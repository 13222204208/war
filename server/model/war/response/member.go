package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/war"

type LoginResponse struct {
	User      war.Member `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}

// 个人详情
type MemberResponse struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	//战损比
	DamageRatio float64 `json:"damageRatio"`
	//战损排行
	WinRateRank int `json:"winRateRank"`
	//战队名称
	TeamName string `json:"teamName"`
	//战队logo
	TeamLogo string `json:"teamLogo"`
	//战队职位
	Role string `json:"Role"`
	//装备
	Equipments []map[string]interface{} `json:"equipments"`
}

// 装备详情
type EquipmentResponse struct {
	//装备分类名称
	CategoryName string `json:"categoryName"`
	//装备名称
	Name string `json:"name"`
	//装备图标
	Icon string `json:"icon"`
}

// 我的kda信息
type MyKdaResponse struct {
	//更新时间
	UpdateTime string `json:"updateTime"`
	//kda
	Kda float64 `json:"kda"`
	//胜局
	Win int `json:"win"`
	//败局
	Lose int `json:"lose"`
}

// 我的战斗信息
type MyBattleResponse struct {
	//场次
	Count int `json:"count"`
	//胜率
	Kda float64 `json:"kda"`
	//胜局
	Win int `json:"win"`
	//败局
	Lose int `json:"lose"`
}

// 用户排行
type MemberRankResponse struct {
	UserId uint `json:"userId"`
	//排名
	Rank int `json:"rank"`
	//头像
	Avatar string `json:"avatar"`
	//昵称
	Nickname string `json:"nickname"`
	//kda
	Kda float64 `json:"kda"`
	//积分
	Score int `json:"score"`
	//战队名称
	TeamName string `json:"teamName"`
	//战队logo
	TeamLogo string `json:"teamLogo"`
	//装备
	Equipments MyEquipmentList `json:"equipments"`
}

// 装备图标
type EquipmentIcon struct {
	//装备图标
	Icon string `json:"icon"`
}
