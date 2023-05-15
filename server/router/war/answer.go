package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AnswerRouter struct {
}

// InitAnswerRouter 初始化 Answer 路由信息
func (s *AnswerRouter) InitAnswerRouter(Router *gin.RouterGroup) {
	answerRouter := Router.Group("answer").Use(middleware.OperationRecord())
	answerRouterWithoutRecord := Router.Group("answer")
	var answerApi = v1.ApiGroupApp.WarApiGroup.AnswerApi
	{
		answerRouter.POST("createAnswer", answerApi.CreateAnswer)   // 新建Answer
		answerRouter.DELETE("deleteAnswer", answerApi.DeleteAnswer) // 删除Answer
		answerRouter.DELETE("deleteAnswerByIds", answerApi.DeleteAnswerByIds) // 批量删除Answer
		answerRouter.PUT("updateAnswer", answerApi.UpdateAnswer)    // 更新Answer
	}
	{
		answerRouterWithoutRecord.GET("findAnswer", answerApi.FindAnswer)        // 根据ID获取Answer
		answerRouterWithoutRecord.GET("getAnswerList", answerApi.GetAnswerList)  // 获取Answer列表
	}
}
