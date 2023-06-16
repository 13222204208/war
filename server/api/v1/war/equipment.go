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

type EquipmentApi struct {
}

var equipmentService = service.ServiceGroupApp.WarServiceGroup.EquipmentService

// CreateEquipment 创建Equipment
// @Tags Equipment
// @Summary 创建Equipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Equipment true "创建Equipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /equipment/createEquipment [post]
func (equipmentApi *EquipmentApi) CreateEquipment(c *gin.Context) {
	var equipment war.Equipment
	err := c.ShouldBindJSON(&equipment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(equipment, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := equipmentService.CreateEquipment(&equipment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEquipment 删除Equipment
// @Tags Equipment
// @Summary 删除Equipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Equipment true "删除Equipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /equipment/deleteEquipment [delete]
func (equipmentApi *EquipmentApi) DeleteEquipment(c *gin.Context) {
	var equipment war.Equipment
	err := c.ShouldBindJSON(&equipment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := equipmentService.DeleteEquipment(equipment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEquipmentByIds 批量删除Equipment
// @Tags Equipment
// @Summary 批量删除Equipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Equipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /equipment/deleteEquipmentByIds [delete]
func (equipmentApi *EquipmentApi) DeleteEquipmentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := equipmentService.DeleteEquipmentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEquipment 更新Equipment
// @Tags Equipment
// @Summary 更新Equipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Equipment true "更新Equipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /equipment/updateEquipment [put]
func (equipmentApi *EquipmentApi) UpdateEquipment(c *gin.Context) {
	var equipment war.Equipment
	err := c.ShouldBindJSON(&equipment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(equipment, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := equipmentService.UpdateEquipment(equipment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEquipment 用id查询Equipment
// @Tags Equipment
// @Summary 用id查询Equipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.Equipment true "用id查询Equipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /equipment/findEquipment [get]
func (equipmentApi *EquipmentApi) FindEquipment(c *gin.Context) {
	var equipment war.Equipment
	err := c.ShouldBindQuery(&equipment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reequipment, err := equipmentService.GetEquipment(equipment.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reequipment": reequipment}, c)
	}
}

// GetEquipmentList 分页获取Equipment列表
// @Tags Equipment
// @Summary 分页获取Equipment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.EquipmentSearch true "分页获取Equipment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /equipment/getEquipmentList [get]
func (equipmentApi *EquipmentApi) GetEquipmentList(c *gin.Context) {
	var pageInfo warReq.EquipmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := equipmentService.GetEquipmentInfoList(pageInfo); err != nil {
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

// 我的装备详情
func (equipmentApi *EquipmentApi) Detail(c *gin.Context) {
	userId := utils.GetUserID(c)
	if reequipment, err := equipmentService.Detail(userId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"info": reequipment}, c)
	}
}
