import service from '@/utils/request'

// @Tags Member
// @Summary 创建Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Member true "创建Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /member/createMember [post]
export const createMember = (data) => {
  return service({
    url: '/member/createMember',
    method: 'post',
    data
  })
}

// @Tags Member
// @Summary 删除Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Member true "删除Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /member/deleteMember [delete]
export const deleteMember = (data) => {
  return service({
    url: '/member/deleteMember',
    method: 'delete',
    data
  })
}

// @Tags Member
// @Summary 删除Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /member/deleteMember [delete]
export const deleteMemberByIds = (data) => {
  return service({
    url: '/member/deleteMemberByIds',
    method: 'delete',
    data
  })
}

// @Tags Member
// @Summary 更新Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Member true "更新Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /member/updateMember [put]
export const updateMember = (data) => {
  return service({
    url: '/member/updateMember',
    method: 'put',
    data
  })
}

// @Tags Member
// @Summary 用id查询Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Member true "用id查询Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /member/findMember [get]
export const findMember = (params) => {
  return service({
    url: '/member/findMember',
    method: 'get',
    params
  })
}

// @Tags Member
// @Summary 分页获取Member列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Member列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /member/getMemberList [get]
export const getMemberList = (params) => {
  return service({
    url: '/member/getMemberList',
    method: 'get',
    params
  })
}
