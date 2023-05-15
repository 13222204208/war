package war

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
	warRes "github.com/flipped-aurora/gin-vue-admin/server/model/war/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/copilot"
	"go.uber.org/zap"
)

type MemberService struct {
}

// CreateMember 创建Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) CreateMember(member *war.Member) (err error) {
	err = global.GVA_DB.Create(member).Error
	return err
}

// DeleteMember 删除Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) DeleteMember(member war.Member) (err error) {
	err = global.GVA_DB.Delete(&member).Error
	return err
}

// DeleteMemberByIds 批量删除Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) DeleteMemberByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.Member{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMember 更新Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) UpdateMember(member war.Member) (err error) {
	err = global.GVA_DB.Save(&member).Error
	return err
}

// GetMember 根据id获取Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) GetMember(id uint) (member war.Member, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&member).Error
	return
}

// GetMemberInfoList 分页获取Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) GetMemberInfoList(info warReq.MemberSearch) (list []war.Member, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.Member{})
	var members []war.Member
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	//如果代号不为空
	if info.Nickname != "" {
		db = db.Where("nickname = ?", info.Nickname)
	}
	//如果姓名不为空
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// err = db.Limit(limit).Offset(offset).Find(&members).Error
	err = db.Table("war_member").Select("war_member.*, war_team.name as team_name").Joins("left join war_team on war_member.team_id = war_team.id").Preload("MemberLevel").Limit(limit).Offset(offset).Find(&members).Error
	return members, total, err
}

// 登陆
func (memberService *MemberService) Login(code string) (member *war.Member, err error) {
	openid, err := copilot.GetOpenId(code)
	if err != nil {
		return
	}
	if openid == "" {
		return nil, errors.New("openid is empty")
	}

	global.GVA_DB.Where("openid = ?", openid).First(&member)

	if member.ID == 0 {
		member.Openid = openid
		err = global.GVA_DB.Create(&member).Error
		if err != nil {
			return
		}
	}
	return member, err
}

// 会员修改信息
func (memberService *MemberService) UpdateMemberInfo(userID uint, member war.Member) (err error) {
	err = global.GVA_DB.Model(&war.Member{}).Where("id = ?", userID).Updates(member).Error
	return err
}

// 获取会员资料
func (memberService *MemberService) GetMemberInfo(userID uint) (member war.Member, err error) {
	err = global.GVA_DB.Where("id = ?", userID).First(&member).Error
	return
}

// 会员增加或减少场次
func (memberService *MemberService) AddOrUpdateMemberMatch(userID, match, matchType uint) (err error) {
	global.GVA_LOG.Info("场次类型", zap.Any("matchType", matchType))
	if matchType == 1 {
		err = AddUserMatch(userID, match, "后台增加场次")
		if err != nil {
			return err
		}
	} else if matchType == 2 {
		err = DeductUserMatch(userID, match, "后台减少场次")
		if err != nil {
			return err
		}
	} else {
		return errors.New("场次类型必须为1或者2")
	}
	return err
}

// 会员增加场次
func AddUserMatch(userID, match uint, remark string) (err error) {
	var member war.Member
	err = global.GVA_DB.Where("id = ?", userID).First(&member).Error
	if err != nil {
		return err
	} else {
		*member.Match += match
		err = global.GVA_DB.Save(&member).Error
		if err != nil {
			return err
		}
	}
	var record war.MatchRecord
	record.UserId = userID
	record.MatchNum = match
	record.MatchType = 1
	record.Remark = remark
	err = global.GVA_DB.Create(&record).Error
	return err
}

// 个人详情
func (memberService *MemberService) GetMemberDetail(userID uint) (response warRes.MemberResponse, err error) {
	var tm war.TeamMember
	err = global.GVA_DB.Where("user_id = ?", userID).Preload("MemberInfo").Preload("TeamRoleInfo").Preload("TeamInfo").First(&tm).Error
	if err != nil {
		return response, err
	}
	response.Avatar = tm.MemberInfo.Avatar
	response.Nickname = tm.MemberInfo.Nickname
	response.Role = tm.TeamRoleInfo.Role
	response.TeamName = tm.TeamInfo.Name
	response.TeamLogo = tm.TeamInfo.Logo
	response.DamageRatio = tm.DamageRatio
	response.WinRateRank = 3
	//获取个人装备
	response.Equipments, err = GetMemberEquip(userID)
	return response, err
}

// 查询个人装备
func GetMemberEquip(userID uint) (m []map[string]interface{}, err error) {
	var response []warRes.EquipmentResponse
	//查询用户装备
	var userEquip []war.UserEquipment
	err = global.GVA_DB.Where("user_id = ?", userID).Preload("Category").Preload("Equipment").Find(&userEquip).Error
	if err != nil {
		return m, err
	}
	for _, v := range userEquip {
		var equip warRes.EquipmentResponse
		equip.Name = v.Equipment.Name
		equip.CategoryName = v.Category.Name
		equip.Icon = v.Equipment.Icon
		response = append(response, equip)
	}
	//装备分类分组
	m, err = GetEquipCategory(response)
	return m, err
}

// 装备分类分组
func GetEquipCategory(response []warRes.EquipmentResponse) (m []map[string]interface{}, err error) {

	categories := make(map[string][]map[string]string)

	for _, equipment := range response {
		category := equipment.CategoryName
		item := map[string]string{
			"name": equipment.Name,
			"icon": equipment.Icon,
		}

		if _, ok := categories[category]; !ok {
			categories[category] = []map[string]string{}
		}

		categories[category] = append(categories[category], item)
	}

	result := []map[string]interface{}{}

	for categoryName, children := range categories {
		category := map[string]interface{}{
			"categoryName": categoryName,
			"children":     children,
		}

		result = append(result, category)
	}
	return result, err
}
