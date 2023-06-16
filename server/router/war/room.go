package war

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RoomRouter struct {
}

// InitRoomRouter 初始化 Room 路由信息
func (s *RoomRouter) InitRoomRouter(Router *gin.RouterGroup) {
	roomRouter := Router.Group("room").Use(middleware.OperationRecord())
	roomRouterWithoutRecord := Router.Group("room")

	roomPrivateRouter := Router.Group("room").Use(middleware.JWTAuth())

	var roomApi = v1.ApiGroupApp.WarApiGroup.RoomApi
	{
		roomRouter.POST("createRoom", roomApi.CreateRoom)             // 新建Room
		roomRouter.DELETE("deleteRoom", roomApi.DeleteRoom)           // 删除Room
		roomRouter.DELETE("deleteRoomByIds", roomApi.DeleteRoomByIds) // 批量删除Room
		roomRouter.PUT("updateRoom", roomApi.UpdateRoom)              // 更新Room
		//开始游戏
		roomRouter.POST("startGame", roomApi.StartGame)
	}
	{
		roomRouterWithoutRecord.GET("findRoom", roomApi.FindRoom)       // 根据ID获取Room
		roomRouterWithoutRecord.GET("getRoomList", roomApi.GetRoomList) // 获取Room列表
	}
	{
		//快速匹配游戏
		roomPrivateRouter.POST("quick", roomApi.QuickMatch)

		//倒计时结束
		roomPrivateRouter.POST("countdown", roomApi.Countdown)

		//获取房间二维码
		roomPrivateRouter.GET("qrcode", roomApi.GetRoomQrCode)
	}
}
