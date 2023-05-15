// 自动生成模板RankLevel
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// RankLevel 结构体
type RankLevel struct {
	global.GVA_MODEL
	Name       string `json:"name" form:"name" gorm:"column:name;comment:军衔名称;"`
	Icon       string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;"`
	Experience int    `json:"experience" form:"experience" gorm:"column:experience;comment:升级军衔所需经验;"`
}

// TableName RankLevel 表名
func (RankLevel) TableName() string {
	return "war_rank_level"
}
