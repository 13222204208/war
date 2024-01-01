package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SlotRouter struct {
}

// InitSlotRouter 初始化 Slot 路由信息
func (s *SlotRouter) InitSlotRouter(Router *gin.RouterGroup) {
	slotRouter := Router.Group("slot").Use(middleware.OperationRecord())
	slotRouterWithoutRecord := Router.Group("slot")
	var slotApi = v1.ApiGroupApp.WarApiGroup.SlotApi
	{
		slotRouter.POST("createSlot", slotApi.CreateSlot)   // 新建Slot
		slotRouter.DELETE("deleteSlot", slotApi.DeleteSlot) // 删除Slot
		slotRouter.DELETE("deleteSlotByIds", slotApi.DeleteSlotByIds) // 批量删除Slot
		slotRouter.PUT("updateSlot", slotApi.UpdateSlot)    // 更新Slot
	}
	{
		slotRouterWithoutRecord.GET("findSlot", slotApi.FindSlot)        // 根据ID获取Slot
		slotRouterWithoutRecord.GET("getSlotList", slotApi.GetSlotList)  // 获取Slot列表
	}
}
