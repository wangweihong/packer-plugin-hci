package iexample

type RepositoryListRequest struct {
	PagingParam
	Cluster           string `json:"cluster_uuid"`
	Tenant            string `json:"tenant"`
	FilterUsage       string `json:"filter_usage"`
	FilterName        string `json:"filter_name"`
	FilterStorageType string `json:"filter_type"`
	Cap               int32  `json:"cap"`
	RepoType          string `json:"repo_type"`
}

type RepositoryACL struct {
	Values []string
}

type RepositoryEntry struct {
	Tenant     string         `json:"tenant"`
	TenantName string         `json:"tenant_name"`
	UUID       string         `json:"uuid"`
	Name       string         `json:"name"`
	Usages     []string       `json:"usages"`
	ACL        *RepositoryACL `json:"acl"`
	Ctime      int64          `json:"ctime"`
	Mtime      int64          `json:"mtime"`
	Capacity   uint64         `json:"capacity"`
	Used       uint64         `json:"used"`
}

type RepositoryListResponse struct {
	ResponseResult
	Data struct {
		TotalCount int               `json:"total_count"`
		List       []RepositoryEntry `json:"list"`
	} `json:"data"`
}

type RepositoryGetRequest struct {
	Cluster string `json:"cluster_uuid"`
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
}

type RepositoryGetResponse struct {
	ResponseResult
	Data *RepositoryEntry `json:"data"`
}

type RepositoryImageListRequest struct {
	PagingParam
	Cluster          string `json:"cluster_uuid"`
	Tenant           string `json:"tenant"`
	FilterName       string `json:"filter_name"`
	FilterRepository string `json:"filter_repository"`
	FilterAvailable  string `json:"filter_available"`
	FilterOsType     string `json:"os_type"`
}

type ImageMeta struct {
	Image struct {
		UUID           string
		Repository     string
		Name           string
		Format         string
		Size           uint64
		Vsize          uint64
		CTime          int64
		LastModified   int64
		Md5            string
		Ref            int32
		Extra          map[string]string
		OsType         string
		RegisterTime   uint64
		VolumeUUID     string
		Available      string
		RepositoryName string
	} `json:"image"`
}

type RepositoryImageListResponse struct {
	ResponseResult
	Data struct {
		Total  int
		Images []ImageMeta
	} `json:"data"`
}

type RepositoryImageGetRequest struct {
	Cluster    string `json:"cluster_uuid"`
	Tenant     string `json:"tenant"`
	UUID       string `json:"uuid"`
	Repository string `json:"repository"`
	Name       string `json:"name"`
}

type RepositoryImageGetResponse struct {
	ResponseResult
	Data *ImageMeta `json:"data"`
}

type RepositoryImageDeleteRequest struct {
	Cluster    string `json:"cluster_uuid"`
	Tenant     string `json:"tenant"`
	UUID       string `json:"uuid"`
	Repository string `json:"repository"`
}

type RepositoryImageDeleteResponse struct {
	ResponseResult
}

type RepositoryImageSyncRequest struct {
	Cluster string `json:"cluster_uuid"`
	Tenant  string `json:"tenant"`
	Image   struct {
		Repository string
	} `json:"image"`
}

type RepositoryImageSyncResponse struct {
	ResponseResult
}
