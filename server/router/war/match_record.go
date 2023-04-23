package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MatchRecordRouter struct {
}

// InitMatchRecordRouter 初始化 MatchRecord 路由信息
func (s *MatchRecordRouter) InitMatchRecordRouter(Router *gin.RouterGroup) {
	matchRecordRouter := Router.Group("matchRecord").Use(middleware.OperationRecord())
	matchRecordRouterWithoutRecord := Router.Group("matchRecord")
	var matchRecordApi = v1.ApiGroupApp.WarApiGroup.MatchRecordApi
	{
		matchRecordRouter.POST("createMatchRecord", matchRecordApi.CreateMatchRecord)   // 新建MatchRecord
		matchRecordRouter.DELETE("deleteMatchRecord", matchRecordApi.DeleteMatchRecord) // 删除MatchRecord
		matchRecordRouter.DELETE("deleteMatchRecordByIds", matchRecordApi.DeleteMatchRecordByIds) // 批量删除MatchRecord
		matchRecordRouter.PUT("updateMatchRecord", matchRecordApi.UpdateMatchRecord)    // 更新MatchRecord
	}
	{
		matchRecordRouterWithoutRecord.GET("findMatchRecord", matchRecordApi.FindMatchRecord)        // 根据ID获取MatchRecord
		matchRecordRouterWithoutRecord.GET("getMatchRecordList", matchRecordApi.GetMatchRecordList)  // 获取MatchRecord列表
	}
}
