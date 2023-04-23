package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type EquipmentService struct {
}

// CreateEquipment 创建Equipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (equipmentService *EquipmentService) CreateEquipment(equipment *war.Equipment) (err error) {
	err = global.GVA_DB.Create(equipment).Error
	return err
}

// DeleteEquipment 删除Equipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (equipmentService *EquipmentService)DeleteEquipment(equipment war.Equipment) (err error) {
	err = global.GVA_DB.Delete(&equipment).Error
	return err
}

// DeleteEquipmentByIds 批量删除Equipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (equipmentService *EquipmentService)DeleteEquipmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.Equipment{},"id in ?",ids.Ids).Error
	return err
}

// UpdateEquipment 更新Equipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (equipmentService *EquipmentService)UpdateEquipment(equipment war.Equipment) (err error) {
	err = global.GVA_DB.Save(&equipment).Error
	return err
}

// GetEquipment 根据id获取Equipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (equipmentService *EquipmentService)GetEquipment(id uint) (equipment war.Equipment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&equipment).Error
	return
}

// GetEquipmentInfoList 分页获取Equipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (equipmentService *EquipmentService)GetEquipmentInfoList(info warReq.EquipmentSearch) (list []war.Equipment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&war.Equipment{})
    var equipments []war.Equipment
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&equipments).Error
	return  equipments, total, err
}
