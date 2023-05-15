import service from '@/utils/request'

// @Tags Room
// @Summary 创建Room
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Room true "创建Room"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /room/createRoom [post]
export const createRoom = (data) => {
  return service({
    url: '/room/createRoom',
    method: 'post',
    data
  })
}

// @Tags Room
// @Summary 删除Room
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Room true "删除Room"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /room/deleteRoom [delete]
export const deleteRoom = (data) => {
  return service({
    url: '/room/deleteRoom',
    method: 'delete',
    data
  })
}

// @Tags Room
// @Summary 删除Room
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Room"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /room/deleteRoom [delete]
export const deleteRoomByIds = (data) => {
  return service({
    url: '/room/deleteRoomByIds',
    method: 'delete',
    data
  })
}

// @Tags Room
// @Summary 更新Room
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Room true "更新Room"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /room/updateRoom [put]
export const updateRoom = (data) => {
  return service({
    url: '/room/updateRoom',
    method: 'put',
    data
  })
}

// @Tags Room
// @Summary 用id查询Room
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Room true "用id查询Room"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /room/findRoom [get]
export const findRoom = (params) => {
  return service({
    url: '/room/findRoom',
    method: 'get',
    params
  })
}

// @Tags Room
// @Summary 分页获取Room列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Room列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /room/getRoomList [get]
export const getRoomList = (params) => {
  return service({
    url: '/room/getRoomList',
    method: 'get',
    params
  })
}

//开始游戏
export const startGame = (data) => {
  return service({
    url: '/room/startGame',
    method: 'post',
    data
  })
}