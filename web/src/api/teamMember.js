import service from '@/utils/request'

// @Tags TeamMember
// @Summary 创建TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamMember true "创建TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamMember/createTeamMember [post]
export const createTeamMember = (data) => {
  return service({
    url: '/teamMember/createTeamMember',
    method: 'post',
    data
  })
}

// @Tags TeamMember
// @Summary 删除TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamMember true "删除TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamMember/deleteTeamMember [delete]
export const deleteTeamMember = (data) => {
  return service({
    url: '/teamMember/deleteTeamMember',
    method: 'delete',
    data
  })
}

// @Tags TeamMember
// @Summary 删除TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamMember/deleteTeamMember [delete]
export const deleteTeamMemberByIds = (data) => {
  return service({
    url: '/teamMember/deleteTeamMemberByIds',
    method: 'delete',
    data
  })
}

// @Tags TeamMember
// @Summary 更新TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamMember true "更新TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamMember/updateTeamMember [put]
export const updateTeamMember = (data) => {
  return service({
    url: '/teamMember/updateTeamMember',
    method: 'put',
    data
  })
}

// @Tags TeamMember
// @Summary 用id查询TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TeamMember true "用id查询TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamMember/findTeamMember [get]
export const findTeamMember = (params) => {
  return service({
    url: '/teamMember/findTeamMember',
    method: 'get',
    params
  })
}

// @Tags TeamMember
// @Summary 分页获取TeamMember列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TeamMember列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamMember/getTeamMemberList [get]
export const getTeamMemberList = (params) => {
  return service({
    url: '/teamMember/getTeamMemberList',
    method: 'get',
    params
  })
}
