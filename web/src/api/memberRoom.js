import service from '@/utils/request'

// @Tags MemberRoom
// @Summary 创建MemberRoom
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemberRoom true "创建MemberRoom"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memberRoom/createMemberRoom [post]
export const createMemberRoom = (data) => {
  return service({
    url: '/memberRoom/createMemberRoom',
    method: 'post',
    data
  })
}

// @Tags MemberRoom
// @Summary 删除MemberRoom
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemberRoom true "删除MemberRoom"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memberRoom/deleteMemberRoom [delete]
export const deleteMemberRoom = (data) => {
  return service({
    url: '/memberRoom/deleteMemberRoom',
    method: 'delete',
    data
  })
}

// @Tags MemberRoom
// @Summary 删除MemberRoom
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemberRoom"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memberRoom/deleteMemberRoom [delete]
export const deleteMemberRoomByIds = (data) => {
  return service({
    url: '/memberRoom/deleteMemberRoomByIds',
    method: 'delete',
    data
  })
}

// @Tags MemberRoom
// @Summary 更新MemberRoom
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemberRoom true "更新MemberRoom"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memberRoom/updateMemberRoom [put]
export const updateMemberRoom = (data) => {
  return service({
    url: '/memberRoom/updateMemberRoom',
    method: 'put',
    data
  })
}

// @Tags MemberRoom
// @Summary 用id查询MemberRoom
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MemberRoom true "用id查询MemberRoom"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memberRoom/findMemberRoom [get]
export const findMemberRoom = (params) => {
  return service({
    url: '/memberRoom/findMemberRoom',
    method: 'get',
    params
  })
}

// @Tags MemberRoom
// @Summary 分页获取MemberRoom列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MemberRoom列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memberRoom/getMemberRoomList [get]
export const getMemberRoomList = (params) => {
  return service({
    url: '/memberRoom/getMemberRoomList',
    method: 'get',
    params
  })
}
