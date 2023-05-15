package war

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MemberRoomRouter struct {
}

// InitMemberRoomRouter 初始化 MemberRoom 路由信息
func (s *MemberRoomRouter) InitMemberRoomRouter(Router *gin.RouterGroup) {
	memberRoomRouter := Router.Group("memberRoom").Use(middleware.OperationRecord())
	memberRoomRouterWithoutRecord := Router.Group("memberRoom")

	memberRoomPrivateRouter := Router.Group("memberRoom").Use(middleware.JWTAuth())
	var memberRoomApi = v1.ApiGroupApp.WarApiGroup.MemberRoomApi
	{
		memberRoomRouter.POST("createMemberRoom", memberRoomApi.CreateMemberRoom)             // 新建MemberRoom
		memberRoomRouter.DELETE("deleteMemberRoom", memberRoomApi.DeleteMemberRoom)           // 删除MemberRoom
		memberRoomRouter.DELETE("deleteMemberRoomByIds", memberRoomApi.DeleteMemberRoomByIds) // 批量删除MemberRoom
		memberRoomRouter.PUT("updateMemberRoom", memberRoomApi.UpdateMemberRoom)              // 更新MemberRoom
	}
	{
		memberRoomRouterWithoutRecord.GET("findMemberRoom", memberRoomApi.FindMemberRoom)       // 根据ID获取MemberRoom
		memberRoomRouterWithoutRecord.GET("getMemberRoomList", memberRoomApi.GetMemberRoomList) // 获取MemberRoom列表
	}
	{
		//用户签到
		memberRoomPrivateRouter.POST("sign", memberRoomApi.SignIn)

		//对局列表
		memberRoomPrivateRouter.GET("list", memberRoomApi.GetMemberRoomListByUserId)

		//对局详情
		memberRoomPrivateRouter.GET(":roomId", memberRoomApi.GetMemberRoomDetailByRoomId)
	}
}
