package example

import (
	"context"
	"fmt"

	"github.com/wangweihong/gotoolbox/pkg/errors"
	"github.com/wangweihong/gotoolbox/pkg/httpcli"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"
)

type Specification struct {
	c *Client
}

func NewSpecification(c *Client) *Specification {
	return &Specification{c: c}
}

func (p *Specification) SpecificationList(
	ctx context.Context,
	req *iexample.SpecListRequest,
	opts ...httpcli.CallOption,
) (*iexample.SpecListResponse, error) {
	resp := &iexample.SpecListResponse{}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/specification/unified_vm/list").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Specification) SpecificationGet(
	ctx context.Context,
	req *iexample.SpecGetRequest,
	opts ...httpcli.CallOption,
) (*iexample.SpecGetResponse, error) {
	resp := &iexample.SpecGetResponse{}

	if req.Name != "" {
		listReq := iexample.SpecListRequest{
			FilterName:   req.Name,
			IncludeShare: true,
		}
		listResp := &iexample.SpecListResponse{}

		r := httpcli.NewHttpRequestBuilder().
			GET().
			WithPath("/v1/specification/unified_vm/list").
			AddQueryParamByObject(listReq).
			Build()
		_, err := p.c.cc.Invoke(ctx, r, listReq, listResp, opts...)
		if err != nil {
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
		WithPath("/v1/specification/unified_vm/inspect").
		AddQueryParam("uuid", req.UUID).
		AddQueryParam("tenant", req.Tenant).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}
	return resp, nil
}
