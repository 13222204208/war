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

type RoomApi struct {
}

var roomService = service.ServiceGroupApp.WarServiceGroup.RoomService

// CreateRoom 创建Room
// @Tags Room
// @Summary 创建Room
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Room true "创建Room"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /room/createRoom [post]
func (roomApi *RoomApi) CreateRoom(c *gin.Context) {
	var room war.Room
	err := c.ShouldBindJSON(&room)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roomService.CreateRoom(&room); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRoom 删除Room
// @Tags Room
// @Summary 删除Room
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Room true "删除Room"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /room/deleteRoom [delete]
func (roomApi *RoomApi) DeleteRoom(c *gin.Context) {
	var room war.Room
	err := c.ShouldBindJSON(&room)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roomService.DeleteRoom(room); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRoomByIds 批量删除Room
// @Tags Room
// @Summary 批量删除Room
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Room"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /room/deleteRoomByIds [delete]
func (roomApi *RoomApi) DeleteRoomByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roomService.DeleteRoomByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRoom 更新Room
// @Tags Room
// @Summary 更新Room
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Room true "更新Room"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /room/updateRoom [put]
func (roomApi *RoomApi) UpdateRoom(c *gin.Context) {
	var room war.Room
	err := c.ShouldBindJSON(&room)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if room.Status == 4 {
		if room.RedScore+room.BlueScore != 5 {
			response.FailWithMessage("场次不对", c)
			return
		}
	}

	if err := roomService.UpdateRoom(room); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRoom 用id查询Room
// @Tags Room
// @Summary 用id查询Room
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.Room true "用id查询Room"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /room/findRoom [get]
func (roomApi *RoomApi) FindRoom(c *gin.Context) {
	var room war.Room
	err := c.ShouldBindQuery(&room)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reroom, err := roomService.GetRoom(room.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reroom": reroom}, c)
	}
}

// GetRoomList 分页获取Room列表
// @Tags Room
// @Summary 分页获取Room列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.RoomSearch true "分页获取Room列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /room/getRoomList [get]
func (roomApi *RoomApi) GetRoomList(c *gin.Context) {
	var pageInfo warReq.RoomSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := roomService.GetRoomInfoList(pageInfo); err != nil {
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

// 快速匹配游戏
func (roomApi *RoomApi) QuickMatch(c *gin.Context) {
	userId := utils.GetUserID(c)
	if err := roomService.QuickMatch(userId); err != nil {
		global.GVA_LOG.Error("快速匹配失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("快速匹配成功", c)
	}
}

// 开始游戏
func (roomApi *RoomApi) StartGame(c *gin.Context) {
	var room war.Room
	err := c.ShouldBindJSON(&room)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roomService.StartGame(room.ID, 2); err != nil {
		global.GVA_LOG.Error("开始游戏失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("开始游戏成功", c)
	}
}

// 倒计时结束解散或开始游戏
func (roomApi *RoomApi) Countdown(c *gin.Context) {

	if err := roomService.Countdown(); err != nil {
		global.GVA_LOG.Error("倒计时结束失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("倒计时结束成功", c)
	}
}

// 获取房间对战海报二维码
func (roomApi *RoomApi) GetRoomQrCode(c *gin.Context) {
	var room warReq.RoomIdReq
	err := c.ShouldBindQuery(&room)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if room.Id == 0 {
		response.FailWithMessage("房间id不能为空", c)
		return
	}
	if url, err := roomService.GetRoomQrCode(room.Id); err != nil {
		global.GVA_LOG.Error("获取房间二维码失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"url": url}, c)
	}
}
