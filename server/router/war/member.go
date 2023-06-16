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
	memberPublicRouter := Router.Group("member")

	exaFileUploadAndDownloadApi := v1.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
	var memberApi = v1.ApiGroupApp.WarApiGroup.MemberApi
	{
		memberRouter.POST("createMember", memberApi.CreateMember)             // 新建Member
		memberRouter.DELETE("deleteMember", memberApi.DeleteMember)           // 删除Member
		memberRouter.DELETE("deleteMemberByIds", memberApi.DeleteMemberByIds) // 批量删除Member
		memberRouter.PUT("updateMember", memberApi.UpdateMember)              // 更新Member

		// 会员增加或修改场次
		memberRouter.PUT("updateMemberMatch", memberApi.AddOrUpdateMemberMatch)

		//导入会员
		memberRouter.POST("importExcel", memberApi.ImportExcel)
	}
	{
		memberRouterWithoutRecord.GET("findMember", memberApi.FindMember)       // 根据ID获取Member
		memberRouterWithoutRecord.GET("getMemberList", memberApi.GetMemberList) // 获取Member列表
	}

	{
		//登陆
		memberRouterWithoutRecord.POST("login", memberApi.Login)

		//获取会员手机号
		memberRouterWithoutRecord.POST("phone", memberApi.GetMemberPhone)
	}
	{
		//修改会员资料
		memberPrivateRouter.PUT("update", memberApi.UpdateMemberInfo)

		//获取会员资料
		memberPrivateRouter.GET("info", memberApi.GetMemberInfo)

		//图片上传
		memberPublicRouter.POST("upload", exaFileUploadAndDownloadApi.UploadFile)

		//获取个人详情
		memberPrivateRouter.GET("detail", memberApi.GetMemberDetail)

		//获取我的战绩
		memberPrivateRouter.GET("kda", memberApi.GetMyKda)

		//获取我的战斗信息
		memberPrivateRouter.GET("battle", memberApi.GetMyBattleInfo)

		//用户排行
		memberPrivateRouter.GET("rank", memberApi.GetMemberRank)

	}
}
