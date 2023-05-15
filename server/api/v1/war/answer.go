package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/war"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AnswerApi struct {
}

var answerService = service.ServiceGroupApp.WarServiceGroup.AnswerService


// CreateAnswer 创建Answer
// @Tags Answer
// @Summary 创建Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Answer true "创建Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /answer/createAnswer [post]
func (answerApi *AnswerApi) CreateAnswer(c *gin.Context) {
	var answer war.Answer
	err := c.ShouldBindJSON(&answer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "QuestionId":{utils.NotEmpty()},
        "Content":{utils.NotEmpty()},
        "IsCorrect":{utils.NotEmpty()},
    }
	if err := utils.Verify(answer, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := answerService.CreateAnswer(&answer); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAnswer 删除Answer
// @Tags Answer
// @Summary 删除Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Answer true "删除Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /answer/deleteAnswer [delete]
func (answerApi *AnswerApi) DeleteAnswer(c *gin.Context) {
	var answer war.Answer
	err := c.ShouldBindJSON(&answer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := answerService.DeleteAnswer(answer); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAnswerByIds 批量删除Answer
// @Tags Answer
// @Summary 批量删除Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /answer/deleteAnswerByIds [delete]
func (answerApi *AnswerApi) DeleteAnswerByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := answerService.DeleteAnswerByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAnswer 更新Answer
// @Tags Answer
// @Summary 更新Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Answer true "更新Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /answer/updateAnswer [put]
func (answerApi *AnswerApi) UpdateAnswer(c *gin.Context) {
	var answer war.Answer
	err := c.ShouldBindJSON(&answer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "QuestionId":{utils.NotEmpty()},
          "Content":{utils.NotEmpty()},
          "IsCorrect":{utils.NotEmpty()},
      }
    if err := utils.Verify(answer, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := answerService.UpdateAnswer(answer); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAnswer 用id查询Answer
// @Tags Answer
// @Summary 用id查询Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.Answer true "用id查询Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /answer/findAnswer [get]
func (answerApi *AnswerApi) FindAnswer(c *gin.Context) {
	var answer war.Answer
	err := c.ShouldBindQuery(&answer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reanswer, err := answerService.GetAnswer(answer.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reanswer": reanswer}, c)
	}
}

// GetAnswerList 分页获取Answer列表
// @Tags Answer
// @Summary 分页获取Answer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.AnswerSearch true "分页获取Answer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /answer/getAnswerList [get]
func (answerApi *AnswerApi) GetAnswerList(c *gin.Context) {
	var pageInfo warReq.AnswerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := answerService.GetAnswerInfoList(pageInfo); err != nil {
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
