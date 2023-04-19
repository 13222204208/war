package war

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/copilot"
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
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&members).Error
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
