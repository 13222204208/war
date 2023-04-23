package war

import (
	"strconv"

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

type BackgroundsApi struct {
}

var backgroundsService = service.ServiceGroupApp.WarServiceGroup.BackgroundsService

// CreateBackgrounds 创建Backgrounds
// @Tags Backgrounds
// @Summary 创建Backgrounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Backgrounds true "创建Backgrounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /backgrounds/createBackgrounds [post]
func (backgroundsApi *BackgroundsApi) CreateBackgrounds(c *gin.Context) {
	var backgrounds war.Backgrounds
	err := c.ShouldBindJSON(&backgrounds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Url":    {utils.NotEmpty()},
		"Type":   {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(backgrounds, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := backgroundsService.CreateBackgrounds(&backgrounds); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBackgrounds 删除Backgrounds
// @Tags Backgrounds
// @Summary 删除Backgrounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Backgrounds true "删除Backgrounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /backgrounds/deleteBackgrounds [delete]
func (backgroundsApi *BackgroundsApi) DeleteBackgrounds(c *gin.Context) {
	var backgrounds war.Backgrounds
	err := c.ShouldBindJSON(&backgrounds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := backgroundsService.DeleteBackgrounds(backgrounds); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBackgroundsByIds 批量删除Backgrounds
// @Tags Backgrounds
// @Summary 批量删除Backgrounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Backgrounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /backgrounds/deleteBackgroundsByIds [delete]
func (backgroundsApi *BackgroundsApi) DeleteBackgroundsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := backgroundsService.DeleteBackgroundsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBackgrounds 更新Backgrounds
// @Tags Backgrounds
// @Summary 更新Backgrounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Backgrounds true "更新Backgrounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /backgrounds/updateBackgrounds [put]
func (backgroundsApi *BackgroundsApi) UpdateBackgrounds(c *gin.Context) {
	var backgrounds war.Backgrounds
	err := c.ShouldBindJSON(&backgrounds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Url":    {utils.NotEmpty()},
		"Type":   {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(backgrounds, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := backgroundsService.UpdateBackgrounds(backgrounds); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBackgrounds 用id查询Backgrounds
// @Tags Backgrounds
// @Summary 用id查询Backgrounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.Backgrounds true "用id查询Backgrounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /backgrounds/findBackgrounds [get]
func (backgroundsApi *BackgroundsApi) FindBackgrounds(c *gin.Context) {
	var backgrounds war.Backgrounds
	err := c.ShouldBindQuery(&backgrounds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebackgrounds, err := backgroundsService.GetBackgrounds(backgrounds.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebackgrounds": rebackgrounds}, c)
	}
}

// GetBackgroundsList 分页获取Backgrounds列表
// @Tags Backgrounds
// @Summary 分页获取Backgrounds列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.BackgroundsSearch true "分页获取Backgrounds列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /backgrounds/getBackgroundsList [get]
func (backgroundsApi *BackgroundsApi) GetBackgroundsList(c *gin.Context) {
	var pageInfo warReq.BackgroundsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := backgroundsService.GetBackgroundsInfoList(pageInfo); err != nil {
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

// 根据类型获取背景图
func (backgroundsApi *BackgroundsApi) GetBackgroundsByType(c *gin.Context) {
	t, err := strconv.Atoi(c.Param("type"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, err := backgroundsService.GetBackgroundsByType(t); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}
