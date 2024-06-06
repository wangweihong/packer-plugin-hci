package example

import (
	"context"

	"github.com/wangweihong/gotoolbox/pkg/errors"
	"github.com/wangweihong/gotoolbox/pkg/httpcli"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"
)

type Volume struct {
	c *Client
}

func NewVolume(c *Client) *Volume {
	return &Volume{c: c}
}

func (p *Volume) PoolList(
	ctx context.Context,
	req *iexample.PoolListRequest,
	opts ...httpcli.CallOption,
) (*iexample.PoolListResponse, error) {
	resp := &iexample.PoolListResponse{}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/pool/internal/list").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Volume) VolumeList(
	ctx context.Context,
	req *iexample.VolumeListRequest,
	opts ...httpcli.CallOption,
) (*iexample.VolumeListResponse, error) {
	resp := &iexample.VolumeListResponse{}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/volume/internal/list").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Volume) VolumeCreate(
	ctx context.Context,
	req *iexample.VolumeCreateRequest,
	opts ...httpcli.CallOption,
) (*iexample.VolumeCreateResponse, error) {
	resp := &iexample.VolumeCreateResponse{}

	r := httpcli.NewHttpRequestBuilder().
		POST().
		WithPath("/v1/volume/internal/create").
		WithBody("", req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}
