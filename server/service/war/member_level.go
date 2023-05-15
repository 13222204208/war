package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type MemberLevelService struct {
}

// CreateMemberLevel 创建MemberLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberLevelService *MemberLevelService) CreateMemberLevel(memberLevel *war.MemberLevel) (err error) {
	err = global.GVA_DB.Create(memberLevel).Error
	return err
}

// DeleteMemberLevel 删除MemberLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberLevelService *MemberLevelService)DeleteMemberLevel(memberLevel war.MemberLevel) (err error) {
	err = global.GVA_DB.Delete(&memberLevel).Error
	return err
}

// DeleteMemberLevelByIds 批量删除MemberLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberLevelService *MemberLevelService)DeleteMemberLevelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.MemberLevel{},"id in ?",ids.Ids).Error
	return err
}

// UpdateMemberLevel 更新MemberLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberLevelService *MemberLevelService)UpdateMemberLevel(memberLevel war.MemberLevel) (err error) {
	err = global.GVA_DB.Save(&memberLevel).Error
	return err
}

// GetMemberLevel 根据id获取MemberLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberLevelService *MemberLevelService)GetMemberLevel(id uint) (memberLevel war.MemberLevel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&memberLevel).Error
	return
}

// GetMemberLevelInfoList 分页获取MemberLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberLevelService *MemberLevelService)GetMemberLevelInfoList(info warReq.MemberLevelSearch) (list []war.MemberLevel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&war.MemberLevel{})
    var memberLevels []war.MemberLevel
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&memberLevels).Error
	return  memberLevels, total, err
}
