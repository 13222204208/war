package war

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type UserEquipmentService struct {
}

// CreateUserEquipment 创建UserEquipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEquipmentService *UserEquipmentService) CreateUserEquipment(add warReq.UserEquipmentAdd) (err error) {
	equipment := add.Equipment
	//json字符串解析为数组

	var equipmentArr []war.UserEquipment
	err = json.Unmarshal([]byte(equipment), &equipmentArr)
	if err != nil {
		return err
	}
	fmt.Println(equipmentArr)
	err = global.GVA_DB.Create(&equipmentArr).Error
	return err
}

// DeleteUserEquipment 删除UserEquipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEquipmentService *UserEquipmentService)DeleteUserEquipment(userEquipment war.UserEquipment) (err error) {
	err = global.GVA_DB.Delete(&userEquipment).Error
	return err
}

// DeleteUserEquipmentByIds 批量删除UserEquipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEquipmentService *UserEquipmentService)DeleteUserEquipmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.UserEquipment{},"id in ?",ids.Ids).Error
	return err
}

// UpdateUserEquipment 更新UserEquipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEquipmentService *UserEquipmentService)UpdateUserEquipment(userEquipment war.UserEquipment) (err error) {
	err = global.GVA_DB.Save(&userEquipment).Error
	return err
}

// GetUserEquipment 根据id获取UserEquipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEquipmentService *UserEquipmentService)GetUserEquipment(id uint) (userEquipment war.UserEquipment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userEquipment).Error
	return
}

// GetUserEquipmentInfoList 分页获取UserEquipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEquipmentService *UserEquipmentService)GetUserEquipmentInfoList(info warReq.UserEquipmentSearch) (list []war.UserEquipment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&war.UserEquipment{})
    var userEquipments []war.UserEquipment
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }

	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}
	
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Preload("User").Preload("Equipment").Preload("Category").Offset(offset).Find(&userEquipments).Error
	return  userEquipments, total, err
}
