import service from '@/utils/request'

// @Tags Question
// @Summary 创建Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question true "创建Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question/createQuestion [post]
export const createQuestion = (data) => {
  return service({
    url: '/question/createQuestion',
    method: 'post',
    data
  })
}

// @Tags Question
// @Summary 删除Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question true "删除Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /question/deleteQuestion [delete]
export const deleteQuestion = (data) => {
  return service({
    url: '/question/deleteQuestion',
    method: 'delete',
    data
  })
}

// @Tags Question
// @Summary 删除Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /question/deleteQuestion [delete]
export const deleteQuestionByIds = (data) => {
  return service({
    url: '/question/deleteQuestionByIds',
    method: 'delete',
    data
  })
}

// @Tags Question
// @Summary 更新Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question true "更新Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /question/updateQuestion [put]
export const updateQuestion = (data) => {
  return service({
    url: '/question/updateQuestion',
    method: 'put',
    data
  })
}

// @Tags Question
// @Summary 用id查询Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Question true "用id查询Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /question/findQuestion [get]
export const findQuestion = (params) => {
  return service({
    url: '/question/findQuestion',
    method: 'get',
    params
  })
}

// @Tags Question
// @Summary 分页获取Question列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Question列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question/getQuestionList [get]
export const getQuestionList = (params) => {
  return service({
    url: '/question/getQuestionList',
    method: 'get',
    params
  })
}
