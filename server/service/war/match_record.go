package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type MatchRecordService struct {
}

// CreateMatchRecord 创建MatchRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchRecordService *MatchRecordService) CreateMatchRecord(matchRecord *war.MatchRecord) (err error) {
	err = global.GVA_DB.Create(matchRecord).Error
	return err
}

// DeleteMatchRecord 删除MatchRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchRecordService *MatchRecordService)DeleteMatchRecord(matchRecord war.MatchRecord) (err error) {
	err = global.GVA_DB.Delete(&matchRecord).Error
	return err
}

// DeleteMatchRecordByIds 批量删除MatchRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchRecordService *MatchRecordService)DeleteMatchRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.MatchRecord{},"id in ?",ids.Ids).Error
	return err
}

// UpdateMatchRecord 更新MatchRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchRecordService *MatchRecordService)UpdateMatchRecord(matchRecord war.MatchRecord) (err error) {
	err = global.GVA_DB.Save(&matchRecord).Error
	return err
}

// GetMatchRecord 根据id获取MatchRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchRecordService *MatchRecordService)GetMatchRecord(id uint) (matchRecord war.MatchRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&matchRecord).Error
	return
}

// GetMatchRecordInfoList 分页获取MatchRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchRecordService *MatchRecordService)GetMatchRecordInfoList(info warReq.MatchRecordSearch) (list []war.MatchRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&war.MatchRecord{})
    var matchRecords []war.MatchRecord
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&matchRecords).Error
	return  matchRecords, total, err
}
