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

type AnnouncementApi struct {
}

var announcementService = service.ServiceGroupApp.WarServiceGroup.AnnouncementService

// CreateAnnouncement 创建Announcement
// @Tags Announcement
// @Summary 创建Announcement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Announcement true "创建Announcement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /announcement/createAnnouncement [post]
func (announcementApi *AnnouncementApi) CreateAnnouncement(c *gin.Context) {
	var announcement war.Announcement
	err := c.ShouldBindJSON(&announcement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Content": {utils.NotEmpty()},
		"Type":    {utils.NotEmpty()},
		"Status":  {utils.NotEmpty()},
	}
	if err := utils.Verify(announcement, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := announcementService.CreateAnnouncement(&announcement); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAnnouncement 删除Announcement
// @Tags Announcement
// @Summary 删除Announcement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Announcement true "删除Announcement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /announcement/deleteAnnouncement [delete]
func (announcementApi *AnnouncementApi) DeleteAnnouncement(c *gin.Context) {
	var announcement war.Announcement
	err := c.ShouldBindJSON(&announcement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := announcementService.DeleteAnnouncement(announcement); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAnnouncementByIds 批量删除Announcement
// @Tags Announcement
// @Summary 批量删除Announcement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Announcement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /announcement/deleteAnnouncementByIds [delete]
func (announcementApi *AnnouncementApi) DeleteAnnouncementByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := announcementService.DeleteAnnouncementByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAnnouncement 更新Announcement
// @Tags Announcement
// @Summary 更新Announcement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Announcement true "更新Announcement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /announcement/updateAnnouncement [put]
func (announcementApi *AnnouncementApi) UpdateAnnouncement(c *gin.Context) {
	var announcement war.Announcement
	err := c.ShouldBindJSON(&announcement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Content": {utils.NotEmpty()},
		"Type":    {utils.NotEmpty()},
		"Status":  {utils.NotEmpty()},
	}
	if err := utils.Verify(announcement, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := announcementService.UpdateAnnouncement(announcement); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAnnouncement 用id查询Announcement
// @Tags Announcement
// @Summary 用id查询Announcement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.Announcement true "用id查询Announcement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /announcement/findAnnouncement [get]
func (announcementApi *AnnouncementApi) FindAnnouncement(c *gin.Context) {
	var announcement war.Announcement
	err := c.ShouldBindQuery(&announcement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reannouncement, err := announcementService.GetAnnouncement(announcement.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reannouncement": reannouncement}, c)
	}
}

// GetAnnouncementList 分页获取Announcement列表
// @Tags Announcement
// @Summary 分页获取Announcement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.AnnouncementSearch true "分页获取Announcement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /announcement/getAnnouncementList [get]
func (announcementApi *AnnouncementApi) GetAnnouncementList(c *gin.Context) {
	var pageInfo warReq.AnnouncementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := announcementService.GetAnnouncementInfoList(pageInfo); err != nil {
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

// 根据公告类型获取公告列表
func (announcementApi *AnnouncementApi) GetAnnouncementListByType(c *gin.Context) {
	t, err := strconv.Atoi(c.Param("type"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if info, err := announcementService.GetAnnouncementByType(t); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"info": info}, c)
	}
}
