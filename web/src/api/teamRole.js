import service from '@/utils/request'

// @Tags TeamRole
// @Summary 创建TeamRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamRole true "创建TeamRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamRole/createTeamRole [post]
export const createTeamRole = (data) => {
  return service({
    url: '/teamRole/createTeamRole',
    method: 'post',
    data
  })
}

// @Tags TeamRole
// @Summary 删除TeamRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamRole true "删除TeamRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamRole/deleteTeamRole [delete]
export const deleteTeamRole = (data) => {
  return service({
    url: '/teamRole/deleteTeamRole',
    method: 'delete',
    data
  })
}

// @Tags TeamRole
// @Summary 删除TeamRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamRole/deleteTeamRole [delete]
export const deleteTeamRoleByIds = (data) => {
  return service({
    url: '/teamRole/deleteTeamRoleByIds',
    method: 'delete',
    data
  })
}

// @Tags TeamRole
// @Summary 更新TeamRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamRole true "更新TeamRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamRole/updateTeamRole [put]
export const updateTeamRole = (data) => {
  return service({
    url: '/teamRole/updateTeamRole',
    method: 'put',
    data
  })
}

// @Tags TeamRole
// @Summary 用id查询TeamRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TeamRole true "用id查询TeamRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamRole/findTeamRole [get]
export const findTeamRole = (params) => {
  return service({
    url: '/teamRole/findTeamRole',
    method: 'get',
    params
  })
}

// @Tags TeamRole
// @Summary 分页获取TeamRole列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TeamRole列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamRole/getTeamRoleList [get]
export const getTeamRoleList = (params) => {
  return service({
    url: '/teamRole/getTeamRoleList',
    method: 'get',
    params
  })
}
