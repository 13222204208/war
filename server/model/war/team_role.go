// 自动生成模板TeamRole
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// TeamRole 结构体
type TeamRole struct {
      global.GVA_MODEL
      Role  string `json:"role" form:"role" gorm:"column:role;comment:角色名称;"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:;"`
}


// TableName TeamRole 表名
func (TeamRole) TableName() string {
  return "war_team_role"
}

