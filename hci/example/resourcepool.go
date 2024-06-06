package example

import (
	"context"
	"fmt"

	"github.com/wangweihong/gotoolbox/pkg/errors"
	"github.com/wangweihong/gotoolbox/pkg/httpcli"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"
)

type ResourcePool struct {
	c *Client
}

func NewResourcePool(c *Client) *ResourcePool {
	return &ResourcePool{c: c}
}

func (p *ResourcePool) ResourcePoolList(
	ctx context.Context,
	req *iexample.ResourcePoolListRequest,
	opts ...httpcli.CallOption,
) (*iexample.ResourcePoolListResponse, error) {
	resp := &iexample.ResourcePoolListResponse{}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/zone/list").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *ResourcePool) ResourcePoolGet(
	ctx context.Context,
	req *iexample.ResourcePoolGetRequest,
	opts ...httpcli.CallOption,
) (*iexample.ResourcePoolGetResponse, error) {
	resp := &iexample.ResourcePoolGetResponse{}

	if req.Name != "" {
		listReq := iexample.ResourcePoolListRequest{
			FilterName: req.Name,
			Cluster:    req.Cluster,
		}
		listResp := &iexample.ResourcePoolListResponse{}

		r := httpcli.NewHttpRequestBuilder().
			GET().
			WithPath("/v1/zone/list").
			AddQueryParamByObject(listReq).
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
		WithPath("/v1/zone/inspect").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}
	return resp, nil

}
