// 自动生成模板UserEquipment
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserEquipment 结构体
type UserEquipment struct {
      global.GVA_MODEL
      UserId  int `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`
      CategoryId  int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:分类ID;"`
      EquipmentId  int `json:"equipmentId" form:"equipmentId" gorm:"column:equipment_id;comment:装备ID;"`
      //会员信息
      User Member `json:"user" form:"user" gorm:"foreignKey:UserId;references:ID;comment:会员信息;"`
      //装备信息
      Equipment Equipment `json:"equipment" form:"equipment" gorm:"foreignKey:EquipmentId;references:ID;comment:装备信息;"`
      //装备分类信息
      Category Equipment `json:"category" form:"category" gorm:"foreignKey:CategoryId;references:ID;comment:装备分类信息;"`
}


// TableName UserEquipment 表名
func (UserEquipment) TableName() string {
  return "war_user_equipment"
}

