// 自动生成模板Team
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Team 结构体
type Team struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Logo        string `json:"logo" form:"logo" gorm:"column:logo;comment:;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:简介;"`
	Score       *int   `json:"score" form:"score" gorm:"column:score;comment:战队积分;"`
	LeaderId    *uint  `json:"leaderId" form:"leaderId" gorm:"column:leader_id;unique;comment:队长的ID;"`
	//队长信息
	LeaderInfo *Member `json:"leaderInfo" form:"leaderInfo" gorm:"foreignKey:LeaderId;references:ID;comment:队长信息;"`
	//战队人数
	TeamMemberNum int `json:"teamMemberNum" form:"teamMemberNum" gorm:"column:team_member_num;comment:战队人数;default:0"`
	//战队成员
	TeamMember []*TeamMember `json:"teamMember" form:"teamMember" gorm:"foreignKey:TeamId;references:ID;comment:战队成员;"`
}

// TableName Team 表名
func (Team) TableName() string {
	return "war_team"
}
