import service from '@/utils/request'

// @Tags GameRecord
// @Summary 创建GameRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GameRecord true "创建GameRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gameRecord/createGameRecord [post]
export const createGameRecord = (data) => {
  return service({
    url: '/gameRecord/createGameRecord',
    method: 'post',
    data
  })
}

// @Tags GameRecord
// @Summary 删除GameRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GameRecord true "删除GameRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gameRecord/deleteGameRecord [delete]
export const deleteGameRecord = (data) => {
  return service({
    url: '/gameRecord/deleteGameRecord',
    method: 'delete',
    data
  })
}

// @Tags GameRecord
// @Summary 删除GameRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GameRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gameRecord/deleteGameRecord [delete]
export const deleteGameRecordByIds = (data) => {
  return service({
    url: '/gameRecord/deleteGameRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags GameRecord
// @Summary 更新GameRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GameRecord true "更新GameRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gameRecord/updateGameRecord [put]
export const updateGameRecord = (data) => {
  return service({
    url: '/gameRecord/updateGameRecord',
    method: 'put',
    data
  })
}

// @Tags GameRecord
// @Summary 用id查询GameRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GameRecord true "用id查询GameRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gameRecord/findGameRecord [get]
export const findGameRecord = (params) => {
  return service({
    url: '/gameRecord/findGameRecord',
    method: 'get',
    params
  })
}

// @Tags GameRecord
// @Summary 分页获取GameRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GameRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gameRecord/getGameRecordList [get]
export const getGameRecordList = (params) => {
  return service({
    url: '/gameRecord/getGameRecordList',
    method: 'get',
    params
  })
}
