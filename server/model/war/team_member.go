// 自动生成模板TeamMember
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TeamMember 结构体
type TeamMember struct {
	global.GVA_MODEL
	UserId      uint    `json:"userId" form:"userId" gorm:"column:user_id;comment:;size:11;default:0;"`
	TeamId      uint    `json:"teamId" form:"teamId" gorm:"column:team_id;comment:;size:10;default:0;"`
	TeamRoleId  uint    `json:"teamRoleId" form:"teamRoleId" gorm:"column:team_role_id;comment:战队角色ID;size:4;default:0;"`
	DamageRatio float64 `json:"damageRatio" form:"damageRatio" gorm:"column:damage_ratio;comment:;size:10,2;default:0;"`
	//会员信息
	MemberInfo *Member `json:"memberInfo" form:"memberInfo" gorm:"foreignKey:UserId;references:ID;comment:会员信息;"`
	//会员角色信息
	TeamRoleInfo *TeamRole `json:"teamRoleInfo" form:"teamRoleInfo" gorm:"foreignKey:TeamRoleId;references:ID;comment:会员角色信息;"`
}

// TableName TeamMember 表名
func (TeamMember) TableName() string {
	return "war_team_member"
}
