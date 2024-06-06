package iexample

type SpecListRequest struct {
	PagingParam
	Tenant       string `json:"tenant"`
	FilterName   string `json:"filter_name"`
	IncludeShare bool   `json:"include_share"`
}

type SpecListResponse struct {
	ResponseResult
	Data struct {
		TotalCount int         `json:"total_count"`
		List       []SpecEntry `json:"list"`
	} `json:"data"`
}

type SpecEntry struct {
	Name          string            `json:"name"`
	UUID          string            `json:"uuid"`
	CreateTime    int               `json:"create_time"`
	Tenant        string            `json:"tenant"`
	ClusterUUID   string            `json:"cluster_uuid"`
	Desc          string            `json:"desc"`
	UnifiedVmSpec *EcsCreateRequest `json:"unified_vm_spec"`
	IsShare       bool              `json:"is_share"`
}

type SpecGetRequest struct {
	Tenant string `json:"tenant"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
}

type SpecGetResponse struct {
	ResponseResult
	Data *SpecEntry `json:"data"`
}
