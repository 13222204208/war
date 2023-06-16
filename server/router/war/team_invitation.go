package war

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamInvitationRouter struct {
}

// InitTeamInvitationRouter 初始化 TeamInvitation 路由信息
func (s *TeamInvitationRouter) InitTeamInvitationRouter(Router *gin.RouterGroup) {
	teamInvitationRouter := Router.Group("teamInvitation").Use(middleware.OperationRecord())
	teamInvitationRouterWithoutRecord := Router.Group("teamInvitation")

	teamInvitationPrivateRouter := Router.Group("teamInvitation").Use(middleware.JWTAuth())

	var teamInvitationApi = v1.ApiGroupApp.WarApiGroup.TeamInvitationApi
	{
		teamInvitationRouter.POST("createTeamInvitation", teamInvitationApi.CreateTeamInvitation)             // 新建TeamInvitation
		teamInvitationRouter.DELETE("deleteTeamInvitation", teamInvitationApi.DeleteTeamInvitation)           // 删除TeamInvitation
		teamInvitationRouter.DELETE("deleteTeamInvitationByIds", teamInvitationApi.DeleteTeamInvitationByIds) // 批量删除TeamInvitation
		teamInvitationRouter.PUT("updateTeamInvitation", teamInvitationApi.UpdateTeamInvitation)              // 更新TeamInvitation
	}
	{
		teamInvitationRouterWithoutRecord.GET("findTeamInvitation", teamInvitationApi.FindTeamInvitation)       // 根据ID获取TeamInvitation
		teamInvitationRouterWithoutRecord.GET("getTeamInvitationList", teamInvitationApi.GetTeamInvitationList) // 获取TeamInvitation列表
	}
	{
		//战队赛邀请
		teamInvitationPrivateRouter.POST("invite", teamInvitationApi.TeamInvitation)
		//战队受邀
		teamInvitationPrivateRouter.GET("list", teamInvitationApi.TeamInvitationList)
	}
}
