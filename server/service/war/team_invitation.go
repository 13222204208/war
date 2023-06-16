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

type TeamInvitationService struct {
}

// CreateTeamInvitation 创建TeamInvitation记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamInvitationService *TeamInvitationService) CreateTeamInvitation(teamInvitation *war.TeamInvitation) (err error) {
	err = global.GVA_DB.Create(teamInvitation).Error
	return err
}

// DeleteTeamInvitation 删除TeamInvitation记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamInvitationService *TeamInvitationService) DeleteTeamInvitation(teamInvitation war.TeamInvitation) (err error) {
	err = global.GVA_DB.Delete(&teamInvitation).Error
	return err
}

// DeleteTeamInvitationByIds 批量删除TeamInvitation记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamInvitationService *TeamInvitationService) DeleteTeamInvitationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.TeamInvitation{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTeamInvitation 更新TeamInvitation记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamInvitationService *TeamInvitationService) UpdateTeamInvitation(teamInvitation war.TeamInvitation) (err error) {
	err = global.GVA_DB.Save(&teamInvitation).Error
	return err
}

// GetTeamInvitation 根据id获取TeamInvitation记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamInvitationService *TeamInvitationService) GetTeamInvitation(id uint) (teamInvitation war.TeamInvitation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&teamInvitation).Error
	return
}

// GetTeamInvitationInfoList 分页获取TeamInvitation记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamInvitationService *TeamInvitationService) GetTeamInvitationInfoList(info warReq.TeamInvitationSearch) (list []war.TeamInvitation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.TeamInvitation{})
	var teamInvitations []war.TeamInvitation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&teamInvitations).Error
	return teamInvitations, total, err
}

// 战队邀请
func (teamInvitationService *TeamInvitationService) TeamInvitation(teamInvitation war.TeamInvitation) (err error) {
	//判断是否已经邀请过
	var teamInvitationInfo war.TeamInvitation
	err = global.GVA_DB.Where("team_id = ? AND invited_team_id = ?", teamInvitation.TeamId, teamInvitation.InvitedTeamId).First(&teamInvitationInfo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//判断邀请的战队和战队长是否存在
			var teamInfo war.Team
			err = global.GVA_DB.Where("leader_id = ?", teamInvitation.TeamCaptainId).First(&teamInfo).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return errors.New("邀请的战队不存在")
				} else {
					return err
				}
			}
			//战队ID
			teamInvitation.TeamId = teamInfo.ID
			// 邀请
			err = global.GVA_DB.Create(&teamInvitation).Error
		} else {
			// 其他错误
			return err
		}
	} else {
		return errors.New("已经邀请过了")
	}
	return
}

// 战队受邀请列表
func (teamInvitationService *TeamInvitationService) TeamInvitationList(userId uint) (response []warRes.TeamInvitationResponse, err error) {
	var teamInvitationList []war.TeamInvitation
	err = global.GVA_DB.Where("invited_team_captain_id = ?", userId).Preload("Team").Preload("TeamCaptain").Find(&teamInvitationList).Error
	if err != nil {
		return
	}
	if len(teamInvitationList) == 0 {
		return
	}
	// 转换数据
	for _, item := range teamInvitationList {
		response = append(response, warRes.TeamInvitationResponse{
			Id:              item.ID,
			TeamName:        item.Team.Name,
			TeamLogo:        item.Team.Logo,
			MatchTime:       item.Date,
			Captain:         item.TeamCaptain.Nickname,
			TeamMemberCount: item.Team.TeamMemberNum,
			Introduction:    item.Team.Description,
			Status:          item.Status,
		})
	}
	return
}
