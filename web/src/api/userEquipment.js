import service from '@/utils/request'

// @Tags UserEquipment
// @Summary 创建UserEquipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserEquipment true "创建UserEquipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userEquipment/createUserEquipment [post]
export const createUserEquipment = (data) => {
  return service({
    url: '/userEquipment/createUserEquipment',
    method: 'post',
    data
  })
}

// @Tags UserEquipment
// @Summary 删除UserEquipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserEquipment true "删除UserEquipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userEquipment/deleteUserEquipment [delete]
export const deleteUserEquipment = (data) => {
  return service({
    url: '/userEquipment/deleteUserEquipment',
    method: 'delete',
    data
  })
}

// @Tags UserEquipment
// @Summary 删除UserEquipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserEquipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userEquipment/deleteUserEquipment [delete]
export const deleteUserEquipmentByIds = (data) => {
  return service({
    url: '/userEquipment/deleteUserEquipmentByIds',
    method: 'delete',
    data
  })
}

// @Tags UserEquipment
// @Summary 更新UserEquipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserEquipment true "更新UserEquipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userEquipment/updateUserEquipment [put]
export const updateUserEquipment = (data) => {
  return service({
    url: '/userEquipment/updateUserEquipment',
    method: 'put',
    data
  })
}

// @Tags UserEquipment
// @Summary 用id查询UserEquipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserEquipment true "用id查询UserEquipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userEquipment/findUserEquipment [get]
export const findUserEquipment = (params) => {
  return service({
    url: '/userEquipment/findUserEquipment',
    method: 'get',
    params
  })
}

// @Tags UserEquipment
// @Summary 分页获取UserEquipment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UserEquipment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userEquipment/getUserEquipmentList [get]
export const getUserEquipmentList = (params) => {
  return service({
    url: '/userEquipment/getUserEquipmentList',
    method: 'get',
    params
  })
}

//获取所有的装备
export const getEquipmentList = (params) => {
  return service({
    url: '/equipment/getEquipmentList',
    method: 'get',
    params
  })
}