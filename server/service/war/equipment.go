package war

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
	warRes "github.com/flipped-aurora/gin-vue-admin/server/model/war/response"
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
func (equipmentService *EquipmentService) DeleteEquipment(equipment war.Equipment) (err error) {
	err = global.GVA_DB.Delete(&equipment).Error
	return err
}

// DeleteEquipmentByIds 批量删除Equipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (equipmentService *EquipmentService) DeleteEquipmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.Equipment{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEquipment 更新Equipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (equipmentService *EquipmentService) UpdateEquipment(equipment war.Equipment) (err error) {
	err = global.GVA_DB.Save(&equipment).Error
	return err
}

// GetEquipment 根据id获取Equipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (equipmentService *EquipmentService) GetEquipment(id uint) (equipment war.Equipment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&equipment).Error
	return
}

// GetEquipmentInfoList 分页获取Equipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (equipmentService *EquipmentService) GetEquipmentInfoList(info warReq.EquipmentSearch) (list []war.Equipment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.Equipment{})
	var equipments []war.Equipment
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&equipments).Error
	return equipments, total, err
}

// 我的装备详情
func (equipmentService *EquipmentService) Detail(userId uint) (e warRes.MyEquipmentList, err error) {
	var userEquipment []war.UserEquipment
	err = global.GVA_DB.Where("user_id = ?", userId).Find(&userEquipment).Error
	if err != nil {
		return
	}

	if len(userEquipment) == 0 {
		return
	}

	equipmentIDs := make([]int, len(userEquipment))

	// 获取用户装备ID
	for i, v := range userEquipment {
		equipmentIDs[i] = v.EquipmentId
	}

	var equipment []war.Equipment

	// 预加载装备数据
	err = global.GVA_DB.Select("id,name,icon,parent_id").Find(&equipment, equipmentIDs).Error
	if err != nil {
		return
	}
	fmt.Println("装备数据", equipment)

	for _, v := range equipment {
		fmt.Println("装备", v.Name, "PID", *v.ParentId)
		pid := *v.ParentId
		switch pid {
		case 1:
			e.MainWeapon.Name = v.Name
			e.MainWeapon.Icon = v.Icon
		case 2:
			e.SecondaryWeapon.Name = v.Name
			e.SecondaryWeapon.Icon = v.Icon
		case 11:
			e.Glasses.Name = v.Name
			e.Glasses.Icon = v.Icon
		case 10:
			e.Clothing.Name = v.Name
			e.Clothing.Icon = v.Icon
		case 8:
			e.Shoes.Name = v.Name
			e.Shoes.Icon = v.Icon
		}
	}
	return
}
