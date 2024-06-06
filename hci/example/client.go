package example

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/wangweihong/gotoolbox/pkg/typeutil"

	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	"github.com/wangweihong/gotoolbox/pkg/errors"
	"github.com/wangweihong/gotoolbox/pkg/httpcli"
	"github.com/wangweihong/gotoolbox/pkg/httpcli/httpconfig"
	"github.com/wangweihong/gotoolbox/pkg/httpcli/interceptorcli"
	"github.com/wangweihong/gotoolbox/pkg/skipper"
	"github.com/wangweihong/gotoolbox/pkg/sliceutil"
	"github.com/wangweihong/gotoolbox/pkg/urlutil"
)

type Client struct {
	cc         *httpcli.Client
	endpoint   string
	tenant     string
	user       string
	password   string
	apiKey     string
	uid        string
	lock       sync.RWMutex
	candidates []string
	cluster    string
}

func (c *Client) setServerIps(master string, ips []string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.candidates = sliceutil.StringSlice(ips).MoveFirst(master)
}

func (c *Client) getServerIps() []string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.candidates
}

func (c *Client) setApiKey(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.apiKey = key
}

func (c *Client) getApiKey() string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.apiKey
}

func (c *Client) GetCluster() string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.cluster
}

func NewClient(endpoint, tenant, user, password string, opt ...httpcli.Option) (*Client, error) {
	if endpoint == "" || tenant == "" || user == "" || password == "" {
		return nil, fmt.Errorf("invalid parameter")
	}

	r, err := login(endpoint, tenant, user, password)
	if err != nil {
		return nil, errors.UpdateStack(err)
	}

	if urlutil.Domain(endpoint) != r.Data.SystemMemberList.Leader {
		var err error
		endpoint, err = urlutil.ReplaceURL(endpoint, nil, typeutil.String(r.Data.SystemMemberList.Leader), nil, nil)
		if err != nil {
			return nil, errors.UpdateStack(err)
		}
	}

	apikey, err := getOrCreateApiKey(endpoint, r.Data.Token)
	if err != nil {
		return nil, errors.UpdateStack(err)
	}

	c := &Client{
		endpoint: endpoint,
		tenant:   tenant,
		user:     user,
		password: password,
		cluster:  r.Data.DefaultCluster,
	}
	var cc *httpcli.Client
	if opt != nil {
		cc, err = httpcli.NewClient(nil, opt...)
	} else {
		// 建立长连接?
		HTTPTransport := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second, // 连接超时时间
				KeepAlive: 60 * time.Second, // 保持长连接的时间
			}).DialContext, // 设置连接的参数
			MaxIdleConns:          500,              // 最大空闲连接
			IdleConnTimeout:       60 * time.Second, // 空闲连接的超时时间
			ExpectContinueTimeout: 30 * time.Second, // 等待服务第一个响应的超时时间
			MaxIdleConnsPerHost:   100,              // 每个host保持的空闲连接数
		}
		cfg := httpconfig.DefaultHttpConfig()
		cc, err = httpcli.NewClient(
			cfg,
			httpcli.WithTransport(HTTPTransport),
			httpcli.WithTimeout(30*time.Second),
			httpcli.WithIntercepts(
				// 注意顺序, 队列也靠后的越早执行调用后逻辑
				ApiKeyAuthInterceptor("apikey", c),
				TryAllServersInterceptor("call", c),
				interceptorcli.StatusCodeInterceptor("NoSuccessStatusCodeInterceptor"),
			),
		)
	}
	if err != nil {
		return nil, errors.UpdateStack(err)
	}
	c.cc = cc
	c.setServerIps(r.Data.SystemMemberList.Leader, r.Data.SystemMemberList.Candidates)
	c.setApiKey(apikey)
	return c, nil
}

func ApiKeyAuthInterceptor(name string, c *Client, skipperFunc ...skipper.SkipperFunc) httpcli.Interceptor {
	return httpcli.NewInterceptor("apikey", func(ctx context.Context, req *httpcli.HttpRequest, arg, reply interface{}, cc *httpcli.Client,
		invoker httpcli.Invoker, opts ...httpcli.CallOption) (*httpcli.HttpResponse, error) {
		if skipper.Skip(req.GetPath(), skipperFunc...) {
			return invoker(ctx, req, arg, reply, cc, opts...)
		}

		req.Builder().AddHeaderParam("ApiKey", c.getApiKey())
		rawResp, err := invoker(ctx, req, arg, reply, cc, opts...)
		return rawResp, errors.UpdateStack(err)
	})
}

// 如果请求失败，尝试其他的服务器
func TryAllServersInterceptor(name string, c *Client, skipperFunc ...skipper.SkipperFunc) httpcli.Interceptor {
	return httpcli.NewInterceptor("all", func(ctx context.Context, req *httpcli.HttpRequest, arg, reply interface{}, cc *httpcli.Client,
		invoker httpcli.Invoker, opts ...httpcli.CallOption) (*httpcli.HttpResponse, error) {
		if skipper.Skip(req.GetPath(), skipperFunc...) {
			return invoker(ctx, req, arg, reply, cc, opts...)
		}

		u, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, errors.UpdateStack(err)
		}

		candidates := c.getServerIps()
		for _, v := range candidates {
			addr := fmt.Sprintf("%s://%s:%s", u.Scheme, v, u.Port())
			req.Builder().WithEndpoint(addr).Build()
			rawResp, err := invoker(ctx, req, arg, reply, cc, opts...)
			if err != nil {
				log.Printf("call %v err:%+v\n", req.GetFullRequestAddress(), err)
				// http请求失败, 无法确认当前是否master节点
				// 尝试下一个节点
				continue
			}

			var ret iexample.ResponseResult
			if err := rawResp.Decode(&ret); err != nil {
				continue
			}

			// call fail
			if ret.ErrorCode != 0 {
				//说明此时请求失败，回应的不是master. 继续尝试其他节点
				if ret.ErrorCode == 11057 {
					continue
				}

				// 说明此时是master回应，只不过报了其他错误. 更新节点列表, 返回错误信息
				c.setServerIps(v, c.getServerIps())
				return rawResp, errors.UpdateStack(ret.Error())
			}

			// master请求成功,更新ip列表
			c.setServerIps(v, c.getServerIps())
			if reply != nil {
				if err := rawResp.Decode(reply); err != nil {
					return rawResp, err
				}
			}
			return rawResp, nil
		}

		return nil, errors.UpdateStack(fmt.Errorf("try request all server fail"))
	})
}

func login(endpoint, tenant, user, password string) (*iexample.AuthResponse, error) {
	h := md5.New()
	_, err := io.WriteString(h, password)
	if err != nil {
		return nil, errors.UpdateStack(err)
	}
	encryptedPassword := fmt.Sprintf("%x", h.Sum(nil))

	rawResp, err := httpcli.NewHttpRequestBuilder().
		WithEndpoint(endpoint).
		POST().
		WithPath("/v1/authentication/login").
		WithBody("", &iexample.AuthRequest{
			User:     user,
			Tenant:   tenant,
			Password: encryptedPassword,
		}).Build().Invoke()
	if err != nil {
		return nil, errors.UpdateStack(err)
	}
	if rawResp.GetStatusCode() != http.StatusOK {
		return nil, errors.UpdateStack(fmt.Errorf("request code is not 200,%v", rawResp.GetStatusCode()))
	}
	//
	//var rcode iexample.ResponseResult
	//if err := rawResp.Decode(&rcode); err != nil {
	//	return nil, errors.UpdateStack(err)
	//}

	var r iexample.AuthResponse
	if err := rawResp.Decode(&r); err != nil {
		return nil, errors.UpdateStack(err)
	}

	if r.ErrorCode != 0 {
		return nil, errors.UpdateStack(r.ResponseResult.Error())
	}

	return &r, nil
}

func getOrCreateApiKey(endpoint, token string) (string, error) {
	rawResp, err := httpcli.NewHttpRequestBuilder().
		WithEndpoint(endpoint).
		GET().
		WithPath("/v1/apikey/list").
		AddHeaderParam("Authorization", "Bearer "+token).
		Build().Invoke()
	if err != nil {
		return "", errors.UpdateStack(err)
	}

	if rawResp.GetStatusCode() != http.StatusOK {
		return "", errors.UpdateStack(fmt.Errorf("request code is not 200"))
	}

	var r iexample.ApiKeyListResponse
	if err := rawResp.Decode(&r); err != nil {
		return "", err
	}

	if r.ErrorCode != 0 {
		return "", r.ResponseResult.Error()
	}
	if r.Data == nil || len(r.Data.List) == 0 {
		rawResp, err := httpcli.NewHttpRequestBuilder().
			WithEndpoint(endpoint).
			POST().
			WithPath("/v1/apikey/create").
			AddHeaderParam("Authorization", "Bearer "+token).
			WithBody("", &iexample.ApiKeyCreateRequest{}).
			Build().Invoke()
		if err != nil {
			return "", errors.UpdateStack(err)
		}

		if rawResp.GetStatusCode() != http.StatusOK {
			return "", errors.UpdateStack(fmt.Errorf("request code is not 200"))
		}

		var r iexample.ApiKeyCreateResponse
		if err := rawResp.Decode(&r); err != nil {
			return "", err
		}

		if r.ErrorCode != 0 {
			return "", r.ResponseResult.Error()
		}
		return r.Data.UUID, nil
	}

	return r.Data.List[0].UUID, nil
}
