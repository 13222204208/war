package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type BackgroundsService struct {
}

// CreateBackgrounds 创建Backgrounds记录
// Author [piexlmax](https://github.com/piexlmax)
func (backgroundsService *BackgroundsService) CreateBackgrounds(backgrounds *war.Backgrounds) (err error) {
	err = global.GVA_DB.Create(backgrounds).Error
	return err
}

// DeleteBackgrounds 删除Backgrounds记录
// Author [piexlmax](https://github.com/piexlmax)
func (backgroundsService *BackgroundsService) DeleteBackgrounds(backgrounds war.Backgrounds) (err error) {
	err = global.GVA_DB.Delete(&backgrounds).Error
	return err
}

// DeleteBackgroundsByIds 批量删除Backgrounds记录
// Author [piexlmax](https://github.com/piexlmax)
func (backgroundsService *BackgroundsService) DeleteBackgroundsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.Backgrounds{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBackgrounds 更新Backgrounds记录
// Author [piexlmax](https://github.com/piexlmax)
func (backgroundsService *BackgroundsService) UpdateBackgrounds(backgrounds war.Backgrounds) (err error) {
	err = global.GVA_DB.Save(&backgrounds).Error
	return err
}

// GetBackgrounds 根据id获取Backgrounds记录
// Author [piexlmax](https://github.com/piexlmax)
func (backgroundsService *BackgroundsService) GetBackgrounds(id uint) (backgrounds war.Backgrounds, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&backgrounds).Error
	return
}

// GetBackgroundsInfoList 分页获取Backgrounds记录
// Author [piexlmax](https://github.com/piexlmax)
func (backgroundsService *BackgroundsService) GetBackgroundsInfoList(info warReq.BackgroundsSearch) (list []war.Backgrounds, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.Backgrounds{})
	var backgroundss []war.Backgrounds
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&backgroundss).Error
	return backgroundss, total, err
}

// 根据类型获取背景图
func (backgroundsService *BackgroundsService) GetBackgroundsByType(t int) (backgrounds war.Backgrounds, err error) {
	err = global.GVA_DB.Where("type = ? AND status = ?", t, 1).First(&backgrounds).Error
	return
}
