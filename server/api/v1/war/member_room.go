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

type MemberRoomApi struct {
}

var memberRoomService = service.ServiceGroupApp.WarServiceGroup.MemberRoomService

// CreateMemberRoom 创建MemberRoom
// @Tags MemberRoom
// @Summary 创建MemberRoom
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.MemberRoom true "创建MemberRoom"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memberRoom/createMemberRoom [post]
func (memberRoomApi *MemberRoomApi) CreateMemberRoom(c *gin.Context) {
	var memberRoom war.MemberRoom
	err := c.ShouldBindJSON(&memberRoom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := memberRoomService.CreateMemberRoom(&memberRoom); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMemberRoom 删除MemberRoom
// @Tags MemberRoom
// @Summary 删除MemberRoom
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.MemberRoom true "删除MemberRoom"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memberRoom/deleteMemberRoom [delete]
func (memberRoomApi *MemberRoomApi) DeleteMemberRoom(c *gin.Context) {
	var memberRoom war.MemberRoom
	err := c.ShouldBindJSON(&memberRoom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := memberRoomService.DeleteMemberRoom(memberRoom); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMemberRoomByIds 批量删除MemberRoom
// @Tags MemberRoom
// @Summary 批量删除MemberRoom
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemberRoom"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /memberRoom/deleteMemberRoomByIds [delete]
func (memberRoomApi *MemberRoomApi) DeleteMemberRoomByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := memberRoomService.DeleteMemberRoomByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMemberRoom 更新MemberRoom
// @Tags MemberRoom
// @Summary 更新MemberRoom
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.MemberRoom true "更新MemberRoom"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memberRoom/updateMemberRoom [put]
func (memberRoomApi *MemberRoomApi) UpdateMemberRoom(c *gin.Context) {
	var memberRoom war.MemberRoom
	err := c.ShouldBindJSON(&memberRoom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := memberRoomService.UpdateMemberRoom(memberRoom); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMemberRoom 用id查询MemberRoom
// @Tags MemberRoom
// @Summary 用id查询MemberRoom
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.MemberRoom true "用id查询MemberRoom"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memberRoom/findMemberRoom [get]
func (memberRoomApi *MemberRoomApi) FindMemberRoom(c *gin.Context) {
	var memberRoom war.MemberRoom
	err := c.ShouldBindQuery(&memberRoom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rememberRoom, err := memberRoomService.GetMemberRoom(memberRoom.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rememberRoom": rememberRoom}, c)
	}
}

// GetMemberRoomList 分页获取MemberRoom列表
// @Tags MemberRoom
// @Summary 分页获取MemberRoom列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.MemberRoomSearch true "分页获取MemberRoom列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memberRoom/getMemberRoomList [get]
func (memberRoomApi *MemberRoomApi) GetMemberRoomList(c *gin.Context) {
	var pageInfo warReq.MemberRoomSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := memberRoomService.GetMemberRoomInfoList(pageInfo); err != nil {
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

// 用户签到
func (memberRoomApi *MemberRoomApi) SignIn(c *gin.Context) {
	userId := utils.GetUserID(c)
	if err := memberRoomService.Sign(userId); err != nil {
		global.GVA_LOG.Error("签到失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("签到成功", c)
	}
}

// 对局列表
func (memberRoomApi *MemberRoomApi) GetMemberRoomListByUserId(c *gin.Context) {
	userId := utils.GetUserID(c)
	if list, err := memberRoomService.GetMemberRoomListByUserId(userId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}

// 对局详情
func (memberRoomApi *MemberRoomApi) GetMemberRoomDetailByRoomId(c *gin.Context) {
	roomId := c.Query("roomId")
	//转为int
	roomIdInt, _ := strconv.Atoi(roomId)
	if roomIdInt == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	if memberRoomInfo, err := memberRoomService.GetMemberRoomInfoByRoomId(roomIdInt); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"memberRoomInfo": memberRoomInfo}, c)
	}
}
