package war

import (
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type ComplaintService struct {
}

// CreateComplaint 创建Complaint记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintService *ComplaintService) CreateComplaint(complaint *war.Complaint) (err error) {
	err = global.GVA_DB.Create(complaint).Error
	return err
}

// DeleteComplaint 删除Complaint记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintService *ComplaintService) DeleteComplaint(complaint war.Complaint) (err error) {
	err = global.GVA_DB.Delete(&complaint).Error
	return err
}

// DeleteComplaintByIds 批量删除Complaint记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintService *ComplaintService) DeleteComplaintByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.Complaint{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateComplaint 更新Complaint记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintService *ComplaintService) UpdateComplaint(complaint war.Complaint) (err error) {
	err = global.GVA_DB.Save(&complaint).Error
	return err
}

// GetComplaint 根据id获取Complaint记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintService *ComplaintService) GetComplaint(id uint) (complaint war.Complaint, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&complaint).Error
	return
}

// GetComplaintInfoList 分页获取Complaint记录
// Author [piexlmax](https://github.com/piexlmax)
func (complaintService *ComplaintService) GetComplaintInfoList(info warReq.ComplaintSearch) (list []war.Complaint, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.Complaint{})
	var complaints []war.Complaint
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&complaints).Error
	return complaints, total, err
}

// 会员投诉
func (complaintService *ComplaintService) MemberComplaint(complaint *warReq.ComplaintInfo, userId uint) (err error) {
	complainee := complaint.Complainee
	complainees := strings.Split(complainee, ",")
	for _, v := range complainees {
		num, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		c := war.Complaint{
			Complainant: userId,
			Complainee:  num,
			RoomId:      complaint.RoomId,
		}
		err = global.GVA_DB.Create(&c).Error
		if err != nil {
			return err
		}
	}
	return nil
}
