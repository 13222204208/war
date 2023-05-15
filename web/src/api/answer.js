import service from '@/utils/request'

// @Tags Answer
// @Summary 创建Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer true "创建Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /answer/createAnswer [post]
export const createAnswer = (data) => {
  return service({
    url: '/answer/createAnswer',
    method: 'post',
    data
  })
}

// @Tags Answer
// @Summary 删除Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer true "删除Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /answer/deleteAnswer [delete]
export const deleteAnswer = (data) => {
  return service({
    url: '/answer/deleteAnswer',
    method: 'delete',
    data
  })
}

// @Tags Answer
// @Summary 删除Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /answer/deleteAnswer [delete]
export const deleteAnswerByIds = (data) => {
  return service({
    url: '/answer/deleteAnswerByIds',
    method: 'delete',
    data
  })
}

// @Tags Answer
// @Summary 更新Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer true "更新Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /answer/updateAnswer [put]
export const updateAnswer = (data) => {
  return service({
    url: '/answer/updateAnswer',
    method: 'put',
    data
  })
}

// @Tags Answer
// @Summary 用id查询Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Answer true "用id查询Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /answer/findAnswer [get]
export const findAnswer = (params) => {
  return service({
    url: '/answer/findAnswer',
    method: 'get',
    params
  })
}

// @Tags Answer
// @Summary 分页获取Answer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Answer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /answer/getAnswerList [get]
export const getAnswerList = (params) => {
  return service({
    url: '/answer/getAnswerList',
    method: 'get',
    params
  })
}

//查询出所有题目
export const getQuestionList = (params) => {
  return service({
    url: '/question/getQuestionList',
    method: 'get',
    params
  })
}