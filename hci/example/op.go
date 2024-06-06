package example

import (
	"context"
	"fmt"

	"github.com/wangweihong/gotoolbox/pkg/errors"
	"github.com/wangweihong/gotoolbox/pkg/httpcli"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"
)

type OperationCenter struct {
	c *Client
}

func NewOperationCenter(c *Client) *OperationCenter {
	return &OperationCenter{c: c}
}

func (p *OperationCenter) TenantGet(
	ctx context.Context,
	req *iexample.TenantGetRequest,
	opts ...httpcli.CallOption,
) (*iexample.TenantGetResponse, error) {
	resp := &iexample.TenantGetResponse{}

	if req.TenantName != nil {
		listReq := &iexample.TenantListRequest{FilterName: *req.TenantName, ShowSystem: true}
		listResp := &iexample.TenantListResponse{}

		r := httpcli.NewHttpRequestBuilder().
			GET().
			WithPath("/v1/tenant/list").
			AddQueryParamByObject(listReq).
			Build()
		_, err := p.c.cc.Invoke(ctx, r, listReq, listResp, opts...)
		if err != nil {
			return nil, errors.UpdateStack(err)
		}

		for _, v := range listResp.Data.List {
			if v.Name == *req.TenantName {
				resp = &iexample.TenantGetResponse{
					ResponseResult: listResp.ResponseResult,
					Data:           v,
				}
				return resp, nil
			}
		}
		return nil, errors.UpdateStack(fmt.Errorf("tenant %v not found", *req.TenantName))
	}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/tenant/inspect").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}
