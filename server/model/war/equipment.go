// 自动生成模板Equipment
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Equipment 结构体
type Equipment struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:装备名称;"`
      Icon  string `json:"icon" form:"icon" gorm:"column:icon;comment:图标地址;"`
      ParentId  *int `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;size:5;"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:;size:3;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:;"`
}


// TableName Equipment 表名
func (Equipment) TableName() string {
  return "war_equipment"
}

