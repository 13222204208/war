package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type QuestionApi struct {
}

var questionService = service.ServiceGroupApp.WarServiceGroup.QuestionService


// CreateQuestion 创建Question
// @Tags Question
// @Summary 创建Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Question true "创建Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question/createQuestion [post]
func (questionApi *QuestionApi) CreateQuestion(c *gin.Context) {
	var question war.Question
	err := c.ShouldBindJSON(&question)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Content":{utils.NotEmpty()},
        "Type":{utils.NotEmpty()},
    }
	if err := utils.Verify(question, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := questionService.CreateQuestion(&question); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuestion 删除Question
// @Tags Question
// @Summary 删除Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Question true "删除Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /question/deleteQuestion [delete]
func (questionApi *QuestionApi) DeleteQuestion(c *gin.Context) {
	var question war.Question
	err := c.ShouldBindJSON(&question)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := questionService.DeleteQuestion(question); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuestionByIds 批量删除Question
// @Tags Question
// @Summary 批量删除Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /question/deleteQuestionByIds [delete]
func (questionApi *QuestionApi) DeleteQuestionByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := questionService.DeleteQuestionByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuestion 更新Question
// @Tags Question
// @Summary 更新Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Question true "更新Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /question/updateQuestion [put]
func (questionApi *QuestionApi) UpdateQuestion(c *gin.Context) {
	var question war.Question
	err := c.ShouldBindJSON(&question)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Content":{utils.NotEmpty()},
          "Type":{utils.NotEmpty()},
      }
    if err := utils.Verify(question, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := questionService.UpdateQuestion(question); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuestion 用id查询Question
// @Tags Question
// @Summary 用id查询Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.Question true "用id查询Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /question/findQuestion [get]
func (questionApi *QuestionApi) FindQuestion(c *gin.Context) {
	var question war.Question
	err := c.ShouldBindQuery(&question)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if requestion, err := questionService.GetQuestion(question.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestion": requestion}, c)
	}
}

// GetQuestionList 分页获取Question列表
// @Tags Question
// @Summary 分页获取Question列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.QuestionSearch true "分页获取Question列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question/getQuestionList [get]
func (questionApi *QuestionApi) GetQuestionList(c *gin.Context) {
	var pageInfo warReq.QuestionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := questionService.GetQuestionInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

//获取所有的问题和答案
func (questionApi *QuestionApi) GetAllQuestion(c *gin.Context) {
	if list, err := questionService.GetQuestionAndAnswer(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}