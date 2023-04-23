// 自动生成模板MatchRecord
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MatchRecord 结构体
type MatchRecord struct {
	global.GVA_MODEL
	UserId    uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:;default:0;"`
	MatchType uint   `json:"matchType" form:"matchType" gorm:"column:match_type;comment:场类型;default:1;"`
	MatchNum  uint   `json:"matchNum" form:"matchNum" gorm:"column:match_num;comment:增加或减少的场次;default:0;"`
	Remark    string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
}

// TableName MatchRecord 表名
func (MatchRecord) TableName() string {
	return "war_match_record"
}
