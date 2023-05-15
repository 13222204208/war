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

type GameRecordApi struct {
}

var gameRecordService = service.ServiceGroupApp.WarServiceGroup.GameRecordService


// CreateGameRecord 创建GameRecord
// @Tags GameRecord
// @Summary 创建GameRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.GameRecord true "创建GameRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gameRecord/createGameRecord [post]
func (gameRecordApi *GameRecordApi) CreateGameRecord(c *gin.Context) {
	var gameRecord war.GameRecord
	err := c.ShouldBindJSON(&gameRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gameRecordService.CreateGameRecord(&gameRecord); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGameRecord 删除GameRecord
// @Tags GameRecord
// @Summary 删除GameRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.GameRecord true "删除GameRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gameRecord/deleteGameRecord [delete]
func (gameRecordApi *GameRecordApi) DeleteGameRecord(c *gin.Context) {
	var gameRecord war.GameRecord
	err := c.ShouldBindJSON(&gameRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gameRecordService.DeleteGameRecord(gameRecord); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGameRecordByIds 批量删除GameRecord
// @Tags GameRecord
// @Summary 批量删除GameRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GameRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /gameRecord/deleteGameRecordByIds [delete]
func (gameRecordApi *GameRecordApi) DeleteGameRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gameRecordService.DeleteGameRecordByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGameRecord 更新GameRecord
// @Tags GameRecord
// @Summary 更新GameRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.GameRecord true "更新GameRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gameRecord/updateGameRecord [put]
func (gameRecordApi *GameRecordApi) UpdateGameRecord(c *gin.Context) {
	var gameRecord war.GameRecord
	err := c.ShouldBindJSON(&gameRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gameRecordService.UpdateGameRecord(gameRecord); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGameRecord 用id查询GameRecord
// @Tags GameRecord
// @Summary 用id查询GameRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.GameRecord true "用id查询GameRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gameRecord/findGameRecord [get]
func (gameRecordApi *GameRecordApi) FindGameRecord(c *gin.Context) {
	var gameRecord war.GameRecord
	err := c.ShouldBindQuery(&gameRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if regameRecord, err := gameRecordService.GetGameRecord(gameRecord.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regameRecord": regameRecord}, c)
	}
}

// GetGameRecordList 分页获取GameRecord列表
// @Tags GameRecord
// @Summary 分页获取GameRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.GameRecordSearch true "分页获取GameRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gameRecord/getGameRecordList [get]
func (gameRecordApi *GameRecordApi) GetGameRecordList(c *gin.Context) {
	var pageInfo warReq.GameRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := gameRecordService.GetGameRecordInfoList(pageInfo); err != nil {
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
