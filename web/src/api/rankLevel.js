import service from '@/utils/request'

// @Tags RankLevel
// @Summary 创建RankLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RankLevel true "创建RankLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rankLevel/createRankLevel [post]
export const createRankLevel = (data) => {
  return service({
    url: '/rankLevel/createRankLevel',
    method: 'post',
    data
  })
}

// @Tags RankLevel
// @Summary 删除RankLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RankLevel true "删除RankLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rankLevel/deleteRankLevel [delete]
export const deleteRankLevel = (data) => {
  return service({
    url: '/rankLevel/deleteRankLevel',
    method: 'delete',
    data
  })
}

// @Tags RankLevel
// @Summary 删除RankLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RankLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rankLevel/deleteRankLevel [delete]
export const deleteRankLevelByIds = (data) => {
  return service({
    url: '/rankLevel/deleteRankLevelByIds',
    method: 'delete',
    data
  })
}

// @Tags RankLevel
// @Summary 更新RankLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RankLevel true "更新RankLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rankLevel/updateRankLevel [put]
export const updateRankLevel = (data) => {
  return service({
    url: '/rankLevel/updateRankLevel',
    method: 'put',
    data
  })
}

// @Tags RankLevel
// @Summary 用id查询RankLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RankLevel true "用id查询RankLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rankLevel/findRankLevel [get]
export const findRankLevel = (params) => {
  return service({
    url: '/rankLevel/findRankLevel',
    method: 'get',
    params
  })
}

// @Tags RankLevel
// @Summary 分页获取RankLevel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取RankLevel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rankLevel/getRankLevelList [get]
export const getRankLevelList = (params) => {
  return service({
    url: '/rankLevel/getRankLevelList',
    method: 'get',
    params
  })
}
