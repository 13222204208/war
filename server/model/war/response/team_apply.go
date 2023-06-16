package response

// 入队申请列表信息
type WarApplyListResponse struct {
	//申请ID
	Id uint `json:"id"`
	//申请人ID
	UserId uint `json:"userId"`
	//申请人昵称
	Nickname string `json:"nickname"`
	//申请人头像
	Avatar string `json:"avatar"`
	//kda
	Kda float64 `json:"kda"`
	//当前状态
	Status int `json:"status"`
	//装备
	Equipments []EquipmentIcon `json:"equipment"`
}
