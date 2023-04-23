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

type TeamMemberApi struct {
}

var teamMemberService = service.ServiceGroupApp.WarServiceGroup.TeamMemberService

// CreateTeamMember 创建TeamMember
// @Tags TeamMember
// @Summary 创建TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.TeamMember true "创建TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamMember/createTeamMember [post]
func (teamMemberApi *TeamMemberApi) CreateTeamMember(c *gin.Context) {
	var teamMember war.TeamMember
	err := c.ShouldBindJSON(&teamMember)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamMemberService.CreateTeamMember(&teamMember); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeamMember 删除TeamMember
// @Tags TeamMember
// @Summary 删除TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.TeamMember true "删除TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamMember/deleteTeamMember [delete]
func (teamMemberApi *TeamMemberApi) DeleteTeamMember(c *gin.Context) {
	var teamMember war.TeamMember
	err := c.ShouldBindJSON(&teamMember)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamMemberService.DeleteTeamMember(teamMember); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeamMemberByIds 批量删除TeamMember
// @Tags TeamMember
// @Summary 批量删除TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teamMember/deleteTeamMemberByIds [delete]
func (teamMemberApi *TeamMemberApi) DeleteTeamMemberByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamMemberService.DeleteTeamMemberByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeamMember 更新TeamMember
// @Tags TeamMember
// @Summary 更新TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.TeamMember true "更新TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamMember/updateTeamMember [put]
func (teamMemberApi *TeamMemberApi) UpdateTeamMember(c *gin.Context) {
	var teamMember war.TeamMember
	err := c.ShouldBindJSON(&teamMember)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamMemberService.UpdateTeamMember(teamMember); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeamMember 用id查询TeamMember
// @Tags TeamMember
// @Summary 用id查询TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.TeamMember true "用id查询TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamMember/findTeamMember [get]
func (teamMemberApi *TeamMemberApi) FindTeamMember(c *gin.Context) {
	var teamMember war.TeamMember
	err := c.ShouldBindQuery(&teamMember)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reteamMember, err := teamMemberService.GetTeamMember(teamMember.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteamMember": reteamMember}, c)
	}
}

// GetTeamMemberList 分页获取TeamMember列表
// @Tags TeamMember
// @Summary 分页获取TeamMember列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.TeamMemberSearch true "分页获取TeamMember列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamMember/getTeamMemberList [get]
func (teamMemberApi *TeamMemberApi) GetTeamMemberList(c *gin.Context) {
	var pageInfo warReq.TeamMemberSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := teamMemberService.GetTeamMemberInfoList(pageInfo); err != nil {
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

// 更改会员的角色
func (teamMemberApi *TeamMemberApi) UpdateTeamMemberRole(c *gin.Context) {
	userID := utils.GetUserID(c)
	var t warReq.UpdateTeamMemberRole
	err := c.ShouldBindJSON(&t)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamMemberService.UpdateTeamMemberRole(userID, t); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
