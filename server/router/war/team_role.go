package war

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamRoleRouter struct {
}

// InitTeamRoleRouter 初始化 TeamRole 路由信息
func (s *TeamRoleRouter) InitTeamRoleRouter(Router *gin.RouterGroup) {
	teamRoleRouter := Router.Group("teamRole").Use(middleware.OperationRecord())
	teamRoleRouterWithoutRecord := Router.Group("teamRole")

	teamRolePrivateRouter := Router.Group("teamRole").Use(middleware.JWTAuth())
	var teamRoleApi = v1.ApiGroupApp.WarApiGroup.TeamRoleApi
	{
		teamRoleRouter.POST("createTeamRole", teamRoleApi.CreateTeamRole)             // 新建TeamRole
		teamRoleRouter.DELETE("deleteTeamRole", teamRoleApi.DeleteTeamRole)           // 删除TeamRole
		teamRoleRouter.DELETE("deleteTeamRoleByIds", teamRoleApi.DeleteTeamRoleByIds) // 批量删除TeamRole
		teamRoleRouter.PUT("updateTeamRole", teamRoleApi.UpdateTeamRole)              // 更新TeamRole
	}
	{
		teamRoleRouterWithoutRecord.GET("findTeamRole", teamRoleApi.FindTeamRole)       // 根据ID获取TeamRole
		teamRoleRouterWithoutRecord.GET("getTeamRoleList", teamRoleApi.GetTeamRoleList) // 获取TeamRole列表
	}
	{
		teamRolePrivateRouter.GET("list", teamRoleApi.TeamRoleList) // 获取TeamRole列表
	}
}
