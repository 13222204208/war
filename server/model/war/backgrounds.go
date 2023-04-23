// 自动生成模板Backgrounds
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Backgrounds 结构体
type Backgrounds struct {
      global.GVA_MODEL
      Url  string `json:"url" form:"url" gorm:"column:url;comment:图片的网络地址;"`
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:;"`
}


// TableName Backgrounds 表名
func (Backgrounds) TableName() string {
  return "war_backgrounds"
}

