import service from '@/utils/request'

// @Tags Slot
// @Summary 创建Slot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Slot true "创建Slot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slot/createSlot [post]
export const createSlot = (data) => {
  return service({
    url: '/slot/createSlot',
    method: 'post',
    data
  })
}

// @Tags Slot
// @Summary 删除Slot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Slot true "删除Slot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slot/deleteSlot [delete]
export const deleteSlot = (data) => {
  return service({
    url: '/slot/deleteSlot',
    method: 'delete',
    data
  })
}

// @Tags Slot
// @Summary 删除Slot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Slot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slot/deleteSlot [delete]
export const deleteSlotByIds = (data) => {
  return service({
    url: '/slot/deleteSlotByIds',
    method: 'delete',
    data
  })
}

// @Tags Slot
// @Summary 更新Slot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Slot true "更新Slot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slot/updateSlot [put]
export const updateSlot = (data) => {
  return service({
    url: '/slot/updateSlot',
    method: 'put',
    data
  })
}

// @Tags Slot
// @Summary 用id查询Slot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Slot true "用id查询Slot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slot/findSlot [get]
export const findSlot = (params) => {
  return service({
    url: '/slot/findSlot',
    method: 'get',
    params
  })
}

// @Tags Slot
// @Summary 分页获取Slot列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Slot列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slot/getSlotList [get]
export const getSlotList = (params) => {
  return service({
    url: '/slot/getSlotList',
    method: 'get',
    params
  })
}
