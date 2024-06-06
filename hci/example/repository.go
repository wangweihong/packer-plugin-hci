package example

import (
	"context"
	"fmt"

	"github.com/wangweihong/gotoolbox/pkg/errors"
	"github.com/wangweihong/gotoolbox/pkg/httpcli"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"
)

type Repository struct {
	c *Client
}

func NewRepository(c *Client) *Repository {
	return &Repository{c: c}
}

func (p *Repository) RepositoryList(
	ctx context.Context,
	req *iexample.RepositoryListRequest,
	opts ...httpcli.CallOption,
) (*iexample.RepositoryListResponse, error) {
	resp := &iexample.RepositoryListResponse{}

	r := httpcli.NewHttpRequestBuilder().
		POST().
		WithPath("/v1/compute/repository/list").
		WithBody("", req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Repository) RepositoryGet(
	ctx context.Context,
	req *iexample.RepositoryGetRequest,
	opts ...httpcli.CallOption,
) (*iexample.RepositoryGetResponse, error) {
	resp := &iexample.RepositoryGetResponse{}

	if req.Name != "" {
		listReq := iexample.RepositoryListRequest{
			FilterName: req.Name,
			Cluster:    req.Cluster,
		}
		listResp := &iexample.RepositoryListResponse{}

		r := httpcli.NewHttpRequestBuilder().
			POST().
			WithPath("/v1/compute/repository/list").
			WithBody("", listReq).
			Build()
		if _, err := p.c.cc.Invoke(ctx, r, listReq, listResp, opts...); err != nil {
			return nil, errors.UpdateStack(err)
		}

		for _, v := range listResp.Data.List {
			if v.Name == req.Name {
				resp.Data = &v
				return resp, nil
			}
		}
		return nil, errors.UpdateStack(fmt.Errorf("repo %v not found", req.Name))
	}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/compute/repository/get").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}
	return resp, nil
}

func (p *Repository) RepositoryImageList(
	ctx context.Context,
	req *iexample.RepositoryImageListRequest,
	opts ...httpcli.CallOption,
) (*iexample.RepositoryImageListResponse, error) {
	resp := &iexample.RepositoryImageListResponse{}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/compute/repository/image/list").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}
	return resp, nil
}

func (p *Repository) RepositoryImageDelete(
	ctx context.Context,
	req *iexample.RepositoryImageDeleteRequest,
	opts ...httpcli.CallOption,
) (*iexample.RepositoryImageDeleteResponse, error) {
	resp := &iexample.RepositoryImageDeleteResponse{}

	r := httpcli.NewHttpRequestBuilder().
		POST().
		WithPath("/v1/compute/repository/image/delete").
		WithBody("", req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}
	return resp, nil
}

// RepositoryImageSync  同步仓库镜像。新创建的镜像，必须同步后才能获取
func (p *Repository) RepositoryImageSync(
	ctx context.Context,
	req *iexample.RepositoryImageSyncRequest,
	opts ...httpcli.CallOption,
) (*iexample.RepositoryImageSyncResponse, error) {
	resp := &iexample.RepositoryImageSyncResponse{}

	r := httpcli.NewHttpRequestBuilder().
		POST().
		WithPath("/v1/compute/repository/image/sync").
		WithBody("", req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}
	return resp, nil
}

func (p *Repository) RepositoryImageGet(
	ctx context.Context,
	req *iexample.RepositoryImageGetRequest,
	opts ...httpcli.CallOption,
) (*iexample.RepositoryImageGetResponse, error) {
	resp := &iexample.RepositoryImageGetResponse{}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/compute/repository/image/get").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}
	return resp, nil
}
