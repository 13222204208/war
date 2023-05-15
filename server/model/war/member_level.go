// 自动生成模板MemberLevel
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// MemberLevel 结构体
type MemberLevel struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:会员名称;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:会员等级状态;"`
}


// TableName MemberLevel 表名
func (MemberLevel) TableName() string {
  return "war_member_level"
}

