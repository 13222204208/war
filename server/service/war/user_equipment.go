package war

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
	warRes "github.com/flipped-aurora/gin-vue-admin/server/model/war/response"
	"gorm.io/gorm"
)

type UserEquipmentService struct {
}

// CreateUserEquipment 创建UserEquipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEquipmentService *UserEquipmentService) CreateUserEquipment(add warReq.UserEquipmentAdd) (err error) {
	// equipment := add.Equipment
	// //json字符串解析为数组

	// var equipmentArr []war.UserEquipment
	// err = json.Unmarshal([]byte(equipment), &equipmentArr)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(equipmentArr)
	// err = global.GVA_DB.Create(&equipmentArr).Error
	// return err
	equipment := add.Equipment
	userId := add.UserId
	//查询装备是否已经添加
	var userEquipment war.UserEquipment
	err = global.GVA_DB.Where("user_id = ? AND equipment_id = ?", userId, equipment).First(&userEquipment).Error
	if err == gorm.ErrRecordNotFound {
		//未添加
		//equipment转为int
		if equipment == 0 {
			return fmt.Errorf("装备id错误")
		}
		//查询装备id的分类id
		var equipmentInfo war.Equipment
		err = global.GVA_DB.Where("id = ?", equipment).First(&equipmentInfo).Error
		if err != nil {
			return err
		}

		err = global.GVA_DB.Create(&war.UserEquipment{UserId: int(userId), EquipmentId: equipment, CategoryId: *equipmentInfo.ParentId}).Error
		if err != nil {
			return err
		}
	} else {
		//已添加
		return fmt.Errorf("该装备已添加")
	}
	return err
}

// DeleteUserEquipment 删除UserEquipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEquipmentService *UserEquipmentService) DeleteUserEquipment(userEquipment war.UserEquipment) (err error) {
	err = global.GVA_DB.Delete(&userEquipment).Error
	return err
}

// DeleteUserEquipmentByIds 批量删除UserEquipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEquipmentService *UserEquipmentService) DeleteUserEquipmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.UserEquipment{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserEquipment 更新UserEquipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEquipmentService *UserEquipmentService) UpdateUserEquipment(userEquipment war.UserEquipment) (err error) {
	err = global.GVA_DB.Save(&userEquipment).Error
	return err
}

// GetUserEquipment 根据id获取UserEquipment记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEquipmentService *UserEquipmentService) GetUserEquipment(id uint) (userEquipment war.UserEquipment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userEquipment).Error
	return
}

// GetUserEquipmentInfoList 分页获取UserEquipment记录
// Author [piexlmax](https://github.com/piexlmax)
// func (userEquipmentService *UserEquipmentService) GetUserEquipmentInfoList(info warReq.UserEquipmentSearch) (list []war.UserEquipment, total int64, err error) {
// 	limit := info.PageSize
// 	offset := info.PageSize * (info.Page - 1)
// 	// 创建db
// 	db := global.GVA_DB.Model(&war.UserEquipment{})
// 	var userEquipments []war.UserEquipment
// 	// 如果有条件搜索 下方会自动创建搜索语句
// 	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
// 		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
// 	}

// 	if info.UserId != 0 {
// 		db = db.Where("user_id = ?", info.UserId)
// 	}

// 	err = db.Count(&total).Error
// 	if err != nil {
// 		return
// 	}

// 	err = db.Limit(limit).Preload("User").Preload("Equipment").Preload("Category").Offset(offset).Find(&userEquipments).Error
// 	return userEquipments, total, err
// }

func (userEquipmentService *UserEquipmentService) GetUserEquipmentInfoList(info warReq.UserEquipmentSearch) (equipments []warRes.Equipments, err error) {
	//查询出所有的装备
	err = global.GVA_DB.Model(&war.Equipment{}).Find(&equipments).Error
	if err != nil {
		return
	}
	//查询出所有的用户装备
	var userEquipments []war.UserEquipment
	err = global.GVA_DB.Where("user_id", info.UserId).Find(&userEquipments).Error
	if err != nil {
		return
	}
	if len(userEquipments) != 0 {
		//如果userEquipments中的装备id等于equipments中的装备id，就把equipments中status等于1，否则等于0
		for i := 0; i < len(equipments); i++ {
			for j := 0; j < len(userEquipments); j++ {
				if equipments[i].ID == uint(userEquipments[j].EquipmentId) {
					equipments[i].Status = 2
				}
			}
		}
	}
	return
}
