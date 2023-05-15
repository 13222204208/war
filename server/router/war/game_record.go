package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GameRecordRouter struct {
}

// InitGameRecordRouter 初始化 GameRecord 路由信息
func (s *GameRecordRouter) InitGameRecordRouter(Router *gin.RouterGroup) {
	gameRecordRouter := Router.Group("gameRecord").Use(middleware.OperationRecord())
	gameRecordRouterWithoutRecord := Router.Group("gameRecord")
	var gameRecordApi = v1.ApiGroupApp.WarApiGroup.GameRecordApi
	{
		gameRecordRouter.POST("createGameRecord", gameRecordApi.CreateGameRecord)   // 新建GameRecord
		gameRecordRouter.DELETE("deleteGameRecord", gameRecordApi.DeleteGameRecord) // 删除GameRecord
		gameRecordRouter.DELETE("deleteGameRecordByIds", gameRecordApi.DeleteGameRecordByIds) // 批量删除GameRecord
		gameRecordRouter.PUT("updateGameRecord", gameRecordApi.UpdateGameRecord)    // 更新GameRecord
	}
	{
		gameRecordRouterWithoutRecord.GET("findGameRecord", gameRecordApi.FindGameRecord)        // 根据ID获取GameRecord
		gameRecordRouterWithoutRecord.GET("getGameRecordList", gameRecordApi.GetGameRecordList)  // 获取GameRecord列表
	}
}
