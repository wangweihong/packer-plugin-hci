package example

import (
	"context"

	"github.com/wangweihong/gotoolbox/pkg/errors"
	"github.com/wangweihong/gotoolbox/pkg/httpcli"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"
)

type Compute struct {
	c *Client
}

func NewCompute(c *Client) *Compute {
	return &Compute{c: c}
}

func (p *Compute) ComputeGroupList(
	ctx context.Context,
	req *iexample.ComputeGroupListRequest,
	opts ...httpcli.CallOption,
) (*iexample.ComputeGroupListResponse, error) {
	resp := &iexample.ComputeGroupListResponse{}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/compute/group/inspect").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Compute) EcsCreate(
	ctx context.Context,
	req *iexample.EcsCreateRequest,
	opts ...httpcli.CallOption,
) (*iexample.EcsCreateResponse, error) {
	resp := &iexample.EcsCreateResponse{}

	r := httpcli.NewHttpRequestBuilder().
		POST().
		WithPath("/v1/compute/create").
		WithBody("", req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Compute) EcsImport(
	ctx context.Context,
	req *iexample.EcsImportRequest,
	opts ...httpcli.CallOption,
) (*iexample.EcsImportResponse, error) {
	resp := &iexample.EcsImportResponse{}

	r := httpcli.NewHttpRequestBuilder().
		POST().
		WithPath("/v1/compute/unified_import").
		WithBody("", req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Compute) EcsDelete(
	ctx context.Context,
	req *iexample.EcsDeleteRequest,
	opts ...httpcli.CallOption,
) (*iexample.EcsCreateResponse, error) {
	resp := &iexample.EcsCreateResponse{}

	r := httpcli.NewHttpRequestBuilder().
		POST().
		WithPath("/v1/compute/delete").
		WithBody("", req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Compute) EcsInspect(
	ctx context.Context,
	req *iexample.EcsInspectRequest,
	opts ...httpcli.CallOption,
) (*iexample.EcsInspectResponse, error) {
	resp := &iexample.EcsInspectResponse{}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/compute/inspect").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Compute) EcsStart(
	ctx context.Context,
	req *iexample.EcsStartRequest,
	opts ...httpcli.CallOption,
) (*iexample.EcsStartResponse, error) {
	resp := &iexample.EcsStartResponse{}

	r := httpcli.NewHttpRequestBuilder().
		POST().
		WithPath("/v1/compute/start").
		WithBody("", req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Compute) EcsShutoff(
	ctx context.Context,
	req *iexample.EcsShutoffRequest,
	opts ...httpcli.CallOption,
) (*iexample.EcsShutoffResponse, error) {
	resp := &iexample.EcsShutoffResponse{}

	r := httpcli.NewHttpRequestBuilder().
		POST().
		WithPath("/v1/compute/shutoff").
		WithBody("", req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Compute) EcsDiskExport(
	ctx context.Context,
	req *iexample.EcsDiskExportRequest,
	opts ...httpcli.CallOption,
) (*iexample.EcsDiskExportResponse, error) {
	resp := &iexample.EcsDiskExportResponse{}

	r := httpcli.NewHttpRequestBuilder().
		POST().
		WithPath("/v1/compute/disk/export").
		WithBody("", req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}

func (p *Compute) EcsDiskExportProgress(
	ctx context.Context,
	req *iexample.EcsDiskExportProgressRequest,
	opts ...httpcli.CallOption,
) (*iexample.EcsDiskExportProgressResponse, error) {
	resp := &iexample.EcsDiskExportProgressResponse{}

	r := httpcli.NewHttpRequestBuilder().
		GET().
		WithPath("/v1/compute/disk/export/progress").
		AddQueryParamByObject(req).
		Build()
	if _, err := p.c.cc.Invoke(ctx, r, req, resp, opts...); err != nil {
		return nil, errors.UpdateStack(err)
	}

	return resp, nil
}
