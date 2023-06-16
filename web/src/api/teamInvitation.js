import service from '@/utils/request'

// @Tags TeamInvitation
// @Summary 创建TeamInvitation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamInvitation true "创建TeamInvitation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamInvitation/createTeamInvitation [post]
export const createTeamInvitation = (data) => {
  return service({
    url: '/teamInvitation/createTeamInvitation',
    method: 'post',
    data
  })
}

// @Tags TeamInvitation
// @Summary 删除TeamInvitation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamInvitation true "删除TeamInvitation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamInvitation/deleteTeamInvitation [delete]
export const deleteTeamInvitation = (data) => {
  return service({
    url: '/teamInvitation/deleteTeamInvitation',
    method: 'delete',
    data
  })
}

// @Tags TeamInvitation
// @Summary 删除TeamInvitation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamInvitation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamInvitation/deleteTeamInvitation [delete]
export const deleteTeamInvitationByIds = (data) => {
  return service({
    url: '/teamInvitation/deleteTeamInvitationByIds',
    method: 'delete',
    data
  })
}

// @Tags TeamInvitation
// @Summary 更新TeamInvitation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamInvitation true "更新TeamInvitation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamInvitation/updateTeamInvitation [put]
export const updateTeamInvitation = (data) => {
  return service({
    url: '/teamInvitation/updateTeamInvitation',
    method: 'put',
    data
  })
}

// @Tags TeamInvitation
// @Summary 用id查询TeamInvitation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TeamInvitation true "用id查询TeamInvitation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamInvitation/findTeamInvitation [get]
export const findTeamInvitation = (params) => {
  return service({
    url: '/teamInvitation/findTeamInvitation',
    method: 'get',
    params
  })
}

// @Tags TeamInvitation
// @Summary 分页获取TeamInvitation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TeamInvitation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamInvitation/getTeamInvitationList [get]
export const getTeamInvitationList = (params) => {
  return service({
    url: '/teamInvitation/getTeamInvitationList',
    method: 'get',
    params
  })
}
