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

type RankLevelApi struct {
}

var rankLevelService = service.ServiceGroupApp.WarServiceGroup.RankLevelService


// CreateRankLevel 创建RankLevel
// @Tags RankLevel
// @Summary 创建RankLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.RankLevel true "创建RankLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rankLevel/createRankLevel [post]
func (rankLevelApi *RankLevelApi) CreateRankLevel(c *gin.Context) {
	var rankLevel war.RankLevel
	err := c.ShouldBindJSON(&rankLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Icon":{utils.NotEmpty()},
        "Experience":{utils.NotEmpty()},
    }
	if err := utils.Verify(rankLevel, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := rankLevelService.CreateRankLevel(&rankLevel); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRankLevel 删除RankLevel
// @Tags RankLevel
// @Summary 删除RankLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.RankLevel true "删除RankLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rankLevel/deleteRankLevel [delete]
func (rankLevelApi *RankLevelApi) DeleteRankLevel(c *gin.Context) {
	var rankLevel war.RankLevel
	err := c.ShouldBindJSON(&rankLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := rankLevelService.DeleteRankLevel(rankLevel); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRankLevelByIds 批量删除RankLevel
// @Tags RankLevel
// @Summary 批量删除RankLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RankLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /rankLevel/deleteRankLevelByIds [delete]
func (rankLevelApi *RankLevelApi) DeleteRankLevelByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := rankLevelService.DeleteRankLevelByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRankLevel 更新RankLevel
// @Tags RankLevel
// @Summary 更新RankLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.RankLevel true "更新RankLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rankLevel/updateRankLevel [put]
func (rankLevelApi *RankLevelApi) UpdateRankLevel(c *gin.Context) {
	var rankLevel war.RankLevel
	err := c.ShouldBindJSON(&rankLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Icon":{utils.NotEmpty()},
          "Experience":{utils.NotEmpty()},
      }
    if err := utils.Verify(rankLevel, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := rankLevelService.UpdateRankLevel(rankLevel); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRankLevel 用id查询RankLevel
// @Tags RankLevel
// @Summary 用id查询RankLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.RankLevel true "用id查询RankLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rankLevel/findRankLevel [get]
func (rankLevelApi *RankLevelApi) FindRankLevel(c *gin.Context) {
	var rankLevel war.RankLevel
	err := c.ShouldBindQuery(&rankLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerankLevel, err := rankLevelService.GetRankLevel(rankLevel.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerankLevel": rerankLevel}, c)
	}
}

// GetRankLevelList 分页获取RankLevel列表
// @Tags RankLevel
// @Summary 分页获取RankLevel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.RankLevelSearch true "分页获取RankLevel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rankLevel/getRankLevelList [get]
func (rankLevelApi *RankLevelApi) GetRankLevelList(c *gin.Context) {
	var pageInfo warReq.RankLevelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := rankLevelService.GetRankLevelInfoList(pageInfo); err != nil {
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
