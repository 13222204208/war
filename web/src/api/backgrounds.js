import service from '@/utils/request'

// @Tags Backgrounds
// @Summary 创建Backgrounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Backgrounds true "创建Backgrounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /backgrounds/createBackgrounds [post]
export const createBackgrounds = (data) => {
  return service({
    url: '/backgrounds/createBackgrounds',
    method: 'post',
    data
  })
}

// @Tags Backgrounds
// @Summary 删除Backgrounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Backgrounds true "删除Backgrounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /backgrounds/deleteBackgrounds [delete]
export const deleteBackgrounds = (data) => {
  return service({
    url: '/backgrounds/deleteBackgrounds',
    method: 'delete',
    data
  })
}

// @Tags Backgrounds
// @Summary 删除Backgrounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Backgrounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /backgrounds/deleteBackgrounds [delete]
export const deleteBackgroundsByIds = (data) => {
  return service({
    url: '/backgrounds/deleteBackgroundsByIds',
    method: 'delete',
    data
  })
}

// @Tags Backgrounds
// @Summary 更新Backgrounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Backgrounds true "更新Backgrounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /backgrounds/updateBackgrounds [put]
export const updateBackgrounds = (data) => {
  return service({
    url: '/backgrounds/updateBackgrounds',
    method: 'put',
    data
  })
}

// @Tags Backgrounds
// @Summary 用id查询Backgrounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Backgrounds true "用id查询Backgrounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /backgrounds/findBackgrounds [get]
export const findBackgrounds = (params) => {
  return service({
    url: '/backgrounds/findBackgrounds',
    method: 'get',
    params
  })
}

// @Tags Backgrounds
// @Summary 分页获取Backgrounds列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Backgrounds列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /backgrounds/getBackgroundsList [get]
export const getBackgroundsList = (params) => {
  return service({
    url: '/backgrounds/getBackgroundsList',
    method: 'get',
    params
  })
}
