import service from '@/utils/request'

// @Tags Announcement
// @Summary 创建Announcement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Announcement true "创建Announcement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /announcement/createAnnouncement [post]
export const createAnnouncement = (data) => {
  return service({
    url: '/announcement/createAnnouncement',
    method: 'post',
    data
  })
}

// @Tags Announcement
// @Summary 删除Announcement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Announcement true "删除Announcement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /announcement/deleteAnnouncement [delete]
export const deleteAnnouncement = (data) => {
  return service({
    url: '/announcement/deleteAnnouncement',
    method: 'delete',
    data
  })
}

// @Tags Announcement
// @Summary 删除Announcement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Announcement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /announcement/deleteAnnouncement [delete]
export const deleteAnnouncementByIds = (data) => {
  return service({
    url: '/announcement/deleteAnnouncementByIds',
    method: 'delete',
    data
  })
}

// @Tags Announcement
// @Summary 更新Announcement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Announcement true "更新Announcement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /announcement/updateAnnouncement [put]
export const updateAnnouncement = (data) => {
  return service({
    url: '/announcement/updateAnnouncement',
    method: 'put',
    data
  })
}

// @Tags Announcement
// @Summary 用id查询Announcement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Announcement true "用id查询Announcement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /announcement/findAnnouncement [get]
export const findAnnouncement = (params) => {
  return service({
    url: '/announcement/findAnnouncement',
    method: 'get',
    params
  })
}

// @Tags Announcement
// @Summary 分页获取Announcement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Announcement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /announcement/getAnnouncementList [get]
export const getAnnouncementList = (params) => {
  return service({
    url: '/announcement/getAnnouncementList',
    method: 'get',
    params
  })
}
