package iexample

type VirtualSwitchListRequest struct {
	PagingParam
	Cluster string `json:"cluster_uuid"`
	Tenant  string `json:"tenant"`
	Zone    string `json:"zone_uuid"`
	// Flag 是否显示镜像网桥
	Flag       bool   `json:"flag"`
	FilterName string `json:"filter_name"`
	Types      string `json:"types"`
	HasVlan    int32  `json:"has_vlan"`
	FilterHost string `json:"filter_host"`
}

type VirtualSwitchListResponse struct {
	ResponseResult
	Data struct {
		TotalCount int                  `json:"total_count"`
		List       []VirtualSwitchEntry `json:"list"`
	}
}

type VirtualSwitchEntry struct {
	UUID string
	Name string
}

type VirtualSwitchGetRequest struct {
	Cluster string `json:"cluster_uuid"`
	Tenant  string `json:"tenant"`
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Zone    string `json:"zone"`
}

type VirtualSwitchGetResponse struct {
	ResponseResult
	Data *struct {
		Layer2Info    VirtualSwitchEntry
		PortGroupList []PortGroupEntry `json:"port_group_list"`
	}
}

type PortGroupEntry struct {
	PortGroupName  string      `json:"port_group_name"`
	PortGroupUUID  string      `json:"port_group_uuid"`
	CreateTime     int         `json:"create_time"`
	UpdateTime     int         `json:"update_time"`
	VlanID         int         `json:"vlan_id"`
	Layer2Name     string      `json:"layer2_name"`
	ExtIds         interface{} `json:"ext_ids"`
	AttachQuantity int         `json:"attach_quantity"`
	Bandwidth      interface{} `json:"bandwidth"`
	VM             interface{} `json:"vm"`
	Container      interface{} `json:"container"`
	ZoneUUID       string      `json:"zone_uuid"`
	ZoneName       string      `json:"zone_name"`
	L2NetworkUUID  string      `json:"l2network_uuid"`
	EnableDhcp     bool        `json:"enable_dhcp"`
	DhcpServer     interface{} `json:"dhcp_server"`
	EnableDhcpv6   bool        `json:"enable_dhcpv6"`
	Dhcpv6Server   interface{} `json:"dhcpv6_server"`
	VMTotal        int         `json:"vm_total"`
	ContainerTotal int         `json:"container_total"`
	NicInfos       interface{} `json:"nic_infos"`
	UsedIPCount    int         `json:"used_ip_count"`
	UsedIpv6Total  int         `json:"used_ipv6_total"`
}

type VirtualSwitchPortGroupGetRequest struct {
	Cluster     string `json:"cluster_uuid"`
	Tenant      string `json:"tenant"`
	VSwitchUUID string `json:"vswitch_uuid"`
	VSwitchName string `json:"vswitch_name"`
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
}

type VirtualSwitchPortGroupGetResponse struct {
	ResponseResult
	Data *PortGroupEntry
}
