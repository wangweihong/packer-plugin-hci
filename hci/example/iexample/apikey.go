package iexample

type ApiKeyData struct {
	UUID string `json:"uuid"`
}

type ApiKeyListRequest struct {
	PagingParam
	FilterUser string `form:"filter_user"`
	Type       string `form:"type"`
}

type ApiKeyListResponse struct {
	ResponseResult
	Data *struct {
		List       []*ApiKeyData `json:"list"`
		TotalCount int           `json:"total_count"`
	} `json:"data"`
}

type ApiKeyCreateRequest struct {
}

type ApiKeyCreateResponse struct {
	ResponseResult
	Data *ApiKeyData `json:"data"`
}
