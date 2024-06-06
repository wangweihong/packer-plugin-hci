package iexample

type ComputeGroupListRequest struct {
	PagingParam
	Cluster    string `json:"cluster_uuid"`
	Tenant     string `json:"tenant"`
	FilterName string `json:"filter_name"`
}

type ComputeGroupListResponse struct {
	ResponseResult
	Data struct {
		TotalCount int           `json:"total_count"`
		List       []VolumeEntry `json:"list"`
	} `json:"data"`
}

type ComputeGroupEntry struct {
	UUID      string
	Namespace string
	Name      string
}

type EcsCreateRequest struct {
	// CPU设置
	EcsCPU
	// 内存设置
	EcsMem
	// 元数据
	Meta        EcsMeta          `json:"meta"`
	Interfaces  []ECSInterface   `json:"interfaces"`
	Disks       []EcsDisk        `json:"disks"`
	USBs        []EcsUsb         `json:"usbs"`
	Controllers []EcsControllers `json:"controllers"`
	GPUs        []ECSGPU         `json:"gpus"`
	// 光驱可选
	CDRom *EcsDiskSpecial `json:"cdrom"`
	// 软驱可选
	Floppy *EcsDiskSpecial `json:"floppy"`
	//// 父虚拟机uuid 克隆 基于模板创建会有这个父uuid
	//ParentUuid *string `json:"parent_uuid"`
	//// 镜像uuid，来自镜像仓库repository/get接口
	//ImageUUID *string `json:"image_uuid"`
	// 创建类型
	CreateType *int32 `json:"create_type"`
	// 虚拟机类型
	Type *string `json:"type"`
	// 基准类型
	Baseline *string `json:"baseline"`
	// 系统可选，可通过cdrom/attach接口装系统
	OS *string `json:"os"`
	// 分配的用户
	UserUUID *string `json:"user_uuid"`
	// 分组, 不传默认为default
	ObjectGroup *string `json:"object_group"`
	// 资源池, 不传默认为default
	Zone        *string `json:"zone"`
	ClusterUUID string  `json:"cluster_uuid"`
	Tenant      string  `json:"tenant"`
	// 创建数量
	CreateNum *int `json:"create_num"`
}

type EcsCreateResponse struct {
	ResponseResult
	Data struct {
		Success int `json:"success"`
		Fail    int `json:"fail"`
		Results []struct {
			ResponseResult
			Data *struct {
				Tenant       string   `json:"tenant"`
				UUID         string   `json:"vm_uuid"`
				Name         string   `json:"vm_name"`
				MacAddresses []string `json:"mac_addresses"`
				Group        string   `json:"group"`
				GroupName    string   `json:"group_name"`
			} `json:"data"`
		} `json:"results"`
	} `json:"data"`
}

type EcsDeleteRequest struct {
	ClusterUUID string `json:"cluster_uuid"`
	Tenant      string `json:"tenant"`
	ECS         string `json:"compute_uuid"`
	// 强制删除,不移入回收站
	RealDelete bool `json:"real_delete"`
	//  是否删除卷
	IsDeleteVolumes bool `json:"is_delete_volumes"`
}

type EcsDeleteResponse struct {
	ResponseResult
}

type EcsStartRequest struct {
	Cluster string `json:"cluster_uuid"`
	Tenant  string `json:"tenant"`
	ECS     string `json:"compute_uuid"`
}

type EcsStartResponse struct {
	ResponseResult
}

type EcsShutoffRequest struct {
	Cluster string `json:"cluster_uuid"`
	Tenant  string `json:"tenant"`
	ECS     string `json:"compute_uuid"`
}

type EcsShutoffResponse struct {
	ResponseResult
}

type EcsInspectRequest struct {
	Cluster       string `json:"cluster_uuid"`
	Tenant        string `json:"tenant"`
	ECS           string `json:"compute_uuid"`
	ShowPlaintext bool   `json:"show_plaintext"`
}

type EcsInspectResponse struct {
	ResponseResult
	Data struct {
		Tenant                  string           `json:"tenant"`
		TenantName              string           `json:"tenant_name"`
		OwnerIsLdap             bool             `json:"owner_is_ldap"`
		Owner                   string           `json:"owner"`
		OwnerName               string           `json:"owner_name"`
		UUID                    string           `json:"uuid"`
		Name                    string           `json:"name"`
		Desc                    string           `json:"desc"`
		Type                    string           `json:"type"`
		CreateTime              int              `json:"create_time"`
		ModifyTime              int              `json:"modify_time"`
		DeletedTime             int              `json:"deleted_time"`
		State                   string           `json:"state"`
		GaState                 string           `json:"ga_state"`
		Action                  string           `json:"action"`
		Os                      string           `json:"os"`
		FirstShownIP            string           `json:"first_shown_ip"`
		VdcAgent                string           `json:"vdc_agent"`
		VdcAgentList            interface{}      `json:"vdc_agent_list"`
		VdcAgentInfo            interface{}      `json:"vdc_agent_info"`
		VdcPoolUUID             string           `json:"vdc_pool_uuid"`
		VdcPoolName             string           `json:"vdc_pool_name"`
		CreateType              string           `json:"create_type"`
		Namespace               string           `json:"namespace"`
		MachineUUID             string           `json:"machine_uuid"`
		ManagerIP               string           `json:"manager_ip"`
		Labels                  interface{}      `json:"labels"`
		InterfaceIps            []EcsInterfaceIP `json:"interface_ips"`
		Attr                    interface{}      `json:"attr"`
		ObjectGroup             string           `json:"object_group"`
		ObjectGroupName         string           `json:"object_group_name"`
		Plan                    string           `json:"plan"`
		PlanRunningObject       string           `json:"plan_running_object"`
		PlanName                string           `json:"plan_name"`
		PlanMasterClusterUUID   string           `json:"plan_master_cluster_uuid"`
		PlanMasterClusterName   string           `json:"plan_master_cluster_name"`
		PlanSlaveClusterUUID    string           `json:"plan_slave_cluster_uuid"`
		PlanSlaveClusterName    string           `json:"plan_slave_cluster_name"`
		RunningCluster          string           `json:"running_cluster"`
		Zone                    string           `json:"zone"`
		ZoneName                string           `json:"zone_name"`
		Strategys               interface{}      `json:"strategys"`
		StrategyNames           interface{}      `json:"strategy_names"`
		StratrgyTypes           interface{}      `json:"stratrgy_types"`
		Strategy                string           `json:"strategy"`
		StrategyName            string           `json:"strategy_name"`
		SpicePortMapping        bool             `json:"spice_port_mapping"`
		ConsoleServerListenPort int              `json:"console_server_listen_port"`
		TaskTag                 int              `json:"task_tag"`
		CPUQosCeil              int              `json:"cpu_qos_ceil"`
		NetworkElementType      struct {
			Type string `json:"type"`
			Spec string `json:"spec"`
		} `json:"network_element_type"`
		PowerTime             int         `json:"power_time"`
		IsSensitive           bool        `json:"is_sensitive"`
		SensitiveCPUUsage     interface{} `json:"sensitive_cpu_usage"`
		MemSize               int64       `json:"mem_size"`
		MemMax                int64       `json:"mem_max"`
		MemoryHugePagesEnable bool        `json:"memory_huge_pages_enable"`
		MemoryQos             struct {
			HardLimit     int `json:"hard_limit"`
			SoftLimit     int `json:"soft_limit"`
			SwapHardLimit int `json:"swap_hard_limit"`
			MinGuarantee  int `json:"min_guarantee"`
		} `json:"memory_qos"`
		MemReserve    bool   `json:"mem_reserve"`
		MemAutoExpand bool   `json:"mem_auto_expand"`
		MemExpandNum  int    `json:"mem_expand_num"`
		ComputeName   string `json:"compute_name"`
		OsLoaderPath  string `json:"os_loader_path"`
		Video         []struct {
			ModelType string `json:"model_type"`
			Heads     int    `json:"heads"`
			RAMSize   int    `json:"ram_size"`
			VgaMem    int    `json:"vga_mem"`
		} `json:"video"`
		StartupSetting struct {
			Hosts  []interface{} `json:"hosts"`
			Labels []interface{} `json:"labels"`
		} `json:"startup_setting"`
		VncEnable     bool `json:"vnc_enable"`
		MemSecEnabled bool `json:"mem_sec_enabled"`
		HaConfig      struct {
			Priority   int         `json:"priority"`
			WarnConfig interface{} `json:"warn_config"`
		} `json:"ha_config"`
		NtpSyncConfig          interface{}   `json:"ntp_sync_config"`
		OsTypeMachine          string        `json:"os_type_machine"`
		HypervSwitch           string        `json:"hyperv_switch"`
		MouseInputType         string        `json:"mouse_input_type"`
		BootMenuTimeout        string        `json:"boot_menu_timeout"`
		CPUMode                string        `json:"cpu_mode"`
		MachineManagerIP       string        `json:"machine_manager_ip"`
		MachineStorageIP       string        `json:"machine_storage_ip"`
		ParentUUID             string        `json:"parent_uuid"`
		MigrateDest            string        `json:"migrate_dest"`
		Baseline               string        `json:"baseline"`
		DeviceName             string        `json:"device_name"`
		ImportMachine          string        `json:"import_machine"`
		ExceedClusterBaseline  bool          `json:"exceed_cluster_baseline"`
		CPUCurrent             string        `json:"cpu_current"`
		CPUMaxCount            int           `json:"cpu_max_count"`
		CPUNrCount             int           `json:"cpu_nr_count"`
		CPUSockets             int           `json:"cpu_sockets"`
		CPUSocketCores         int           `json:"cpu_socket_cores"`
		Shares                 int           `json:"shares"`
		VncPort                int           `json:"vnc_port"`
		VncPassword            string        `json:"vnc_password"`
		VncPasswordPlaintext   string        `json:"vnc_password_plaintext"`
		StreamingMode          string        `json:"streaming_mode"`
		CompressionMode        string        `json:"compression_mode"`
		SpicePort              int           `json:"spice_port"`
		SpicePassword          string        `json:"spice_password"`
		SpicePasswordPlaintext string        `json:"spice_password_plaintext"`
		Panic                  string        `json:"panic"`
		Interfaces             []interface{} `json:"interfaces"`
		Cdrom                  struct {
			Attached       bool   `json:"attached"`
			Name           string `json:"name"`
			RepositoryUUID string `json:"repository_uuid"`
			TargetDev      string `json:"target_dev"`
		} `json:"cdrom"`
		Floppy struct {
			TargetDev string `json:"target_dev"`
		} `json:"floppy"`
		Disks []struct {
			DiskTargetDevice string `json:"disk_target_device"`
			DiskTargetBus    string `json:"disk_target_bus"`
			DiskVsdSource    struct {
				VolumeUUID      string      `json:"volume_uuid"`
				VolumeName      string      `json:"volume_name"`
				DiskCapacity    int64       `json:"disk_capacity"`
				PoolUUID        string      `json:"pool_uuid"`
				PoolName        string      `json:"pool_name"`
				Replica         string      `json:"replica"`
				DriveType       string      `json:"drive_type"`
				Redundancy      int         `json:"redundancy"`
				Attribute       interface{} `json:"attribute"`
				VsdVendor       string      `json:"vsd_vendor"`
				DataDevType     string      `json:"data_dev_type"`
				DataType        string      `json:"data_type"`
				ComponentSetNum string      `json:"component_set_num"`
				StripeShift     string      `json:"stripe_shift"`
				ComponentShift  string      `json:"component_shift"`
			} `json:"disk_vsd_source"`
			BackingStore struct {
				Index        int         `json:"index"`
				Format       interface{} `json:"format"`
				Source       interface{} `json:"source"`
				BackingStore interface{} `json:"backing_store"`
			} `json:"backing_store"`
			DisksSerial   string `json:"disks_serial"`
			Wwn           string `json:"wwn"`
			DiskIsSystem  bool   `json:"disk_is_system"`
			ReadBytesSec  int    `json:"read_bytes_sec"`
			WriteBytesSec int    `json:"write_bytes_sec"`
			ReadIopsSec   int    `json:"read_iops_sec"`
			WriteIopsSec  int    `json:"write_iops_sec"`
			DiskDiscard   string `json:"disk_discard"`
		} `json:"disks"`
		DeviceOrder []string `json:"device_order"`
		DiskType    []struct {
			Type  int    `json:"type"`
			Value string `json:"value"`
		} `json:"disk_type"`
		Usbs []struct {
			Bus         int    `json:"bus"`
			Port        string `json:"port"`
			ProductID   string `json:"product_id"`
			ProductName string `json:"product_name"`
			VendorID    string `json:"vendor_id"`
			VendorName  string `json:"vendor_name"`
			SourceType  string `json:"source_type"`
			Postscript  string `json:"postscript"`
		} `json:"usbs"`
		Controllers []struct {
			Type  string `json:"type"`
			Index int    `json:"index"`
			Model string `json:"model"`
		} `json:"controllers"`
		IsExternalDisk         bool `json:"is_external_disk"`
		HasPassthroughDisk     bool `json:"has_passthrough_disk"`
		KvmClockEnable         bool `json:"kvm_clock_enable"`
		VirtNestedEnable       bool `json:"virt_nested_enable"`
		CPUReserveEnable       bool `json:"cpu_reserve_enable"`
		CPUReserveNuma         bool `json:"cpu_reserve_numa"`
		CPUNumaNetAffinity     bool `json:"cpu_numa_net_affinity"`
		MemoryReserveEnable    bool `json:"memory_reserve_enable"`
		CPUHygonFeatureDisable bool `json:"cpu_hygon_feature_disable"`
		ExpandConfig           struct {
			CPUAutoExpandEnable    bool `json:"cpu_auto_expand_enable"`
			MemoryAutoExpandEnable bool `json:"memory_auto_expand_enable"`
			MaxExpandCPU           int  `json:"max_expand_cpu"`
			MaxExpandMem           int  `json:"max_expand_mem"`
			ExpandCPU              int  `json:"expand_cpu"`
			ExpandMem              int  `json:"expand_mem"`
		} `json:"expand_config"`
		DataLocalizationModeOn string      `json:"data_localization_mode_on"`
		PciDevices             interface{} `json:"pci_devices"`
		MdevDevices            interface{} `json:"mdev_devices"`
		TpmDevices             interface{} `json:"tpm_devices"`
		VgpuDevices            interface{} `json:"vgpu_devices"`
		RecoverToRelatedZones  int         `json:"recover_to_related_zones"`
		MasterZoneTakeOver     int         `json:"master_zone_take_over"`
		RestoreType            string      `json:"restore_type"`
		RestorePeriod          string      `json:"restore_period"`
		ShowPlaintext          bool        `json:"show_plaintext"`
		IsAuth                 bool        `json:"is_auth"`
		NetworkSetting         struct {
			State   string `json:"state"`
			Disable bool   `json:"disable"`
		} `json:"network_setting"`
		RestoreDrive       interface{} `json:"restore_drive"`
		VsdVolumeCount     int         `json:"vsd_volume_count"`
		ShareVolumeCount   int         `json:"share_volume_count"`
		AdpProxyEnabled    bool        `json:"adp_proxy_enabled"`
		AdpProxyScheme     string      `json:"adp_proxy_scheme"`
		AdpProxyRouterIP   string      `json:"adp_proxy_router_ip"`
		AdpProxyRouterPort int         `json:"adp_proxy_router_port"`
		AdpProxyHostIP     string      `json:"adp_proxy_host_ip"`
		AdpProxyHostPort   int         `json:"adp_proxy_host_port"`
		AdpProxyErrReason  string      `json:"adp_proxy_err_reason"`
		AdpProxyGroupUUID  string      `json:"adp_proxy_group_uuid"`
		AdpProxyGroupName  string      `json:"adp_proxy_group_name"`
		AllowDirectDesktop bool        `json:"allow_direct_desktop"`
		SupportSetStaticIP int         `json:"support_set_static_ip"`
		ClusterUUID        string      `json:"cluster_uuid"`
		ClusterName        string      `json:"cluster_name"`
	} `json:"data"`
}

type EcsInterfaceIP struct {
	AddrType string `json:"addr_type"`
	IP       string `json:"ip"`
	Address  string `json:"address"`
	Family   string `json:"family"`
	Prefix   int    `json:"prefix"`
	Peer     string `json:"peer"`
	Name     string `json:"name"`
}

type EcsImportRequest struct {
	// CPU设置
	EcsCPU
	// 内存设置
	EcsMem
	// 元数据
	Meta        EcsMeta          `json:"meta"`
	Interfaces  []ECSInterface   `json:"interfaces"`
	Disks       []EcsDisk        `json:"disks"`
	USBs        []EcsUsb         `json:"usbs"`
	Controllers []EcsControllers `json:"controllers"`
	GPUs        []ECSGPU         `json:"gpus"`
	// 光驱可选
	CDRom *EcsDiskSpecial `json:"cdrom"`
	// 软驱可选
	Floppy *EcsDiskSpecial `json:"floppy"`
	//// 父虚拟机uuid 克隆 基于模板创建会有这个父uuid
	//ParentUuid *string `json:"parent_uuid"`
	//// 镜像uuid，来自镜像仓库repository/get接口
	//ImageUUID *string `json:"image_uuid"`
	// 创建类型
	CreateType *int32 `json:"create_type"`
	// 虚拟机类型
	Type *string `json:"type"`
	// 基准类型
	Baseline *string `json:"baseline"`
	// 系统可选，可通过cdrom/attach接口装系统
	OS *string `json:"os"`
	// 分配的用户
	UserUUID *string `json:"user_uuid"`
	// 分组, 不传默认为default
	ObjectGroup *string `json:"object_group"`
	// 资源池, 不传默认为default
	Zone        *string `json:"zone"`
	ClusterUUID string  `json:"cluster_uuid"`
	Tenant      string  `json:"tenant"`
	// 创建数量
	CreateNum *int `json:"create_num"`
}

type EcsImportResponse struct {
	ResponseResult
	Data struct {
		Success int `json:"success"`
		Fail    int `json:"fail"`
		Results []struct {
			ResponseResult
			Data *struct {
				Tenant       string   `json:"tenant"`
				VmUUID       string   `json:"vm_uuid"`
				Index        int      `json:"index"`
				VMName       string   `json:"vm_name"`
				Volumes      []string `json:"volume_uuids"`
				MacAddresses []string `json:"mac_addresses"`
				Group        string   `json:"group"`
				GroupName    string   `json:"group_name"`
			} `json:"data"`
		} `json:"results"`
	} `json:"data"`
}

type EcsDiskExportRequest struct {
	Cluster  string              `json:"cluster_uuid"`
	Tenant   string              `json:"tenant"`
	ECS      string              `json:"compute_uuid"`
	Compress bool                `json:"compress"`
	Disk     []EcsDiskExportInfo `json:"disk"`
}

type EcsDiskExportInfo struct {
	// sda
	Device       string `json:"device"`
	SnapshotUUID string `json:"snapshot_uuid"`
	ImageName    string `json:"image_name"`
	//raw,qcow2
	ImageType      string `json:"image_type"`
	RepositoryUUID string `json:"repository_uuid"`
}

type EcsDiskExportResponse struct {
	ResponseResult
}

type EcsDiskExportProgressRequest struct {
	Cluster string `json:"cluster_uuid"`
	Tenant  string `json:"tenant"`
	ECS     string `json:"compute_uuid"`
}

type EcsDiskExportProgressResponse struct {
	ResponseResult
	Data struct {
		Progress int `json:"progress"`
	}
}
