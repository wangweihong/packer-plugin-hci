package example

import (
	"context"
	"fmt"

	"github.com/wangweihong/gotoolbox/pkg/errors"
	"github.com/wangweihong/gotoolbox/pkg/httpcli"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"
)

type Network struct {
	c *Client
}

func NewNetwork(c *Client) *Network {
	return &Network{c: c}
}

func (p *Network) VirtualSwitchList(
	ctx context.Context,
	req *iexample.VirtualSwitchListRequest,
	opts ...httpcli.CallOption,
) (*iexample.VirtualSwitchListResponse, error) {
	resp := &iexample.VirtualSwitchListResponse{}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/network/layer2/list").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Network) VirtualSwitchGet(
	ctx context.Context,
	req *iexample.VirtualSwitchGetRequest,
	opts ...httpcli.CallOption,
) (*iexample.VirtualSwitchGetResponse, error) {
	resp := &iexample.VirtualSwitchGetResponse{}

	if req.Name != "" {
		listReq := iexample.VirtualSwitchListRequest{
			FilterName: req.Name,
			Cluster:    req.Cluster,
		}
		listResp := &iexample.VirtualSwitchListResponse{}

		r := httpcli.NewHttpRequestBuilder().
			GET().
			WithPath("/v1/network/layer2/list").
			AddQueryParamByObject(req).
			Build()
		if _, err := p.c.cc.Invoke(ctx, r, listReq, listResp, opts...); err != nil {
			return nil, errors.UpdateStack(err)
		}

		for _, v := range listResp.Data.List {
			if v.Name == req.Name {
				req.UUID = v.UUID
				break
			}
		}
	}

	if req.UUID == "" {
		return nil, errors.UpdateStack(fmt.Errorf("vswitch uuid not set"))
	}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/network/layer2/inspect").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Network) VirtualSwitchPortGroupGet(
	ctx context.Context,
	req *iexample.VirtualSwitchPortGroupGetRequest,
	opts ...httpcli.CallOption,
) (*iexample.VirtualSwitchPortGroupGetResponse, error) {
	resp := &iexample.VirtualSwitchPortGroupGetResponse{}

	switchResp, err := p.VirtualSwitchGet(ctx, &iexample.VirtualSwitchGetRequest{
		Cluster: req.Cluster,
		Tenant:  req.Tenant,
		UUID:    req.VSwitchUUID,
		Name:    req.VSwitchName,
	}, opts...)
	if err != nil {
		return nil, errors.UpdateStack(err)
	}

	for _, v := range switchResp.Data.PortGroupList {
		if req.Name == v.PortGroupName {
			resp.Data = &v
			return resp, nil
		}

		if req.UUID == v.PortGroupUUID {
			resp.Data = &v
			return resp, nil
		}
	}

	return nil, fmt.Errorf("portgroup not found")
}
