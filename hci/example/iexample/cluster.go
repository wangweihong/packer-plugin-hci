package iexample

type ClusterListRequest struct {
	PagingParam
	FilterName string `json:"filter_name"` //按名字过滤
	FilterUUID string `form:"filter_uuid"` //按uuid过滤
}

type ClusterEntry struct {
	ClusterUUID string `json:"cluster_uuid"`
	ClusterName string `json:"cluster_name"`
	Ctime       int64
	Candidates  []string
	IsStop      bool `json:"is_stop"`
	IsHealth    bool `json:"is_health"`
	Type        string
	Desc        string `json:"desc"`
}

type ClusterListResponse struct {
	Data struct {
		List       []ClusterEntry
		TotalCount int `json:"total_count"`
	}
}

type ClusterGetRequest struct {
	ClusterUUID string `json:"cluster_uuid"`
	ClusterName string `json:"cluster_Name"`
}
