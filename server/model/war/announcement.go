// 自动生成模板Announcement
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Announcement 结构体
type Announcement struct {
      global.GVA_MODEL
      Content  string `json:"content" form:"content" gorm:"column:content;comment:;size:500;"`
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:;"`
}


// TableName Announcement 表名
func (Announcement) TableName() string {
  return "war_announcement"
}

