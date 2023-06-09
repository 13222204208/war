package war

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamApplyRouter struct {
}

// InitTeamApplyRouter 初始化 TeamApply 路由信息
func (s *TeamApplyRouter) InitTeamApplyRouter(Router *gin.RouterGroup) {
	teamApplyRouter := Router.Group("teamApply").Use(middleware.OperationRecord())
	teamApplyRouterWithoutRecord := Router.Group("teamApply")
	teamApplyPrivateRouter := Router.Group("teamApply").Use(middleware.JWTAuth())
	var teamApplyApi = v1.ApiGroupApp.WarApiGroup.TeamApplyApi
	{
		teamApplyRouter.POST("createTeamApply", teamApplyApi.CreateTeamApply)             // 新建TeamApply
		teamApplyRouter.DELETE("deleteTeamApply", teamApplyApi.DeleteTeamApply)           // 删除TeamApply
		teamApplyRouter.DELETE("deleteTeamApplyByIds", teamApplyApi.DeleteTeamApplyByIds) // 批量删除TeamApply
		teamApplyRouter.PUT("updateTeamApply", teamApplyApi.UpdateTeamApply)              // 更新TeamApply
	}
	{
		teamApplyRouterWithoutRecord.GET("findTeamApply", teamApplyApi.FindTeamApply)       // 根据ID获取TeamApply
		teamApplyRouterWithoutRecord.GET("getTeamApplyList", teamApplyApi.GetTeamApplyList) // 获取TeamApply列表
	}
	{
		teamApplyPrivateRouter.POST("join", teamApplyApi.SaveTeamApply) // 战队申请
		//加入战队申请列表
		teamApplyPrivateRouter.GET("list", teamApplyApi.GetTeamApplyListByUserId)

		//审批用户的申请
		teamApplyPrivateRouter.PUT(":id", teamApplyApi.ApprovalTeamApply)
	}
}
