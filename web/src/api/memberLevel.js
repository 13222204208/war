import service from '@/utils/request'

// @Tags MemberLevel
// @Summary 创建MemberLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemberLevel true "创建MemberLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memberLevel/createMemberLevel [post]
export const createMemberLevel = (data) => {
  return service({
    url: '/memberLevel/createMemberLevel',
    method: 'post',
    data
  })
}

// @Tags MemberLevel
// @Summary 删除MemberLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemberLevel true "删除MemberLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memberLevel/deleteMemberLevel [delete]
export const deleteMemberLevel = (data) => {
  return service({
    url: '/memberLevel/deleteMemberLevel',
    method: 'delete',
    data
  })
}

// @Tags MemberLevel
// @Summary 删除MemberLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MemberLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /memberLevel/deleteMemberLevel [delete]
export const deleteMemberLevelByIds = (data) => {
  return service({
    url: '/memberLevel/deleteMemberLevelByIds',
    method: 'delete',
    data
  })
}

// @Tags MemberLevel
// @Summary 更新MemberLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MemberLevel true "更新MemberLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /memberLevel/updateMemberLevel [put]
export const updateMemberLevel = (data) => {
  return service({
    url: '/memberLevel/updateMemberLevel',
    method: 'put',
    data
  })
}

// @Tags MemberLevel
// @Summary 用id查询MemberLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MemberLevel true "用id查询MemberLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /memberLevel/findMemberLevel [get]
export const findMemberLevel = (params) => {
  return service({
    url: '/memberLevel/findMemberLevel',
    method: 'get',
    params
  })
}

// @Tags MemberLevel
// @Summary 分页获取MemberLevel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MemberLevel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /memberLevel/getMemberLevelList [get]
export const getMemberLevelList = (params) => {
  return service({
    url: '/memberLevel/getMemberLevelList',
    method: 'get',
    params
  })
}
