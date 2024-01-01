package war

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamMemberRouter struct {
}

// InitTeamMemberRouter 初始化 TeamMember 路由信息
func (s *TeamMemberRouter) InitTeamMemberRouter(Router *gin.RouterGroup) {
	teamMemberRouter := Router.Group("teamMember").Use(middleware.OperationRecord())
	teamMemberRouterWithoutRecord := Router.Group("teamMember")
	teamMemberPrivateRouter := Router.Group("teamMember").Use(middleware.JWTAuth())
	var teamMemberApi = v1.ApiGroupApp.WarApiGroup.TeamMemberApi
	{
		teamMemberRouter.POST("createTeamMember", teamMemberApi.CreateTeamMember)             // 新建TeamMember
		teamMemberRouter.DELETE("deleteTeamMember", teamMemberApi.DeleteTeamMember)           // 删除TeamMember
		teamMemberRouter.DELETE("deleteTeamMemberByIds", teamMemberApi.DeleteTeamMemberByIds) // 批量删除TeamMember
		teamMemberRouter.PUT("updateTeamMember", teamMemberApi.UpdateTeamMember)              // 更新TeamMember
	}
	{
		teamMemberRouterWithoutRecord.GET("findTeamMember", teamMemberApi.FindTeamMember)       // 根据ID获取TeamMember
		teamMemberRouterWithoutRecord.GET("getTeamMemberList", teamMemberApi.GetTeamMemberList) // 获取TeamMember列表
	}
	{
		//更改用户的角色
		teamMemberPrivateRouter.PUT("updateRole", teamMemberApi.UpdateTeamMemberRole)

		//退出战队
		teamMemberPrivateRouter.DELETE("exit", teamMemberApi.ExitTeam)
	}
}
