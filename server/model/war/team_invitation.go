// 自动生成模板TeamInvitation
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TeamInvitation 结构体
type TeamInvitation struct {
	global.GVA_MODEL
	TeamId uint `json:"teamId" form:"teamId" gorm:"column:team_id;comment:邀请战队的ID;"`
	//战队信息
	Team Team `json:"team" form:"team" gorm:"foreignKey:TeamId;references:ID;comment:邀请战队的信息;"`

	TeamCaptainId uint `json:"teamCaptainId" form:"teamCaptainId" gorm:"column:team_captain_id;comment:邀请战队长ID;"`
	//战队长信息
	TeamCaptain Member `json:"teamCaptain" form:"teamCaptain" gorm:"foreignKey:TeamCaptainId;references:ID;comment:邀请战队长的信息;"`

	Date                 string `json:"date" form:"date" gorm:"column:date;comment:预约的战队赛时间;"`
	InvitedTeamId        uint   `json:"invitedTeamId" form:"invitedTeamId" gorm:"column:invited_team_id;comment:被邀请的战队ID;"`
	InvitedTeamCaptainId uint   `json:"invitedTeamCaptainId" form:"invitedTeamCaptainId" gorm:"column:invited_team_captain_id;comment:被邀请战队的战队长ID;"`
	Status               int    `json:"status" form:"status" gorm:"column:status;comment:被邀请战队状态 0 待审批 1 通过，2拒绝;default:0"`
}

// TableName TeamInvitation 表名
func (TeamInvitation) TableName() string {
	return "war_team_invitation"
}
