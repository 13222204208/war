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

type TeamApi struct {
}

var teamService = service.ServiceGroupApp.WarServiceGroup.TeamService

// 创建战队
func (teamApi *TeamApi) CreateTeam(c *gin.Context) {
	userID := utils.GetUserID(c)
	var team war.Team
	err := c.ShouldBindJSON(&team)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//判断战队名称是否有空格，和大于36个字符 一个中文字符占3个字符
	if team.Name == "" || len(team.Name) > 21 {
		response.FailWithMessage("战队名称长度不合法", c)
		return
	}
	//判断战队简介是否有空格，和大于36个字符
	if team.Description == "" || len(team.Description) > 36 {
		response.FailWithMessage("战队简介长度不合法", c)
		return
	}
	//获取解析后的token用户id
	team.LeaderId = &userID
	if err := teamService.CreateTeam(&team); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// 修改战队信息
func (teamApi *TeamApi) UpdateTeamInfo(c *gin.Context) {
	var team war.Team
	err := c.ShouldBindJSON(&team)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//判断战队名称是否有空格，和大于21个字符 一个中文字符占3个字符
	if team.Name == "" || len(team.Name) > 21 {
		response.FailWithMessage("战队名称长度不合法", c)
		return
	}
	//判断战队简介是否有空格，和大于36个字符
	if team.Description == "" || len(team.Description) > 36 {
		response.FailWithMessage("战队简介长度不合法", c)
		return
	}
	//获取解析后的token用户id
	userID := utils.GetUserID(c)
	team.LeaderId = &userID
	if err := teamService.UpdateTeamInfo(&team); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// 获取全险战队信息，显示队长信息 ，并按战队积分从大到小排序
func (teamApi *TeamApi) GetAllTeam(c *gin.Context) {
	if list, err := teamService.GetAllTeam(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(list, c)
	}
}

// 获取战队详情
func (teamApi *TeamApi) GetTeamDetail(c *gin.Context) {
	id := c.Param("id")
	//转为Uint
	if team, err := teamService.GetTeamDetail(id); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(team, c)
	}
}

// 获取我的战队详情
func (teamApi *TeamApi) GetMyTeamDetail(c *gin.Context) {
	userID := utils.GetUserID(c)
	if team, err := teamService.GetMyTeamDetail(userID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(team, c)
	}
}

// CreateTeam 创建Team
// @Tags Team
// @Summary 创建Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Team true "创建Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /team/createTeam [post]
// func (teamApi *TeamApi) CreateTeam(c *gin.Context) {
// 	var team war.Team
// 	err := c.ShouldBindJSON(&team)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err := teamService.CreateTeam(&team); err != nil {
//         global.GVA_LOG.Error("创建失败!", zap.Error(err))
// 		response.FailWithMessage("创建失败", c)
// 	} else {
// 		response.OkWithMessage("创建成功", c)
// 	}
// }

// DeleteTeam 删除Team
// @Tags Team
// @Summary 删除Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Team true "删除Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /team/deleteTeam [delete]
func (teamApi *TeamApi) DeleteTeam(c *gin.Context) {
	var team war.Team
	err := c.ShouldBindJSON(&team)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamService.DeleteTeam(team); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeamByIds 批量删除Team
// @Tags Team
// @Summary 批量删除Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /team/deleteTeamByIds [delete]
func (teamApi *TeamApi) DeleteTeamByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamService.DeleteTeamByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeam 更新Team
// @Tags Team
// @Summary 更新Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Team true "更新Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /team/updateTeam [put]
func (teamApi *TeamApi) UpdateTeam(c *gin.Context) {
	var team war.Team
	err := c.ShouldBindJSON(&team)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teamService.UpdateTeam(team); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeam 用id查询Team
// @Tags Team
// @Summary 用id查询Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.Team true "用id查询Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /team/findTeam [get]
func (teamApi *TeamApi) FindTeam(c *gin.Context) {
	var team war.Team
	err := c.ShouldBindQuery(&team)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reteam, err := teamService.GetTeam(team.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteam": reteam}, c)
	}
}

// GetTeamList 分页获取Team列表
// @Tags Team
// @Summary 分页获取Team列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.TeamSearch true "分页获取Team列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /team/getTeamList [get]
func (teamApi *TeamApi) GetTeamList(c *gin.Context) {
	var pageInfo warReq.TeamSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := teamService.GetTeamInfoList(pageInfo); err != nil {
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
