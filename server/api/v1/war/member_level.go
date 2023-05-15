package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/war"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type MemberLevelApi struct {
}

var memberLevelService = service.ServiceGroupApp.WarServiceGroup.MemberLevelService


// CreateMemberLevel 创建MemberLevel
// @Tags MemberLevel
// @Summary 创建MemberLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.MemberLevel true "创建MemberLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memberLevel/createMemberLevel [post]
func (memberLevelApi *MemberLevelApi) CreateMemberLevel(c *gin.Context) {
	var memberLevel war.MemberLevel
	err := c.ShouldBindJSON(&memberLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Status":{utils.NotEmpty()},
    }
	if err := utils.Verify(memberLevel, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := memberLevelService.CreateMemberLevel(&memberLevel); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMemberLevel 删除MemberLevel
// @Tags MemberLevel
// @Summary 删除MemberLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.MemberLevel true "删除MemberLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memberLevel/deleteMemberLevel [delete]
func (memberLevelApi *MemberLevelApi) DeleteMemberLevel(c *gin.Context) {
	var memberLevel war.MemberLevel
	err := c.ShouldBindJSON(&memberLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := memberLevelService.DeleteMemberLevel(memberLevel); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMemberLevelByIds 批量删除MemberLevel
// @Tags MemberLevel
// @Summary 批量删除MemberLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemberLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /memberLevel/deleteMemberLevelByIds [delete]
func (memberLevelApi *MemberLevelApi) DeleteMemberLevelByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := memberLevelService.DeleteMemberLevelByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMemberLevel 更新MemberLevel
// @Tags MemberLevel
// @Summary 更新MemberLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.MemberLevel true "更新MemberLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memberLevel/updateMemberLevel [put]
func (memberLevelApi *MemberLevelApi) UpdateMemberLevel(c *gin.Context) {
	var memberLevel war.MemberLevel
	err := c.ShouldBindJSON(&memberLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Status":{utils.NotEmpty()},
      }
    if err := utils.Verify(memberLevel, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := memberLevelService.UpdateMemberLevel(memberLevel); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMemberLevel 用id查询MemberLevel
// @Tags MemberLevel
// @Summary 用id查询MemberLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.MemberLevel true "用id查询MemberLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memberLevel/findMemberLevel [get]
func (memberLevelApi *MemberLevelApi) FindMemberLevel(c *gin.Context) {
	var memberLevel war.MemberLevel
	err := c.ShouldBindQuery(&memberLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rememberLevel, err := memberLevelService.GetMemberLevel(memberLevel.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rememberLevel": rememberLevel}, c)
	}
}

// GetMemberLevelList 分页获取MemberLevel列表
// @Tags MemberLevel
// @Summary 分页获取MemberLevel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.MemberLevelSearch true "分页获取MemberLevel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memberLevel/getMemberLevelList [get]
func (memberLevelApi *MemberLevelApi) GetMemberLevelList(c *gin.Context) {
	var pageInfo warReq.MemberLevelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := memberLevelService.GetMemberLevelInfoList(pageInfo); err != nil {
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
