package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserEquipmentRouter struct {
}

// InitUserEquipmentRouter 初始化 UserEquipment 路由信息
func (s *UserEquipmentRouter) InitUserEquipmentRouter(Router *gin.RouterGroup) {
	userEquipmentRouter := Router.Group("userEquipment").Use(middleware.OperationRecord())
	userEquipmentRouterWithoutRecord := Router.Group("userEquipment")
	var userEquipmentApi = v1.ApiGroupApp.WarApiGroup.UserEquipmentApi
	{
		userEquipmentRouter.POST("createUserEquipment", userEquipmentApi.CreateUserEquipment)   // 新建UserEquipment
		userEquipmentRouter.DELETE("deleteUserEquipment", userEquipmentApi.DeleteUserEquipment) // 删除UserEquipment
		userEquipmentRouter.DELETE("deleteUserEquipmentByIds", userEquipmentApi.DeleteUserEquipmentByIds) // 批量删除UserEquipment
		userEquipmentRouter.PUT("updateUserEquipment", userEquipmentApi.UpdateUserEquipment)    // 更新UserEquipment
	}
	{
		userEquipmentRouterWithoutRecord.GET("findUserEquipment", userEquipmentApi.FindUserEquipment)        // 根据ID获取UserEquipment
		userEquipmentRouterWithoutRecord.GET("getUserEquipmentList", userEquipmentApi.GetUserEquipmentList)  // 获取UserEquipment列表
	}
}
