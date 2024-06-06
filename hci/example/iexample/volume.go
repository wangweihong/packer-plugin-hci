package iexample

type PoolListRequest struct {
	PagingParam
	Cluster                     string `json:"cluster_uuid"`
	Tenant                      string `json:"tenant"`
	FilterNotSystemPool         bool   `json:"filter_not_system_pool"`
	FilterName                  string `json:"filter_name"`
	FilterVolumeDevType         string `json:"filter_volume_dev_type"`
	FilterVolumeShareType       string `json:"filter_volume_share_type"`
	FilterFuzzy                 string `json:"filter_fuzzy"`
	FilterQuotaMoreThan         uint64 `json:"filter_quota_more_than"`
	FilterAvailMoreThan         uint64 `json:"filter_avail_more_than"`
	FilterLabelName             string `json:"filter_label_name"`
	FilterDefaultDiskLabelValue string `json:"filter_default_disk_label_value"`
	FillSupportedDriveType      bool   `json:"fill_supported_drive_type"`
}

type PoolListResponse struct {
	ResponseResult
	Data struct {
		TotalCount int `json:"TotalCount"`
		List       []struct {
			Tenant      string `json:"tenant"`
			TenantName  string `json:"tenant_name"`
			Namespace   string `json:"namespace"`
			Name        string `json:"name"`
			UUID        string `json:"uuid"`
			StorageSize int    `json:"storage_size"`
			UsedSize    int64  `json:"used_size"`
			Available   int    `json:"available"`
			CreateTime  int    `json:"create_time"`
			VolumeNum   int    `json:"volume_num"`
			Attr        struct {
				ComputeRepository string `json:"ComputeRepository"`
			} `json:"attr"`
			VolumeTotalCount int    `json:"volume_total_count"`
			ClusterUUID      string `json:"cluster_uuid"`
			ClusterName      string `json:"cluster_name"`
		} `json:"List"`
	} `json:"data"`
}

type VolumeListRequest struct {
	PagingParam
	NoDetail         *bool   `json:"no_detail"`
	FilterName       *string `json:"filter_name"`
	FilterDevType    *string `json:"filter_dev_type"`
	FilterMountType  *int    `json:"filter_mount_type"`
	FilterVmUUID     *string `json:"filter_vm_uuid"`
	FilterReadOnly   *string `json:"filter_read_only"`
	FilterSetRelated *string `json:"filter_set_related"`
	FilterS3View     *string `json:"filter_s3_view"`
	FilterUserUUID   *string `json:"filter_user_uuid"`
	FilterShareType  *string `json:"filter_share_type"`
	FilterSambaPerm  *string `json:"filter_samba_perm"`
	FilterUUIDs      *string `json:"filter_uuids"`
}

type VolumeListResponse struct {
	ResponseResult
	Data struct {
		TotalCount int           `json:"total_count"`
		List       []VolumeEntry `json:"list"`
	}
}

type VolumeEntry struct {
	Cluster           string `json:"cluster_uuid"`
	Namespace         string
	Pool              string
	Name              string
	UUID              string
	Capacity          uint64
	SnapCapacity      uint64
	SliceSize         uint64
	Status            int32
	Vendor            string
	RW                []string
	RO                []string
	AccessPath        map[string]string
	InvalidAccessPath []string
	Attr              map[string]string
	UserPerms         map[string]int32
	UserGroupPerms    map[string]int32
	SharePoint        map[string]string
	Ctime             int64
	EventState        int64
	Label             map[string]string
	DeletedTime       int64
	PoolName          string
	FinishPercent     float64
	VsdState          map[string]int32
	// 已挂载到导入导出路径
	ImportExportMount bool `json:"import_export_mount"`
}

type VolumeCtrlAccount struct {
	// 用户名
	Username string `json:"user_name"`
	// 密码
	Password string `json:"password"`
	DevType  string `json:"dev_type"`
	// flag = ivsm.ACCOUNT_ADD (1) 时代表新增卷用户，flag = ivsm.ACCOUNT_DEL (2) 时代表新增卷用户，代表需要删除的卷用户
	Flag     uint32 `json:"flag"`
	Outgoing bool   `json:"outgoing"`
}

type VolumeCreateAttribute struct {
	VolumeName      string            `json:"volume_name"`
	Pool            string            `json:"pool_uuid"`
	Capacity        uint64            `json:"capacity"`
	Attribute       map[string]string `json:"attribute"`
	SnapCapacity    *uint64           `json:"snap_capacity"`
	RebuildPriority int32             `json:"rebuild_priority"`
}

type VolumeCreateRequest struct {
	VolumeCreateAttribute
	ClusterUUID string `json:"cluster_uuid"`
	Tenant      string `json:"tenant"`
	// 创建数量
	Count int `json:"count"`
	//  指定挂载主机的UUID
	Hosts []string `json:"hosts"`
	//自动挂载主机数
	AutoMapHostCount *int `json:"auto_mount_host_count"`
}

type VolumeCreateResponse struct {
	ResponseResult
	Data struct {
		Volume struct {
			Tenant       string `json:"tenant"`
			Name         string `json:"name"`
			UUID         string `json:"uuid"`
			Pool         string `json:"pool"`
			Namespace    string `json:"namespace"`
			Capacity     uint64 `json:"capacity"`
			SnapCapacity uint64 `json:"snap_capacity"`
			Status       int32  `json:"status"`
			Ctime        int64  `json:"ctime"`
			Flag         int32  `json:"flag"`
		} `json:"volume"`
	}
}

type VolumeDeleteRequest struct {
	Cluster    string `json:"cluster_uuid"`
	Tenant     string `json:"tenant"`
	Volume     string `json:"volume_uuid"`
	RealDelete *bool  `json:"real_delete"`
}

type VolumeDeleteResponse struct {
	ResponseResult
}
