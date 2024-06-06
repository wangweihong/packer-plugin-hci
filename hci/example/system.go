package example

import (
	"context"

	"github.com/wangweihong/gotoolbox/pkg/errors"
	"github.com/wangweihong/gotoolbox/pkg/httpcli"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"
)

type System struct {
	c *Client
}

func NewSystem(c *Client) *System {
	return &System{c: c}
}

func (p *System) ServerIPList(
	ctx context.Context,
	req *iexample.ServerIPListRequest,
	opts ...httpcli.CallOption,
) (*iexample.ServerIPListResponse, error) {
	resp := &iexample.ServerIPListResponse{}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/system/member/list").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}
