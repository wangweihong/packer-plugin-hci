package iexample

type EcsUsb struct {
	PassThroughUSBBus  *int   `json:"pass_through_usb_bus" description:"usb透传总线 不为空则指定USB透传 0: uhci(usb1.1, 2 ports)  1: ehci(usb2.0, 6 ports)  2: xhci(usb3.0, 4 ports)"` //// 不为空则指定USB透传 //0: uhci(usb1.1, 2 ports)  1: ehci(usb2.0, 6 ports)  2: xhci(usb3.0, 4 ports)
	PassThroughUSBPort *int   `json:"pass_through_usb_port" description:"usb透传端口 可忽略，1~n"`                                                                                 //// 可忽略，1~n
	RedirectUSBBus     *int   `json:"redirect_usb_bus" description:"usb重定向总线"`
	RedirectUSBPort    *int   `json:"redirect_usb_port" description:"usb重定向端口"`
	SpiceVmcUSBBus     *int   `json:"spice_vmc_usb_bus" description:"usb控制台重定向总线"`
	SpiceVmcUSBPort    *int   `json:"spice_vmc_usb_port" description:"usb控制台重定向总线"`
	SourceUSBHostUUID  string `json:"source_usb_host_uuid" description:"指定USB设备所在主机, 虚拟机使用其他主机的usb需要重定向"` // // 指定USB设备所在主机, 虚拟机使用其他主机的usb需要重定向
	SourceUSBBus       int    `json:"source_usb_bus" description:"指定USB设备bus,数据与列表获取的一致"`
	SourceUSBDevice    int    `json:"source_usb_device" description:"指定USB设备device,数据与列表获取的一致"` //// 指定USB设备bus，device，数据与列表获取的一致
	UsbPostscript      string `json:"usb_postscript" description:"用户自定义usb描述"`
}

type InterfaceBound struct {
	// 平均值
	Average int32 `json:"average"`
	// 峰值
	Peak int32 `json:"peak"`
	// 突发值
	Burst int32 `json:"burst"`
}

type GPUScheme struct {
	Framebuffer   string `json:"framebuffer"`
	FrlConfig     string `json:"frl_config"`
	AuthorizeType string `json:"authorize_type"`
	GPUTypeID     string `json:"gpu_type_id"`
	MaxInstance   string `json:"max_instance"`
	MaxResolution string `json:"max_resolution"`
	NumHeads      string `json:"num_heads"`
	Model         string `json:"model"`
	Passthrough   bool   `json:"passthrough"`
}

type ECSGPU struct {
	UUID     string    `json:"gpu_uuid"`
	Scheme   GPUScheme `json:"scheme"`
	VGpuUUID string    `json:"vgpu_uuid"`
	Model    string    `json:"model"`
	Display  string    `json:"display" `
}

type InterfaceIP struct {
	AddrType string `json:"addr_type"`
	Ip       string `json:"ip"`
	Address  string `json:"address"`
	Family   string `json:"family"`
	Prefix   uint32 `json:"prefix"`
	Peer     string `json:"peer"`
	Name     string `json:"name"`
}

type ECSInterface struct {
	// 网卡UUID
	InterfaceID string `json:"interface_id"`
	// 网卡驱动
	InterfaceModelType string `json:"interface_model_type"`
	// 网桥名称
	InterfaceType   string `json:"interface_type"`
	InterfaceBridge string `json:"interface_bridge"`
	// mac地址
	MacAddress string        `json:"mac_address"`
	Ips        []InterfaceIP `json:"ips"`
	TargetDev  string        `json:"target_dev"`
	// 接收速率
	InBound *InterfaceBound `json:"in_bount"`
	// 发送速率
	OutBound *InterfaceBound `json:"out_bount"`
	// vpc交换机UUID
	SwitchUUID string `json:"switch_uuid"`
	// vpc出口UUID
	GatewayUUID       string `json:"gateway_uuid"`
	ConnectDomainUUID string `json:"connect_domain_uuid"`
	// 端口组UUID
	PortGroup  string   `json:"port_group"`
	VlanTagIDs []uint32 `json:"vlan_tag_ids"`
	//是否老的网卡
	IsOld  bool    `json:"is_old"`
	Driver *string `json:"driver"`
	// 网卡多队列
	VhostQueue uint32 `json:"vhost_queue"`
	BoundIPv6  string `json:"bound_ipv6"`
	//绑定IP地址
	BoundIP string `json:"bound_ip"`
	// ip绑定模式
	Peer string `json:"peer"`
	Name string `json:"name"`
	// windows本地ipv6
	WindowsLocalIPv6 string `json:"windows_local_ipv6"`
	AddressPCISlot   uint32 `json:"address_pci_slot"`
	// 发包延时
	Latency uint32 `json:"in_bound_latency"`
	// 发包丢包率
	Loss uint32 `json:"in_bound_loss"`
}

type DiskNetworkSource struct {
	// iscsi盘的iqn
	Name string `json:"name"`
	// 主机ip
	HostName string `json:"host_name"`
	// 主机端口号
	HostPort string `json:"host_port"`
}

type DiskVolumeSource struct {
	// 存储池
	Pool string `json:"pool"`
	// 卷名称
	Volume string `json:"volume"`
	// 卷uuid
	VolumeUUID string `json:"volume_uuid"`
	// 磁盘容量
	DiskCapacity uint64 `json:"disk_capacity"`
	Replica      string `json:"replica"`
	DriveType    string `json:"drive_type"`
	DataDevType  string `json:"data_dev_type"`
	// 池的可用挂载主机，为空代表可以挂载资源池下所有主机
	PoolAvailHosts []string `json:"pool_avail_hosts"`
}

type DiskVsdSource struct {
	VolumeUUID string `json:"volume_uuid"`
	// 存储池uuid
	PoolUUID    string            `json:"pool_uuid"`
	Replica     string            `json:"replica"`
	DriveType   string            `json:"drive_type"`
	Redundancy  uint32            `json:"redundancy"`
	Attribute   map[string]string `json:"attribute"`
	DataDevType string            `json:"data_dev_type"`
	DataType    string            `json:"data_type"`
}

type DiskBlockSource struct {
	VolumeUUID   string `json:"volume_uuid"`
	DataDevType  string `json:"data_dev_type"`
	DataType     string `json:"data_type"`
	IOAccelerate bool   `json:"io_accelerate"`
}

type DiskPassThroughSource struct {
	DiskId   string `json:"disk_id"`
	HostUUID string `json:"host_uuid"`
	HostIP   string `json:"host_ip"`
}

type DiskImageFileSource struct {
	Name string `json:"name"`
	Size uint64 `json:"size"`
	Path string `json:"path"`
}

type DiskUnifiedVsdSource struct {
	VolumeCreateAttribute
}

type EcsDisk struct {
	Disk
	// 创建新的卷作为磁盘
	DiskUnifiedVsd *DiskUnifiedVsdSource `json:"disk_unified_vsd,omitempty"`
	Redundancy     uint32                `json:"redundancy,omitempty"`
	Device         string                `json:"device,omitempty"`
	//ova磁盘文件路径
	FilePath string `json:"file_path,omitempty"`
	// 导入磁盘镜像仓库
	RepositoryUUID string `json:"repository_uuid,omitempty"`
	// 导入镜像
	ImageUUID string `json:"image_uuid,omitempty"`
}

type Disk struct {
	DiskRWRate
	// 虚拟机盘符，以下卷挂载到主机后的盘符。 指定驱动时该项为空.hda~hdc, sda~sdo, vda~vdo
	DiskTargetDevice string `json:"disk_target_device"`
	// 有效值： ide, virtio, scsi
	DiskTargetBus string `json:"disk_target_bus"`
	// iscsi
	DiskSourceNetwork *DiskNetworkSource `json:"disk_network_source"`
	// 外置存储
	DiskSourceVolume *DiskVolumeSource `json:"disk_volume_source"`
	// vsd 内置卷
	DiskSourceVsd *DiskVsdSource `json:"disk_vsd_source"`
	// 块设备
	DiskSourceBlock *DiskBlockSource `json:"disk_block_source"`
	// 透传物理硬盘
	DiskSourcePassThrough *DiskPassThroughSource `json:"disk_source_passthrough"`
	// 映像文件作为磁盘
	DiskImageFileSource *DiskImageFileSource `json:"disk_image_file_source"`
	// 缓存模式 none:无缓存, writeback:写回模式, writethrough:写通模式"
	DiskCacheMode *string `json:"disk_cache_mode"`
	DisksSerial   *string `json:"disks_serial"`
	WWN           string  `json:"wwn,omitempty"`
	// 是否为系统盘,用于指定某个盘来启动。唯一
	DiskIsSystem bool `json:"disk_is_system"`
	// 磁盘是否启用io悬挂
	DiskIoHang *bool `json:"disk_io_hang"`

	// 磁盘空间回收
	DiskDiscard string `json:"disk_discard"`
}

type DiskRWRate struct {
	// 每秒读速
	ReadBytesSec uint64 `json:"read_bytes_sec"`
	// 每秒写速
	WriteBytesSec uint64 `json:"write_bytes_sec"`
	// 每秒读iops
	ReadIopsSec uint64 `json:"read_iops_sec"`
	// 每秒写iops
	WriteIopsSec uint64 `json:"write_iops_sec"`
}

type EcsControllers struct {
	// 控制器类型
	Type string `json:"type"`
	// 序号
	Index uint32 `json:"index"`
	// 控制器型号
	Model string `json:"model"`
	// only present in type "usb"
	// usb控制器端口
	USBPorts *uint32 `json:"usb_ports,omitempty"`
	// usb主控制器端口
	USBMasterPort *uint32 `json:"usb_master_port,omitempty"`
}

type EcsVideo struct {
	// displays on this video device
	Heads uint32 `json:"heads"`

	// 显示类型 qxl,vga
	ModelType string `json:"model_type"`
	// 显存大小
	RamSize uint32 `json:"ram_size"`
	// 决定显示器的最大分辨率
	VGAMem uint32 `json:"vga_mem"`
}

type NtpConfig struct {
	// ntp服务器地址来源，1：默认用集群配置好的ntp服务器地址；2：自定义
	Source int32
	// 当前使用的ntp服务器允许用户设置多个备选的ntp服务器，但必须要选择一个主ntp服务器来表示当前使用的ntp服务器
	MainNtpServer string
	// 备选ntp服务器地址
	SlaveNtpServers []string
	// 刷新频率，间隔多久时间同步一次ntp服务器时间
	FlushPeriod int64
	// 故障重试周期，如：ntp服务器故障时，会间隔 RetryPeriod 时间之后尝试重新同步时间
	RetryPeriod int64
}

type NtpSyncConfig struct {
	// 时区id，1为UTC（0时区）、2为CST（中国时区）
	TimeZone int32
	// ntp配置
	NtpConfig *NtpConfig
}

type EcsMeta struct {
	// 修改/设置 虚拟机的名字
	ComputeName string `json:"compute_name"`
	// HA模式　auto,on,off,suggest
	HAMode string `json:"ha_mode,omitempty"`
	// 启动方式 BIOS, UEFI
	OsLoaderPath string      `json:"os_loader_path,omitempty"`
	Video        []*EcsVideo `json:"video,omitempty"`
	// 蓝屏检测(异常检测)　可选:isa, none
	PanicModel string `json:"panic_model,omitempty"`
	// 开启VNC
	VncEnable *bool `json:"vnc_enable"`
	// 非空表示修改内存安全功能设置
	MemSecEnabled *bool `json:"mem_sec_enabled,omitempty"`
	// 虚拟机多个ip，选择显示的ip
	ShownIp string `json:"shown_ip,omitempty"`
	// 创建时嵌套虚拟化
	VirtNestedEnabled *bool `json:"virt_nested_enable,omitempty"`
	// 创建时数据本地化
	DataLocalizationModeOn string `json:"data_localization_mode_on,omitempty"`
	RecoverToRelatedZones  int32  `json:"recover_to_related_zones,omitempty"`
	MasterZoneTakeOver     int32  `json:"master_zone_take_over,omitempty"`
	// 虚拟机描述
	ComputeDesc string `json:"compute_desc,omitempty"`
	// 虚拟机ntp时间同步配置
	NtpSyncConfig *NtpSyncConfig `json:"ntp_sync_config"`
	OsTypeMachine string         `json:"os_type_machine,omitempty"`
	HypervSwitch  string         `json:"hyperv_switch"`
	// 鼠标输入类型 ps2 usb
	DomainMouseInputType string `json:"mouse_input_type"`
	// 启动菜单超时时间 必须为数字0~65535
	DomainBootMenuTimeOut string `json:"boot_menu_timeout"`
	// 虚拟机cpu模式:host-passthrough直通 custom透传
	DomainCpuMode string `json:"cpu_mode"`
}

type EcsCPU struct {
	CPUnr int32 `json:"cpunr"`
	// 最大cpu可选
	MaxCPUnr int32 `json:"max_cpu_nr"`
	// Cpu插槽数
	CpuSockets *int `json:"cpu_sockets"`
	//Cpu核数
	CpuSocketCores *int `json:"cpu_socket_cores"`
	// 是否开启cpu预留
	CPUReserve bool `json:"cpu_reserve"`
	// 是否开启numa亲和
	CPUNuma bool `json:"cpu_numa"`
	// 是否开启cpu自动扩展
	CPUAutoExpand bool `json:"cpu_auto_expand"`
	// cpu自动扩展数量
	CPUExpandNum int32 `json:"cpu_expand_num"`
	// numa网络亲和
	CPUNumaNetAffinity bool `json:"cpu_numa_net_affinity"`
}

type EcsMem struct {
	// 内存大小,默认单位是字节
	MemorySize uint64 `json:"mem_size" `
	// 内存上限
	MemoryMax uint64 `json:"mem_max"`
	// 是否开启内存大页
	MemoryHugePagesEnabled *bool `json:"memory_huge_pages_enable"`
	// 内存Qos
	MemoryQos *EcsMemQos `json:"memory_qos"`
	// 是否开启内存预留
	MemoryReserve *bool `json:"mem_reserve" `
	// 是否开启内存自动扩展
	MemoryAutoExpand *bool `json:"mem_auto_expand"`
	// 内存自动扩展值
	MemoryExpandNum *uint64 `json:"mem_expand_num"`
}

type EcsMemQos struct {
	// 虚拟机实际能使用的最大内存
	HardLimit uint64 `json:"hard_limit" `
	// 资源紧张的时候，虚拟机能争取到的最大内存
	SoftLimit uint64 `json:"soft_limit"`
	// 虚拟机能使用的内存加上交换区的内存，不能小于HardLimit
	SwapHardLimit uint64 `json:"swap_hard_limit"`
}

type EcsDiskSpecial struct {
	Name string `json:"name"`
	// 镜像仓库uuid
	RepositoryUUID string `json:"repository_uuid"`
	// 镜像uuid
	ImageUUID string `json:"image_uuid"`
}
