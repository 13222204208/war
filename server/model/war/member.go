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
}

// TableName Member 表名
func (Member) TableName() string {
	return "war_member"
}