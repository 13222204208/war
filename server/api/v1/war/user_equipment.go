package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserEquipmentApi struct {
}

var userEquipmentService = service.ServiceGroupApp.WarServiceGroup.UserEquipmentService

// CreateUserEquipment 创建UserEquipment
// @Tags UserEquipment
// @Summary 创建UserEquipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.UserEquipment true "创建UserEquipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userEquipment/createUserEquipment [post]
func (userEquipmentApi *UserEquipmentApi) CreateUserEquipment(c *gin.Context) {
	var userEquipment warReq.UserEquipmentAdd
	err := c.ShouldBindJSON(&userEquipment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userEquipmentService.CreateUserEquipment(userEquipment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserEquipment 删除UserEquipment
// @Tags UserEquipment
// @Summary 删除UserEquipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.UserEquipment true "删除UserEquipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userEquipment/deleteUserEquipment [delete]
func (userEquipmentApi *UserEquipmentApi) DeleteUserEquipment(c *gin.Context) {
	var userEquipment war.UserEquipment
	err := c.ShouldBindJSON(&userEquipment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userEquipmentService.DeleteUserEquipment(userEquipment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserEquipmentByIds 批量删除UserEquipment
// @Tags UserEquipment
// @Summary 批量删除UserEquipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserEquipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userEquipment/deleteUserEquipmentByIds [delete]
func (userEquipmentApi *UserEquipmentApi) DeleteUserEquipmentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userEquipmentService.DeleteUserEquipmentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserEquipment 更新UserEquipment
// @Tags UserEquipment
// @Summary 更新UserEquipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.UserEquipment true "更新UserEquipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userEquipment/updateUserEquipment [put]
func (userEquipmentApi *UserEquipmentApi) UpdateUserEquipment(c *gin.Context) {
	var userEquipment war.UserEquipment
	err := c.ShouldBindJSON(&userEquipment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userEquipmentService.UpdateUserEquipment(userEquipment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserEquipment 用id查询UserEquipment
// @Tags UserEquipment
// @Summary 用id查询UserEquipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.UserEquipment true "用id查询UserEquipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userEquipment/findUserEquipment [get]
func (userEquipmentApi *UserEquipmentApi) FindUserEquipment(c *gin.Context) {
	var userEquipment war.UserEquipment
	err := c.ShouldBindQuery(&userEquipment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reuserEquipment, err := userEquipmentService.GetUserEquipment(userEquipment.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserEquipment": reuserEquipment}, c)
	}
}

// GetUserEquipmentList 分页获取UserEquipment列表
// @Tags UserEquipment
// @Summary 分页获取UserEquipment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.UserEquipmentSearch true "分页获取UserEquipment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userEquipment/getUserEquipmentList [get]
func (userEquipmentApi *UserEquipmentApi) GetUserEquipmentList(c *gin.Context) {
	var pageInfo warReq.UserEquipmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// if list, total, err := userEquipmentService.GetUserEquipmentInfoList(pageInfo); err != nil {
	//     global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//     response.FailWithMessage("获取失败", c)
	// } else {
	//     response.OkWithDetailed(response.PageResult{
	//         List:     list,
	//         Total:    total,
	//         Page:     pageInfo.Page,
	//         PageSize: pageInfo.PageSize,
	//     }, "获取成功", c)
	// }
	if list, err := userEquipmentService.GetUserEquipmentInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(len(list)),
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
