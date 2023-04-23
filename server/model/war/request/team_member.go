package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
)

type TeamMemberSearch struct {
	war.TeamMember
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

// 更改会员的角色
type UpdateTeamMemberRole struct {
	TeamMemberId uint `json:"teamMemberId"`
	RoleId       uint `json:"roleId"`
}
