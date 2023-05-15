package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type GameRecordService struct {
}

// CreateGameRecord 创建GameRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (gameRecordService *GameRecordService) CreateGameRecord(gameRecord *war.GameRecord) (err error) {
	err = global.GVA_DB.Create(gameRecord).Error
	return err
}

// DeleteGameRecord 删除GameRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (gameRecordService *GameRecordService)DeleteGameRecord(gameRecord war.GameRecord) (err error) {
	err = global.GVA_DB.Delete(&gameRecord).Error
	return err
}

// DeleteGameRecordByIds 批量删除GameRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (gameRecordService *GameRecordService)DeleteGameRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.GameRecord{},"id in ?",ids.Ids).Error
	return err
}

// UpdateGameRecord 更新GameRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (gameRecordService *GameRecordService)UpdateGameRecord(gameRecord war.GameRecord) (err error) {
	err = global.GVA_DB.Save(&gameRecord).Error
	return err
}

// GetGameRecord 根据id获取GameRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (gameRecordService *GameRecordService)GetGameRecord(id uint) (gameRecord war.GameRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&gameRecord).Error
	return
}

// GetGameRecordInfoList 分页获取GameRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (gameRecordService *GameRecordService)GetGameRecordInfoList(info warReq.GameRecordSearch) (list []war.GameRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&war.GameRecord{})
    var gameRecords []war.GameRecord
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Faction != nil {
        db = db.Where("faction = ?",info.Faction)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&gameRecords).Error
	return  gameRecords, total, err
}
