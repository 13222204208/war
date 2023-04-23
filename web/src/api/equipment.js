import service from '@/utils/request'

// @Tags Equipment
// @Summary 创建Equipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Equipment true "创建Equipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /equipment/createEquipment [post]
export const createEquipment = (data) => {
  return service({
    url: '/equipment/createEquipment',
    method: 'post',
    data
  })
}

// @Tags Equipment
// @Summary 删除Equipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Equipment true "删除Equipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /equipment/deleteEquipment [delete]
export const deleteEquipment = (data) => {
  return service({
    url: '/equipment/deleteEquipment',
    method: 'delete',
    data
  })
}

// @Tags Equipment
// @Summary 删除Equipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Equipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /equipment/deleteEquipment [delete]
export const deleteEquipmentByIds = (data) => {
  return service({
    url: '/equipment/deleteEquipmentByIds',
    method: 'delete',
    data
  })
}

// @Tags Equipment
// @Summary 更新Equipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Equipment true "更新Equipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /equipment/updateEquipment [put]
export const updateEquipment = (data) => {
  return service({
    url: '/equipment/updateEquipment',
    method: 'put',
    data
  })
}

// @Tags Equipment
// @Summary 用id查询Equipment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Equipment true "用id查询Equipment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /equipment/findEquipment [get]
export const findEquipment = (params) => {
  return service({
    url: '/equipment/findEquipment',
    method: 'get',
    params
  })
}

// @Tags Equipment
// @Summary 分页获取Equipment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Equipment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /equipment/getEquipmentList [get]
export const getEquipmentList = (params) => {
  return service({
    url: '/equipment/getEquipmentList',
    method: 'get',
    params
  })
}
