package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RankLevelRouter struct {
}

// InitRankLevelRouter 初始化 RankLevel 路由信息
func (s *RankLevelRouter) InitRankLevelRouter(Router *gin.RouterGroup) {
	rankLevelRouter := Router.Group("rankLevel").Use(middleware.OperationRecord())
	rankLevelRouterWithoutRecord := Router.Group("rankLevel")
	var rankLevelApi = v1.ApiGroupApp.WarApiGroup.RankLevelApi
	{
		rankLevelRouter.POST("createRankLevel", rankLevelApi.CreateRankLevel)   // 新建RankLevel
		rankLevelRouter.DELETE("deleteRankLevel", rankLevelApi.DeleteRankLevel) // 删除RankLevel
		rankLevelRouter.DELETE("deleteRankLevelByIds", rankLevelApi.DeleteRankLevelByIds) // 批量删除RankLevel
		rankLevelRouter.PUT("updateRankLevel", rankLevelApi.UpdateRankLevel)    // 更新RankLevel
	}
	{
		rankLevelRouterWithoutRecord.GET("findRankLevel", rankLevelApi.FindRankLevel)        // 根据ID获取RankLevel
		rankLevelRouterWithoutRecord.GET("getRankLevelList", rankLevelApi.GetRankLevelList)  // 获取RankLevel列表
	}
}
