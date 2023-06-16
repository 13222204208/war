package war

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
	warRes "github.com/flipped-aurora/gin-vue-admin/server/model/war/response"

	"gorm.io/gorm"
)

type TeamApplyService struct {
}

// CreateTeamApply 创建TeamApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamApplyService *TeamApplyService) CreateTeamApply(teamApply *war.TeamApply) (err error) {
	err = global.GVA_DB.Create(teamApply).Error
	return err
}

// DeleteTeamApply 删除TeamApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamApplyService *TeamApplyService) DeleteTeamApply(teamApply war.TeamApply) (err error) {
	err = global.GVA_DB.Delete(&teamApply).Error
	return err
}

// DeleteTeamApplyByIds 批量删除TeamApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamApplyService *TeamApplyService) DeleteTeamApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.TeamApply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTeamApply 更新TeamApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamApplyService *TeamApplyService) UpdateTeamApply(teamApply war.TeamApply) (err error) {
	err = global.GVA_DB.Save(&teamApply).Error
	return err
}

// GetTeamApply 根据id获取TeamApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamApplyService *TeamApplyService) GetTeamApply(id uint) (teamApply war.TeamApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&teamApply).Error
	return
}

// GetTeamApplyInfoList 分页获取TeamApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamApplyService *TeamApplyService) GetTeamApplyInfoList(info warReq.TeamApplySearch) (list []war.TeamApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.TeamApply{})
	var teamApplys []war.TeamApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&teamApplys).Error
	return teamApplys, total, err
}

// 保存战队申请的记录
func (teamApplyService *TeamApplyService) SaveTeamApply(teamApply war.TeamApply) (err error) {
	//判断申请用户是否已经申请过和是否已经加入了战队
	var team war.TeamMember
	err = global.GVA_DB.Where("user_id = ?", teamApply.UserId).First(&team).Error
	if err == nil {
		return errors.New("该用户已经加入了战队")
	}
	//判断用户是否已经申请过这个战队
	var teamApplys war.TeamApply
	err = global.GVA_DB.Where("user_id = ? and team_id = ?", teamApply.UserId, teamApply.TeamId).First(&teamApplys).Error
	if err == nil {
		return errors.New("该用户已经申请过这个战队")
	}
	err = global.GVA_DB.Save(&teamApply).Error
	return err
}

// 申请加入战队的列表
func (teamApplyService *TeamApplyService) GetTeamApplyList(userID uint) (res []warRes.WarApplyListResponse, err error) {
	//查询用户的身份是否是战队长或者副队长
	var teamMember war.TeamMember
	err = global.GVA_DB.Where("user_id = ?", userID).First(&teamMember).Error
	if err != nil {
		return nil, err
	}

	if teamMember.TeamRoleId != 1 && teamMember.TeamRoleId != 2 {
		return nil, errors.New("用户不是战队长或者副队长")
	}

	var list []war.TeamApply
	err = global.GVA_DB.Where("team_id = ? AND status = ?", teamMember.TeamId, 0).Preload("User").Find(&list).Error
	if err != nil {
		return nil, err
	}

	response := make([]warRes.WarApplyListResponse, 0, len(list))

	for _, r := range list {
		res := warRes.WarApplyListResponse{
			Id:       r.ID,
			UserId:   r.UserId,
			Nickname: r.User.Nickname,
			Avatar:   r.User.Avatar,
			Kda:      r.User.Kda,
			Status:   *r.Status,
		}

		var userEquip []war.UserEquipment
		err := global.GVA_DB.Where("user_id = ?", r.UserId).Preload("Equipment").Find(&userEquip).Error
		if err != nil {
			return nil, err
		}

		for _, v := range userEquip {
			equip := warRes.EquipmentIcon{
				Icon: v.Equipment.Icon,
			}
			res.Equipments = append(res.Equipments, equip)
		}

		response = append(response, res)
	}

	return response, nil
}

// 同意或拒绝入队申请
func (teamApplyService *TeamApplyService) ApprovalTeamApply(userId, teamApplyId, status int) (err error) {
	// 查询用户的身份是否是战队长或者副队长
	var teamMember war.TeamMember
	err = global.GVA_DB.Where("user_id = ?", userId).First(&teamMember).Error
	if err != nil {
		return err
	}
	if teamMember.TeamRoleId != 1 && teamMember.TeamRoleId != 2 {
		return errors.New("用户不是战队长或者副队长")
	}

	var teamApply war.TeamApply
	err = global.GVA_DB.Where("id = ?", teamApplyId).First(&teamApply).Error
	if err != nil {
		return err
	}
	if *teamApply.Status != 0 {
		return errors.New("该申请已经处理过了")
	}

	if status == 1 {
		// 添加用户到战队成员表
		teamMember := war.TeamMember{
			TeamId: teamApply.TeamId,
			UserId: teamApply.UserId,
		}
		err = global.GVA_DB.Create(&teamMember).Error
		if err != nil {
			return err
		}

		err = UpdateTeamMember(teamApply)
		if err != nil {
			return err
		}
	}

	teamApply.Status = &status
	err = global.GVA_DB.Save(&teamApply).Error
	return err
}

// 修改会员的战队id和战队名称
func UpdateTeamMember(teamApply war.TeamApply) (err error) {
	var m war.Member
	err = global.GVA_DB.Where("id = ?", teamApply.UserId).First(&m).Error
	if err != nil {
		return errors.New("该用户不存在")
	}

	m.TeamID = &teamApply.TeamId

	var team war.Team
	//战队人数自增1
	global.GVA_DB.Model(&team).Where("id = ?", teamApply.TeamId).Update("team_member_num", gorm.Expr("team_member_num + ?", 1))
	err = global.GVA_DB.Where("id = ?", teamApply.TeamId).First(&team).Error
	if err != nil {
		return errors.New("该战队不存在")
	}
	m.TeamName = team.Name

	err = global.GVA_DB.Save(&m).Error
	return err
}
