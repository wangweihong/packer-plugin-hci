package iexample

type StoragePoolCreateRequest struct {
	Cluster  string     `json:"cluster_uuid"`
	Tenant   string     `json:"tenant"`
	PoolName string     `json:"pool_name"`
	Quota    *PoolQuota `json:"quota"`
}

type PoolQuota struct {
	StorageSize uint64 `json:"storage_size"`
}

type StoragePoolCreateResponse struct {
	ResponseResult
	Pool StoragePoolBaseInfo `json:"pool"`
}

// DeleteRequest 删除存储池请求
type StoragePoolDeleteRequest struct {
	Pool string `json:"pool_uuid"`
}

// DeleteResponse 删除存储池响应
type StoragePoolDeleteResponse struct {
	Pool StoragePoolBaseInfo `json:"pool"`
}

type StoragePoolGetRequest struct {
	Cluster  string `json:"cluster_uuid"`
	Tenant   string `json:"tenant"`
	PoolUUID string `form:"pool_uuid"`
}

type StoragePoolGetByNameRequest struct {
	Cluster  string `json:"cluster_uuid"`
	Tenant   string `json:"tenant"`
	PoolName string `form:"pool_name"`
}

type StoragePoolGetResponse struct {
	Pool StoragePoolDetail `json:"pool"`
}

type StoragePoolListRequest struct {
	PagingParam
	Cluster    string `json:"cluster_uuid"`
	Tenant     string `json:"tenant"`
	FilterName string `json:"filter_name"`
}

type StoragePoolBaseInfo struct {
	Namespace        string            `json:"namespace"`
	Name             string            `json:"name"`
	UUID             string            `json:"uuid"`
	StorageSize      uint64            `json:"storage_size"`
	UsedSize         uint64            `json:"used_size"`
	Available        uint64            `json:"available"`
	CreateTime       int64             `json:"create_time"`
	VolumeNum        int               `json:"volume_num"`
	Attr             map[string]string `json:"attr"`
	VolumeTotalCount int32             `json:"volume_total_count"`
	Cluster          string            `json:"cluster_uuid"`
	Tenant           string            `json:"tenant"`
}

type StoragePoolDetail struct {
	StoragePoolBaseInfo
	FilterVolumeNum int `json:"filter_volume_num"`
}

type StoragePoolListResponse struct {
	TotalCount int
	List       []StoragePoolBaseInfo `json:"list"`
}
