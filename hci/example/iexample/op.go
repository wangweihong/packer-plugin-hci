package iexample

type TenantGetRequest struct {
	Tenant     string  `json:"tenant"`
	TenantName *string `json:"tenant_name"`
}

type TenantGetResponse struct {
	ResponseResult
	Data TenantListEntry `json:"data"`
}

type TenantListRequest struct {
	PagingParam
	ShowSystem bool   `json:"system" form:"system"`
	FilterName string `json:"filter_name" form:"filter_name"`
}

type TenantListEntry struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// ListResponse 租户列表回应
type TenantListResponse struct {
	ResponseResult
	Data struct {
		List       []TenantListEntry `json:"list"`
		TotalCount int               `json:"total_count"`
	} `json:"data"`
}
