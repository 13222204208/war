package war

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type TeamMemberService struct {
}

// CreateTeamMember 创建TeamMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamMemberService *TeamMemberService) CreateTeamMember(teamMember *war.TeamMember) (err error) {
	err = global.GVA_DB.Create(teamMember).Error
	return err
}

// DeleteTeamMember 删除TeamMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamMemberService *TeamMemberService) DeleteTeamMember(teamMember war.TeamMember) (err error) {
	err = global.GVA_DB.Delete(&teamMember).Error
	return err
}

// DeleteTeamMemberByIds 批量删除TeamMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamMemberService *TeamMemberService) DeleteTeamMemberByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.TeamMember{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTeamMember 更新TeamMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamMemberService *TeamMemberService) UpdateTeamMember(teamMember war.TeamMember) (err error) {
	err = global.GVA_DB.Save(&teamMember).Error
	return err
}

// GetTeamMember 根据id获取TeamMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamMemberService *TeamMemberService) GetTeamMember(id uint) (teamMember war.TeamMember, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&teamMember).Error
	return
}

// GetTeamMemberInfoList 分页获取TeamMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamMemberService *TeamMemberService) GetTeamMemberInfoList(info warReq.TeamMemberSearch) (list []war.TeamMember, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.TeamMember{})
	var teamMembers []war.TeamMember
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TeamRoleId != 0 {
		db = db.Where("team_role_id = ?", info.TeamRoleId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&teamMembers).Error
	return teamMembers, total, err
}

// 更改会员的角色
func (teamMemberService *TeamMemberService) UpdateTeamMemberRole(userID uint, t warReq.UpdateTeamMemberRole) (err error) {
	//判断userID 是不是队长
	var team war.Team
	err = global.GVA_DB.Where("leader_id = ?", userID).First(&team).Error
	if err != nil {
		return err
	}
	if team.ID == 0 {
		return errors.New("您不是队长，无法更改队员角色")
	}

	err = global.GVA_DB.Model(&war.TeamMember{}).Where("id = ?", t.TeamMemberId).Update("team_role_id", t.RoleId).Error
	return err
}
