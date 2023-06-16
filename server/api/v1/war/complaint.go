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

type ComplaintApi struct {
}

var complaintService = service.ServiceGroupApp.WarServiceGroup.ComplaintService

// CreateComplaint 创建Complaint
// @Tags Complaint
// @Summary 创建Complaint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Complaint true "创建Complaint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /complaint/createComplaint [post]
func (complaintApi *ComplaintApi) CreateComplaint(c *gin.Context) {
	var complaint war.Complaint
	err := c.ShouldBindJSON(&complaint)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"RoomId":     {utils.NotEmpty()},
		"Complainee": {utils.NotEmpty()},
	}
	if err := utils.Verify(complaint, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := complaintService.CreateComplaint(&complaint); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteComplaint 删除Complaint
// @Tags Complaint
// @Summary 删除Complaint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Complaint true "删除Complaint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /complaint/deleteComplaint [delete]
func (complaintApi *ComplaintApi) DeleteComplaint(c *gin.Context) {
	var complaint war.Complaint
	err := c.ShouldBindJSON(&complaint)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := complaintService.DeleteComplaint(complaint); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteComplaintByIds 批量删除Complaint
// @Tags Complaint
// @Summary 批量删除Complaint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Complaint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /complaint/deleteComplaintByIds [delete]
func (complaintApi *ComplaintApi) DeleteComplaintByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := complaintService.DeleteComplaintByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateComplaint 更新Complaint
// @Tags Complaint
// @Summary 更新Complaint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Complaint true "更新Complaint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /complaint/updateComplaint [put]
func (complaintApi *ComplaintApi) UpdateComplaint(c *gin.Context) {
	var complaint war.Complaint
	err := c.ShouldBindJSON(&complaint)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"RoomId":     {utils.NotEmpty()},
		"Complainee": {utils.NotEmpty()},
	}
	if err := utils.Verify(complaint, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := complaintService.UpdateComplaint(complaint); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindComplaint 用id查询Complaint
// @Tags Complaint
// @Summary 用id查询Complaint
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.Complaint true "用id查询Complaint"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /complaint/findComplaint [get]
func (complaintApi *ComplaintApi) FindComplaint(c *gin.Context) {
	var complaint war.Complaint
	err := c.ShouldBindQuery(&complaint)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recomplaint, err := complaintService.GetComplaint(complaint.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recomplaint": recomplaint}, c)
	}
}

// GetComplaintList 分页获取Complaint列表
// @Tags Complaint
// @Summary 分页获取Complaint列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.ComplaintSearch true "分页获取Complaint列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /complaint/getComplaintList [get]
func (complaintApi *ComplaintApi) GetComplaintList(c *gin.Context) {
	var pageInfo warReq.ComplaintSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := complaintService.GetComplaintInfoList(pageInfo); err != nil {
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

// 会员投诉
func (complaintApi *ComplaintApi) MemberComplaint(c *gin.Context) {
	var complaint warReq.ComplaintInfo
	err := c.ShouldBindJSON(&complaint)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	verify := utils.Rules{
		"RoomId":     {utils.NotEmpty()},
		"Complainee": {utils.NotEmpty()},
	}
	if err := utils.Verify(complaint, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := complaintService.MemberComplaint(&complaint, userId); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
