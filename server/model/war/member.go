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
	Gender   *int   `json:"gender" form:"gender" gorm:"column:gender;comment:1男，2 女;default:0;"`
	Height   *int   `json:"height" form:"height" gorm:"column:height;comment:单位 kg;default:0;"`
	Weight   *int   `json:"weight" form:"weight" gorm:"column:weight;comment:单位 cm;default:0;"`
	Phone    string `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
	Openid   string `json:"openId" form:"openId" gorm:"column:openid;comment:;"`
	Match    *uint  `json:"match" form:"match" gorm:"column:match;comment:比赛的剩余场次;default:0;"`
	//族别
	Clan string `json:"clan" form:"clan" gorm:"column:clan;comment:族别;default:汉族;"`
	//身份证号码
	IdCard string `json:"idCard" form:"idCard" gorm:"column:id_card;comment:身份证号码;default:null;"`
	//荣誉分
	Honor int `json:"honor" form:"honor" gorm:"column:honor;comment:荣誉分;default:100;"`
	//积分
	Score int `json:"score" form:"score" gorm:"column:score;comment:战斗积分;default:0;"`
	//消费积分
	ConsumeScore int `json:"consumeScore" form:"consumeScore" gorm:"column:consume_score;comment:消费积分;default:0;"`
	//经验值
	Exp int `json:"exp" form:"exp" gorm:"column:exp;comment:经验值;default:0;"`
	//军衔等级
	RankLevelId uint    `json:"rankLevelId" form:"rankLevelId" gorm:"column:rank_level_id;comment:军衔等级;default:1;"`
	Kda         float64 `json:"kda" form:"kda" gorm:"column:kda;decimal(4,2);comment:kda;default:0;"`
	//军衔
	RankLevel RankLevel `json:"rankLevel" form:"rankLevel" gorm:"foreignKey:RankLevelId;references:ID;"`

	//会员等级
	MemberLevelId uint `json:"memberLevelId" form:"memberLevelId" gorm:"column:member_level_id;comment:会员等级;default:11;"`
	//会员
	MemberLevel MemberLevel `json:"memberLevel" form:"memberLevel" gorm:"foreignKey:MemberLevelId;references:ID;"`
	//所属战队ID
	TeamID   *uint  `json:"teamId" form:"teamId" gorm:"column:team_id;comment:所属战队ID;default:0;"`
	TeamName string `json:"teamName" form:"teamName" gorm:"column:team_name;comment:所属战队名称;default:null;"`
}

// TableName Member 表名
func (Member) TableName() string {
	return "war_member"
}
