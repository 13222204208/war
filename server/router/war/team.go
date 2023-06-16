package war

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamRouter struct {
}

// InitTeamRouter 初始化 Team 路由信息
func (s *TeamRouter) InitTeamRouter(Router *gin.RouterGroup) {
	teamRouter := Router.Group("team").Use(middleware.OperationRecord())
	teamRouterWithoutRecord := Router.Group("team")
	teamPrivateRouter := Router.Group("team").Use(middleware.JWTAuth())
	var teamApi = v1.ApiGroupApp.WarApiGroup.TeamApi
	{
		// teamRouter.POST("createTeam", teamApi.CreateTeam)   // 新建Team
		teamRouter.DELETE("deleteTeam", teamApi.DeleteTeam)           // 删除Team
		teamRouter.DELETE("deleteTeamByIds", teamApi.DeleteTeamByIds) // 批量删除Team
		teamRouter.PUT("updateTeam", teamApi.UpdateTeam)              // 更新Team
	}
	{
		teamRouterWithoutRecord.GET("findTeam", teamApi.FindTeam)       // 根据ID获取Team
		teamRouterWithoutRecord.GET("getTeamList", teamApi.GetTeamList) // 获取Team列表
	}
	{
		//创建战队
		teamPrivateRouter.POST("create", teamApi.CreateTeam)
		//修改战队信息
		teamPrivateRouter.PUT("update", teamApi.UpdateTeamInfo)
		//获取全部战队信息
		teamPrivateRouter.GET("all", teamApi.GetAllTeam)
		//获取战队详情
		teamPrivateRouter.GET(":id", teamApi.GetTeamDetail)
		//获取我的战队详情
		teamPrivateRouter.GET("my", teamApi.GetMyTeamDetail)

		//战队邀请海报
		teamPrivateRouter.GET("poster", teamApi.GetTeamPoster)
	}
}
