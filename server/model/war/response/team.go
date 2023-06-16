package response

// 战队邀请海报
type TeamInvitePoster struct {
	//战队名称
	TeamName string `json:"teamName"`
	//人数
	Num int `json:"num"`
	//战队logo
	TeamLogo string `json:"teamLogo"`
	//战队二维码
	TeamQrCode string `json:"teamQrCode"`
}

// 战队信息
type TeamInfo struct {
	//战队logo
	Logo string `json:"logo"`
	//战队名称
	Name string `json:"name"`
	//简介
	Description string `json:"description"`
	//战队人数
	TeamMemberNum int              `json:"teamMemberNum"`
	TeamMember    []TeamMemberInfo `json:"memberInfo"`
}

// 战队成员信息
type TeamMemberInfo struct {
	MemberInfo
	RoleName string `json:"roleName"`
	//装备
	Equipments []EquipmentIcon `json:"equipments"`
}
