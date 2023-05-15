package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MemberLevelRouter struct {
}

// InitMemberLevelRouter 初始化 MemberLevel 路由信息
func (s *MemberLevelRouter) InitMemberLevelRouter(Router *gin.RouterGroup) {
	memberLevelRouter := Router.Group("memberLevel").Use(middleware.OperationRecord())
	memberLevelRouterWithoutRecord := Router.Group("memberLevel")
	var memberLevelApi = v1.ApiGroupApp.WarApiGroup.MemberLevelApi
	{
		memberLevelRouter.POST("createMemberLevel", memberLevelApi.CreateMemberLevel)   // 新建MemberLevel
		memberLevelRouter.DELETE("deleteMemberLevel", memberLevelApi.DeleteMemberLevel) // 删除MemberLevel
		memberLevelRouter.DELETE("deleteMemberLevelByIds", memberLevelApi.DeleteMemberLevelByIds) // 批量删除MemberLevel
		memberLevelRouter.PUT("updateMemberLevel", memberLevelApi.UpdateMemberLevel)    // 更新MemberLevel
	}
	{
		memberLevelRouterWithoutRecord.GET("findMemberLevel", memberLevelApi.FindMemberLevel)        // 根据ID获取MemberLevel
		memberLevelRouterWithoutRecord.GET("getMemberLevelList", memberLevelApi.GetMemberLevelList)  // 获取MemberLevel列表
	}
}
