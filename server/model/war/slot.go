// 自动生成模板Slot
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Slot 结构体
type Slot struct {
      global.GVA_MODEL
      Times  string `json:"times" form:"times" gorm:"column:times;comment:可选时间;"`
}


// TableName Slot 表名
func (Slot) TableName() string {
  return "war_time_slot"
}

