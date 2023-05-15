package war

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QuestionRouter struct {
}

// InitQuestionRouter 初始化 Question 路由信息
func (s *QuestionRouter) InitQuestionRouter(Router *gin.RouterGroup) {
	questionRouter := Router.Group("question").Use(middleware.OperationRecord())
	questionRouterWithoutRecord := Router.Group("question")
	var questionApi = v1.ApiGroupApp.WarApiGroup.QuestionApi
	{
		questionRouter.POST("createQuestion", questionApi.CreateQuestion)   // 新建Question
		questionRouter.DELETE("deleteQuestion", questionApi.DeleteQuestion) // 删除Question
		questionRouter.DELETE("deleteQuestionByIds", questionApi.DeleteQuestionByIds) // 批量删除Question
		questionRouter.PUT("updateQuestion", questionApi.UpdateQuestion)    // 更新Question
	}
	{
		questionRouterWithoutRecord.GET("findQuestion", questionApi.FindQuestion)        // 根据ID获取Question
		questionRouterWithoutRecord.GET("getQuestionList", questionApi.GetQuestionList)  // 获取Question列表
	}
	{
		//获取所有的问题和答案
		questionRouterWithoutRecord.GET("list", questionApi.GetAllQuestion)
	}
}
