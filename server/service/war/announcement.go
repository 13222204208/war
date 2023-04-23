package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type AnnouncementService struct {
}

// CreateAnnouncement 创建Announcement记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementService *AnnouncementService) CreateAnnouncement(announcement *war.Announcement) (err error) {
	err = global.GVA_DB.Create(announcement).Error
	return err
}

// DeleteAnnouncement 删除Announcement记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementService *AnnouncementService) DeleteAnnouncement(announcement war.Announcement) (err error) {
	err = global.GVA_DB.Delete(&announcement).Error
	return err
}

// DeleteAnnouncementByIds 批量删除Announcement记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementService *AnnouncementService) DeleteAnnouncementByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.Announcement{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAnnouncement 更新Announcement记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementService *AnnouncementService) UpdateAnnouncement(announcement war.Announcement) (err error) {
	err = global.GVA_DB.Save(&announcement).Error
	return err
}

// GetAnnouncement 根据id获取Announcement记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementService *AnnouncementService) GetAnnouncement(id uint) (announcement war.Announcement, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&announcement).Error
	return
}

// GetAnnouncementInfoList 分页获取Announcement记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementService *AnnouncementService) GetAnnouncementInfoList(info warReq.AnnouncementSearch) (list []war.Announcement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.Announcement{})
	var announcements []war.Announcement
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&announcements).Error
	return announcements, total, err
}

// 根据公告类型获取公告
func (announcementService *AnnouncementService) GetAnnouncementByType(announcementType int) (announcement war.Announcement, err error) {
	err = global.GVA_DB.Where("type = ? AND status = ?", announcementType, 1).First(&announcement).Error
	return
}
