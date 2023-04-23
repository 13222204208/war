// 自动生成模板TeamApply
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TeamApply 结构体
type TeamApply struct {
	global.GVA_MODEL
	UserId uint `json:"userId" form:"userId" gorm:"column:user_id;comment:;size:11;"`
	TeamId uint `json:"teamId" form:"teamId" gorm:"column:team_id;comment:;size:10;"`
	Status *int `json:"status" form:"status" gorm:"column:status;comment:;"`
}

// TableName TeamApply 表名
func (TeamApply) TableName() string {
	return "war_team_apply"
}
