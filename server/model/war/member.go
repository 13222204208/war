// 自动生成模板Member
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Member 结构体
type Member struct {
	global.GVA_MODEL
	Avatar   string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:;"`
	Nickname string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:昵称代号;"`
	Name     string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Gender   *int   `json:"gender" form:"gender" gorm:"column:gender;comment:1男，2 女;"`
	Height   *int   `json:"height" form:"height" gorm:"column:height;comment:单位 kg;"`
	Weight   *int   `json:"weight" form:"weight" gorm:"column:weight;comment:单位 cm;"`
	Phone    string `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
	Openid   string `json:"openId" form:"openId" gorm:"column:openid;comment:;"`
	Match    *uint  `json:"match" form:"match" gorm:"column:match;comment:比赛的剩余场次;default:0;"`
	//经验值
	Exp int `json:"exp" form:"exp" gorm:"column:exp;comment:经验值;default:0;"`
	//军衔等级
	RankLevelId uint `json:"rankLevelId" form:"rankLevelId" gorm:"column:rank_level_id;comment:军衔等级;default:0;"`
	//军衔
	RankLevel RankLevel `json:"rankLevel" form:"rankLevel" gorm:"foreignKey:RankLevelId;references:ID;"`

	//会员等级
	MemberLevelId uint `json:"memberLevelId" form:"memberLevelId" gorm:"column:member_level_id;comment:会员等级;default:1;"`
	//会员
	MemberLevel MemberLevel `json:"memberLevel" form:"memberLevel" gorm:"foreignKey:MemberLevelId;references:ID;"`
	//所属战队ID
	TeamID   *uint  `json:"teamId" form:"teamId" gorm:"column:team_id;comment:所属战队ID;default:0;"`
	TeamName string `json:"teamName" form:"teamName"`
}

// TableName Member 表名
func (Member) TableName() string {
	return "war_member"
}
