package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type SlotService struct {
}

// CreateSlot 创建Slot记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotService *SlotService) CreateSlot(slot *war.Slot) (err error) {
	err = global.GVA_DB.Create(slot).Error
	return err
}

// DeleteSlot 删除Slot记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotService *SlotService)DeleteSlot(slot war.Slot) (err error) {
	err = global.GVA_DB.Delete(&slot).Error
	return err
}

// DeleteSlotByIds 批量删除Slot记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotService *SlotService)DeleteSlotByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.Slot{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSlot 更新Slot记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotService *SlotService)UpdateSlot(slot war.Slot) (err error) {
	err = global.GVA_DB.Save(&slot).Error
	return err
}

// GetSlot 根据id获取Slot记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotService *SlotService)GetSlot(id uint) (slot war.Slot, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slot).Error
	return
}

// GetSlotInfoList 分页获取Slot记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotService *SlotService)GetSlotInfoList(info warReq.SlotSearch) (list []war.Slot, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&war.Slot{})
    var slots []war.Slot
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&slots).Error
	return  slots, total, err
}
