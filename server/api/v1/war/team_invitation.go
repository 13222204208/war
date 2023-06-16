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

type TeamInvitationApi struct {
}

var teamInvitationService = service.ServiceGroupApp.WarServiceGroup.TeamInvitationService

// CreateTeamInvitation 创建TeamInvitation
// @Tags TeamInvitation
// @Summary 创建TeamInvitation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.TeamInvitation true "创建TeamInvitation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamInvitation/createTeamInvitation [post]
func (teamInvitationApi *TeamInvitationApi) CreateTeamInvitation(c *gin.Context) {
	var teamInvitation war.TeamInvitation
	err := c.ShouldBindJSON(&teamInvitation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"TeamId":               {utils.NotEmpty()},
		"TeamCaptainId":        {utils.NotEmpty()},
		"Date":                 {utils.NotEmpty()},
		"InvitedTeamId":        {utils.NotEmpty()},
		"InvitedTeamCaptainId": {utils.NotEmpty()},
	}
	if err := utils.Verify(teamInvitation, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamInvitationService.CreateTeamInvitation(&teamInvitation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeamInvitation 删除TeamInvitation
// @Tags TeamInvitation
// @Summary 删除TeamInvitation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.TeamInvitation true "删除TeamInvitation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamInvitation/deleteTeamInvitation [delete]
func (teamInvitationApi *TeamInvitationApi) DeleteTeamInvitation(c *gin.Context) {
	var teamInvitation war.TeamInvitation
	err := c.ShouldBindJSON(&teamInvitation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamInvitationService.DeleteTeamInvitation(teamInvitation); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeamInvitationByIds 批量删除TeamInvitation
// @Tags TeamInvitation
// @Summary 批量删除TeamInvitation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamInvitation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teamInvitation/deleteTeamInvitationByIds [delete]
func (teamInvitationApi *TeamInvitationApi) DeleteTeamInvitationByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamInvitationService.DeleteTeamInvitationByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeamInvitation 更新TeamInvitation
// @Tags TeamInvitation
// @Summary 更新TeamInvitation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.TeamInvitation true "更新TeamInvitation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamInvitation/updateTeamInvitation [put]
func (teamInvitationApi *TeamInvitationApi) UpdateTeamInvitation(c *gin.Context) {
	var teamInvitation war.TeamInvitation
	err := c.ShouldBindJSON(&teamInvitation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"TeamId":               {utils.NotEmpty()},
		"TeamCaptainId":        {utils.NotEmpty()},
		"Date":                 {utils.NotEmpty()},
		"InvitedTeamId":        {utils.NotEmpty()},
		"InvitedTeamCaptainId": {utils.NotEmpty()},
	}
	if err := utils.Verify(teamInvitation, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamInvitationService.UpdateTeamInvitation(teamInvitation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeamInvitation 用id查询TeamInvitation
// @Tags TeamInvitation
// @Summary 用id查询TeamInvitation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.TeamInvitation true "用id查询TeamInvitation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamInvitation/findTeamInvitation [get]
func (teamInvitationApi *TeamInvitationApi) FindTeamInvitation(c *gin.Context) {
	var teamInvitation war.TeamInvitation
	err := c.ShouldBindQuery(&teamInvitation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reteamInvitation, err := teamInvitationService.GetTeamInvitation(teamInvitation.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteamInvitation": reteamInvitation}, c)
	}
}

// GetTeamInvitationList 分页获取TeamInvitation列表
// @Tags TeamInvitation
// @Summary 分页获取TeamInvitation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.TeamInvitationSearch true "分页获取TeamInvitation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamInvitation/getTeamInvitationList [get]
func (teamInvitationApi *TeamInvitationApi) GetTeamInvitationList(c *gin.Context) {
	var pageInfo warReq.TeamInvitationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := teamInvitationService.GetTeamInvitationInfoList(pageInfo); err != nil {
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

// 战队邀请
func (teamInvitationApi *TeamInvitationApi) TeamInvitation(c *gin.Context) {
	var teamInvitation war.TeamInvitation
	err := c.ShouldBindJSON(&teamInvitation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	teamInvitation.TeamCaptainId = userID
	if err := teamInvitationService.TeamInvitation(teamInvitation); err != nil {
		global.GVA_LOG.Error("邀请失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("邀请成功", c)
	}
}

// 战队受邀列表
func (teamInvitationApi *TeamInvitationApi) TeamInvitationList(c *gin.Context) {
	userID := utils.GetUserID(c)
	if list, err := teamInvitationService.TeamInvitationList(userID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}
