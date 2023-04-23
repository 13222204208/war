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
)

type MatchRecordApi struct {
}

var matchRecordService = service.ServiceGroupApp.WarServiceGroup.MatchRecordService


// CreateMatchRecord 创建MatchRecord
// @Tags MatchRecord
// @Summary 创建MatchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.MatchRecord true "创建MatchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /matchRecord/createMatchRecord [post]
func (matchRecordApi *MatchRecordApi) CreateMatchRecord(c *gin.Context) {
	var matchRecord war.MatchRecord
	err := c.ShouldBindJSON(&matchRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := matchRecordService.CreateMatchRecord(&matchRecord); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMatchRecord 删除MatchRecord
// @Tags MatchRecord
// @Summary 删除MatchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.MatchRecord true "删除MatchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /matchRecord/deleteMatchRecord [delete]
func (matchRecordApi *MatchRecordApi) DeleteMatchRecord(c *gin.Context) {
	var matchRecord war.MatchRecord
	err := c.ShouldBindJSON(&matchRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := matchRecordService.DeleteMatchRecord(matchRecord); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMatchRecordByIds 批量删除MatchRecord
// @Tags MatchRecord
// @Summary 批量删除MatchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MatchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /matchRecord/deleteMatchRecordByIds [delete]
func (matchRecordApi *MatchRecordApi) DeleteMatchRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := matchRecordService.DeleteMatchRecordByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMatchRecord 更新MatchRecord
// @Tags MatchRecord
// @Summary 更新MatchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.MatchRecord true "更新MatchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /matchRecord/updateMatchRecord [put]
func (matchRecordApi *MatchRecordApi) UpdateMatchRecord(c *gin.Context) {
	var matchRecord war.MatchRecord
	err := c.ShouldBindJSON(&matchRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := matchRecordService.UpdateMatchRecord(matchRecord); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMatchRecord 用id查询MatchRecord
// @Tags MatchRecord
// @Summary 用id查询MatchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.MatchRecord true "用id查询MatchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /matchRecord/findMatchRecord [get]
func (matchRecordApi *MatchRecordApi) FindMatchRecord(c *gin.Context) {
	var matchRecord war.MatchRecord
	err := c.ShouldBindQuery(&matchRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rematchRecord, err := matchRecordService.GetMatchRecord(matchRecord.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rematchRecord": rematchRecord}, c)
	}
}

// GetMatchRecordList 分页获取MatchRecord列表
// @Tags MatchRecord
// @Summary 分页获取MatchRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.MatchRecordSearch true "分页获取MatchRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /matchRecord/getMatchRecordList [get]
func (matchRecordApi *MatchRecordApi) GetMatchRecordList(c *gin.Context) {
	var pageInfo warReq.MatchRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := matchRecordService.GetMatchRecordInfoList(pageInfo); err != nil {
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
