import service from '@/utils/request'

// @Tags MatchRecord
// @Summary 创建MatchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MatchRecord true "创建MatchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /matchRecord/createMatchRecord [post]
export const createMatchRecord = (data) => {
  return service({
    url: '/matchRecord/createMatchRecord',
    method: 'post',
    data
  })
}

// @Tags MatchRecord
// @Summary 删除MatchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MatchRecord true "删除MatchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /matchRecord/deleteMatchRecord [delete]
export const deleteMatchRecord = (data) => {
  return service({
    url: '/matchRecord/deleteMatchRecord',
    method: 'delete',
    data
  })
}

// @Tags MatchRecord
// @Summary 删除MatchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MatchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /matchRecord/deleteMatchRecord [delete]
export const deleteMatchRecordByIds = (data) => {
  return service({
    url: '/matchRecord/deleteMatchRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags MatchRecord
// @Summary 更新MatchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MatchRecord true "更新MatchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /matchRecord/updateMatchRecord [put]
export const updateMatchRecord = (data) => {
  return service({
    url: '/matchRecord/updateMatchRecord',
    method: 'put',
    data
  })
}

// @Tags MatchRecord
// @Summary 用id查询MatchRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MatchRecord true "用id查询MatchRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /matchRecord/findMatchRecord [get]
export const findMatchRecord = (params) => {
  return service({
    url: '/matchRecord/findMatchRecord',
    method: 'get',
    params
  })
}

// @Tags MatchRecord
// @Summary 分页获取MatchRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MatchRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /matchRecord/getMatchRecordList [get]
export const getMatchRecordList = (params) => {
  return service({
    url: '/matchRecord/getMatchRecordList',
    method: 'get',
    params
  })
}
