package iexample

type AuthRequest struct {
	User     string
	Tenant   string
	Password string
}

// LoginResponse 用户登录回应
type AuthResponse struct {
	ResponseResult
	Data struct {
		Role             string
		Token            string
		UserUuid         string `json:"user_uuid"`
		UserName         string `json:"user_name"`
		EnableLDAP       int    `json:"enable_ldap"`
		Tenant           string
		SystemMemberList struct {
			Leader     string
			Candidates []string `json:"candidates"`
		} `json:"system_member_list"`
		DefaultCluster string `json:"default_cluster_uuid"`
	}
}
