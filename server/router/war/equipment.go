package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EquipmentRouter struct {
}

// InitEquipmentRouter 初始化 Equipment 路由信息
func (s *EquipmentRouter) InitEquipmentRouter(Router *gin.RouterGroup) {
	equipmentRouter := Router.Group("equipment").Use(middleware.OperationRecord())
	equipmentRouterWithoutRecord := Router.Group("equipment")
	var equipmentApi = v1.ApiGroupApp.WarApiGroup.EquipmentApi
	{
		equipmentRouter.POST("createEquipment", equipmentApi.CreateEquipment)   // 新建Equipment
		equipmentRouter.DELETE("deleteEquipment", equipmentApi.DeleteEquipment) // 删除Equipment
		equipmentRouter.DELETE("deleteEquipmentByIds", equipmentApi.DeleteEquipmentByIds) // 批量删除Equipment
		equipmentRouter.PUT("updateEquipment", equipmentApi.UpdateEquipment)    // 更新Equipment
	}
	{
		equipmentRouterWithoutRecord.GET("findEquipment", equipmentApi.FindEquipment)        // 根据ID获取Equipment
		equipmentRouterWithoutRecord.GET("getEquipmentList", equipmentApi.GetEquipmentList)  // 获取Equipment列表
	}
}
