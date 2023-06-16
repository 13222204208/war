package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
)

type UserEquipmentSearch struct {
	war.UserEquipment
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

// 给用户添加装备
type UserEquipmentAdd struct {
	UserId    uint `json:"userId" form:"userId"`
	Equipment int  `json:"equipment" form:"equipment"`
}
