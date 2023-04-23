package war

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AnnouncementRouter struct {
}

// InitAnnouncementRouter 初始化 Announcement 路由信息
func (s *AnnouncementRouter) InitAnnouncementRouter(Router *gin.RouterGroup) {
	announcementRouter := Router.Group("announcement").Use(middleware.OperationRecord())
	announcementRouterWithoutRecord := Router.Group("announcement")
	var announcementApi = v1.ApiGroupApp.WarApiGroup.AnnouncementApi
	{
		announcementRouter.POST("createAnnouncement", announcementApi.CreateAnnouncement)             // 新建Announcement
		announcementRouter.DELETE("deleteAnnouncement", announcementApi.DeleteAnnouncement)           // 删除Announcement
		announcementRouter.DELETE("deleteAnnouncementByIds", announcementApi.DeleteAnnouncementByIds) // 批量删除Announcement
		announcementRouter.PUT("updateAnnouncement", announcementApi.UpdateAnnouncement)              // 更新Announcement
	}
	{
		announcementRouterWithoutRecord.GET("findAnnouncement", announcementApi.FindAnnouncement)       // 根据ID获取Announcement
		announcementRouterWithoutRecord.GET("getAnnouncementList", announcementApi.GetAnnouncementList) // 获取Announcement列表
		//根据公告类型获取公告列表
		announcementRouterWithoutRecord.GET(":type", announcementApi.GetAnnouncementListByType) // 获取Announcement列表
	}
}
