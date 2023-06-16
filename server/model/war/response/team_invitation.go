package response

// 战队受邀列表
type TeamInvitationResponse struct {
	//邀请id
	Id uint `json:"id"`
	//战队名称
	TeamName string `json:"teamName"`
	//战队logo
	TeamLogo string `json:"teamLogo"`
	//对局时间
	MatchTime string `json:"matchTime"`
	//战队长
	Captain string `json:"captain"`
	//战队人数
	TeamMemberCount int `json:"teamMemberCount"`
	//简介
	Introduction string `json:"introduction"`
	//状态
	Status int `json:"status"`
}
