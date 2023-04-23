package war

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BackgroundsRouter struct {
}

// InitBackgroundsRouter 初始化 Backgrounds 路由信息
func (s *BackgroundsRouter) InitBackgroundsRouter(Router *gin.RouterGroup) {
	backgroundsRouter := Router.Group("backgrounds").Use(middleware.OperationRecord())
	backgroundsRouterWithoutRecord := Router.Group("backgrounds")
	var backgroundsApi = v1.ApiGroupApp.WarApiGroup.BackgroundsApi
	{
		backgroundsRouter.POST("createBackgrounds", backgroundsApi.CreateBackgrounds)             // 新建Backgrounds
		backgroundsRouter.DELETE("deleteBackgrounds", backgroundsApi.DeleteBackgrounds)           // 删除Backgrounds
		backgroundsRouter.DELETE("deleteBackgroundsByIds", backgroundsApi.DeleteBackgroundsByIds) // 批量删除Backgrounds
		backgroundsRouter.PUT("updateBackgrounds", backgroundsApi.UpdateBackgrounds)              // 更新Backgrounds
	}
	{
		backgroundsRouterWithoutRecord.GET("findBackgrounds", backgroundsApi.FindBackgrounds)       // 根据ID获取Backgrounds
		backgroundsRouterWithoutRecord.GET("getBackgroundsList", backgroundsApi.GetBackgroundsList) // 获取Backgrounds列表

		//根据类型获取背景图
		backgroundsRouterWithoutRecord.GET(":type", backgroundsApi.GetBackgroundsByType)
	}
}
