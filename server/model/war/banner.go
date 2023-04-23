// 自动生成模板Banner
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Banner 结构体
type Banner struct {
      global.GVA_MODEL
      Url  string `json:"url" form:"url" gorm:"column:url;comment:;"`
      Path  string `json:"path" form:"path" gorm:"column:path;comment:;"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:;"`
}


// TableName Banner 表名
func (Banner) TableName() string {
  return "war_banner"
}

