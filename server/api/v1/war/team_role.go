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

type TeamRoleApi struct {
}

var teamRoleService = service.ServiceGroupApp.WarServiceGroup.TeamRoleService

// CreateTeamRole 创建TeamRole
// @Tags TeamRole
// @Summary 创建TeamRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.TeamRole true "创建TeamRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamRole/createTeamRole [post]
func (teamRoleApi *TeamRoleApi) CreateTeamRole(c *gin.Context) {
	var teamRole war.TeamRole
	err := c.ShouldBindJSON(&teamRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Role":   {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(teamRole, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamRoleService.CreateTeamRole(&teamRole); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeamRole 删除TeamRole
// @Tags TeamRole
// @Summary 删除TeamRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.TeamRole true "删除TeamRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamRole/deleteTeamRole [delete]
func (teamRoleApi *TeamRoleApi) DeleteTeamRole(c *gin.Context) {
	var teamRole war.TeamRole
	err := c.ShouldBindJSON(&teamRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamRoleService.DeleteTeamRole(teamRole); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeamRoleByIds 批量删除TeamRole
// @Tags TeamRole
// @Summary 批量删除TeamRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teamRole/deleteTeamRoleByIds [delete]
func (teamRoleApi *TeamRoleApi) DeleteTeamRoleByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamRoleService.DeleteTeamRoleByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeamRole 更新TeamRole
// @Tags TeamRole
// @Summary 更新TeamRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.TeamRole true "更新TeamRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamRole/updateTeamRole [put]
func (teamRoleApi *TeamRoleApi) UpdateTeamRole(c *gin.Context) {
	var teamRole war.TeamRole
	err := c.ShouldBindJSON(&teamRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Role":   {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(teamRole, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamRoleService.UpdateTeamRole(teamRole); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeamRole 用id查询TeamRole
// @Tags TeamRole
// @Summary 用id查询TeamRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.TeamRole true "用id查询TeamRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamRole/findTeamRole [get]
func (teamRoleApi *TeamRoleApi) FindTeamRole(c *gin.Context) {
	var teamRole war.TeamRole
	err := c.ShouldBindQuery(&teamRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reteamRole, err := teamRoleService.GetTeamRole(teamRole.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteamRole": reteamRole}, c)
	}
}

// GetTeamRoleList 分页获取TeamRole列表
// @Tags TeamRole
// @Summary 分页获取TeamRole列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.TeamRoleSearch true "分页获取TeamRole列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamRole/getTeamRoleList [get]
func (teamRoleApi *TeamRoleApi) GetTeamRoleList(c *gin.Context) {
	var pageInfo warReq.TeamRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := teamRoleService.GetTeamRoleInfoList(pageInfo); err != nil {
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

// 获取战队角色列表
func (teamRoleApi *TeamRoleApi) TeamRoleList(c *gin.Context) {

	if list, err := teamRoleService.GetTeamRoleList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"list": list}, "获取成功", c)
	}
}
