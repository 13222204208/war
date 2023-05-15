package war

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MemberRouter struct {
}

// InitMemberRouter 初始化 Member 路由信息
func (s *MemberRouter) InitMemberRouter(Router *gin.RouterGroup) {
	memberRouter := Router.Group("member").Use(middleware.OperationRecord())
	memberRouterWithoutRecord := Router.Group("member")

	memberPrivateRouter := Router.Group("member").Use(middleware.JWTAuth())

	exaFileUploadAndDownloadApi := v1.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
	var memberApi = v1.ApiGroupApp.WarApiGroup.MemberApi
	{
		memberRouter.POST("createMember", memberApi.CreateMember)             // 新建Member
		memberRouter.DELETE("deleteMember", memberApi.DeleteMember)           // 删除Member
		memberRouter.DELETE("deleteMemberByIds", memberApi.DeleteMemberByIds) // 批量删除Member
		memberRouter.PUT("updateMember", memberApi.UpdateMember)              // 更新Member

		// 会员增加或修改场次
		memberRouter.PUT("updateMemberMatch", memberApi.AddOrUpdateMemberMatch)
	}
	{
		memberRouterWithoutRecord.GET("findMember", memberApi.FindMember)       // 根据ID获取Member
		memberRouterWithoutRecord.GET("getMemberList", memberApi.GetMemberList) // 获取Member列表
	}

	{
		//登陆
		memberRouterWithoutRecord.POST("login", memberApi.Login)
	}
	{
		//修改会员资料
		memberPrivateRouter.PUT("update", memberApi.UpdateMemberInfo)

		//获取会员资料
		memberPrivateRouter.GET("info", memberApi.GetMemberInfo)

		//图片上传
		memberPrivateRouter.POST("upload", exaFileUploadAndDownloadApi.UploadFile)

		//获取个人详情
		memberPrivateRouter.GET("detail", memberApi.GetMemberDetail)
	}
}
