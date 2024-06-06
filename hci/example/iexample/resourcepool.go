package iexample

type ResourcePoolListRequest struct {
	PagingParam
	Cluster    string `json:"cluster_uuid"`
	FilterName string `json:"filter_name"`
}

type ResourcePoolListResponse struct {
	ResponseResult
	Data struct {
		TotalCount int                 `json:"total_count"`
		List       []ResourcePoolEntry `json:"list"`
	} `json:"data"`
}

type ResourcePoolEntry struct {
	ClusterUUID string `json:"cluster_uuid"`
	Name        string
	UUID        string
}

type ResourcePoolGetRequest struct {
	Cluster string `json:"cluster_uuid"`
	UUID    string `json:"zone_uuid"`
	Name    string `json:"zone_name"`
}

type ResourcePoolGetResponse struct {
	ResponseResult
	Data *ResourcePoolEntry `json:"data"`
}
