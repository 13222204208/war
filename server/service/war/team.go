package war

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
	"go.uber.org/zap"
)

type TeamService struct {
}

// CreateTeam 创建Team记录
// Author [piexlmax](https://github.com/piexlmax)
// func (teamService *TeamService) CreateTeam(team *war.Team) (err error) {
// 	err = global.GVA_DB.Create(team).Error
// 	return err
// }

// DeleteTeam 删除Team记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamService *TeamService) DeleteTeam(team war.Team) (err error) {
	err = global.GVA_DB.Delete(&team).Error
	return err
}

// DeleteTeamByIds 批量删除Team记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamService *TeamService) DeleteTeamByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.Team{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTeam 更新Team记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamService *TeamService) UpdateTeam(team war.Team) (err error) {
	err = global.GVA_DB.Save(&team).Error
	return err
}

// GetTeam 根据id获取Team记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamService *TeamService) GetTeam(id uint) (team war.Team, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&team).Error
	return
}

// GetTeamInfoList 分页获取Team记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamService *TeamService) GetTeamInfoList(info warReq.TeamSearch) (list []war.Team, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.Team{})
	var teams []war.Team
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&teams).Error
	return teams, total, err
}

// 创建战队
func (teamService *TeamService) CreateTeam(team *war.Team) (err error) {
	match := global.GVA_CONFIG.Team.Create //获取战队创建所需的场次
	global.GVA_LOG.Info("创建战队所需的场次", zap.Any("match", match))
	ok := IsUserMatchFull(*team.LeaderId, match)
	if !ok {
		return errors.New("场次不足")
	}
	if ok := IsTeamNameExist(team.Name); ok {
		return errors.New("战队名已经存在")
	}
	//创建事务
	tx := global.GVA_DB.Begin()
	err = tx.Create(&team).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var teamMember war.TeamMember
	teamMember.TeamId = team.ID
	teamMember.UserId = *team.LeaderId
	teamMember.TeamRoleId = 1
	err = tx.Create(&teamMember).Error //保存队长到战队成员表
	if err != nil {
		tx.Rollback()
		return err
	}
	err = DeductUserMatch(*team.LeaderId, match, "创建战队扣除场次")
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// 修改战队信息
func (teamService *TeamService) UpdateTeamInfo(team *war.Team) (err error) {
	//先根据leaderId查询战队信息
	var t war.Team
	err = global.GVA_DB.Where("leader_id = ?", team.LeaderId).First(&t).Error
	if err != nil {
		return err
	}
	if team.Name != t.Name {
		//判断战队笱是否存在，如果存在则提示已经存在
		if ok := IsTeamNameExist(team.Name); ok {
			return errors.New("战队名已经存在")
		}
		//扣除修改战队名的场次
		err = DeductUserMatch(*team.LeaderId, global.GVA_CONFIG.Team.Update, "修改战队名扣除场次")
		if err != nil {
			return err
		}
	}
	err = global.GVA_DB.Model(&t).Updates(team).Error
	return err
}

// 获取全险战队信息，显示队长信息 ，并按战队积分从大到小排序
func (teamService *TeamService) GetAllTeam() (list []war.Team, err error) {
	err = global.GVA_DB.Preload("LeaderInfo").Order("score DESC").Find(&list).Error
	return list, err
}

// 获取战队详情
func (teamService *TeamService) GetTeamDetail(teamId string) (team war.Team, err error) {
	err = global.GVA_DB.Preload("TeamMember").Preload("TeamMember.MemberInfo").Preload("TeamMember.TeamRoleInfo").Where("id = ?", teamId).First(&team).Error
	team.TeamMemberNum = len(team.TeamMember)
	return team, err
}

// 获取我的战队详情
func (teamService *TeamService) GetMyTeamDetail(userId uint) (team war.Team, err error) {
	err = global.GVA_DB.Preload("TeamMember").Preload("TeamMember.MemberInfo").Preload("TeamMember.TeamRoleInfo").Where("leader_id = ?", userId).First(&team).Error
	team.TeamMemberNum = len(team.TeamMember)
	return team, err
}

// 判断战队名称是否已经存在
// 存在返回true
func IsTeamNameExist(name string) bool {
	var count int64
	global.GVA_DB.Model(&war.Team{}).Where("name = ?", name).Count(&count)
	return count > 0
}

// 判断用户的场次是否已经满足
// 如果满足则返回true
// userId 用户id
// match 场次
func IsUserMatchFull(userId, match uint) bool {
	var m war.Member
	err := global.GVA_DB.Where("id = ?", userId).First(&m).Error
	if err != nil {
		global.GVA_LOG.Error("查询用户信息失败", zap.Error(err))
	}
	count := *m.Match
	return count >= match
}

// 扣除用户的场次
func DeductUserMatch(userId, match uint, remark string) (err error) {
	var m war.Member
	err = global.GVA_DB.Where("id = ?", userId).First(&m).Error
	if err != nil {
		return err
	}
	count := *m.Match
	if count >= match {
		//count减match 扣除场次
		count -= match
		m.Match = &count
		err = global.GVA_DB.Save(&m).Error
		if err != nil {
			return err
		}
	} else {
		return errors.New("场次不足")
	}
	//保存扣除场次的记录
	var record war.MatchRecord
	record.UserId = userId
	record.MatchNum = match
	record.MatchType = 2
	record.Remark = remark
	err = global.GVA_DB.Create(&record).Error
	return err
}
