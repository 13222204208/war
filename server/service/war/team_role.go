package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type TeamRoleService struct {
}

// CreateTeamRole 创建TeamRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamRoleService *TeamRoleService) CreateTeamRole(teamRole *war.TeamRole) (err error) {
	err = global.GVA_DB.Create(teamRole).Error
	return err
}

// DeleteTeamRole 删除TeamRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamRoleService *TeamRoleService) DeleteTeamRole(teamRole war.TeamRole) (err error) {
	err = global.GVA_DB.Delete(&teamRole).Error
	return err
}

// DeleteTeamRoleByIds 批量删除TeamRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamRoleService *TeamRoleService) DeleteTeamRoleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.TeamRole{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTeamRole 更新TeamRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamRoleService *TeamRoleService) UpdateTeamRole(teamRole war.TeamRole) (err error) {
	err = global.GVA_DB.Save(&teamRole).Error
	return err
}

// GetTeamRole 根据id获取TeamRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamRoleService *TeamRoleService) GetTeamRole(id uint) (teamRole war.TeamRole, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&teamRole).Error
	return
}

// GetTeamRoleInfoList 分页获取TeamRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamRoleService *TeamRoleService) GetTeamRoleInfoList(info warReq.TeamRoleSearch) (list []war.TeamRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.TeamRole{})
	var teamRoles []war.TeamRole
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&teamRoles).Error
	return teamRoles, total, err
}

// 获取战队角色列表
func (teamRoleService *TeamRoleService) GetTeamRoleList() (list []war.TeamRole, err error) {
	err = global.GVA_DB.Where("status = ?", 1).Find(&list).Error
	return
}
