package war

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type RankLevelService struct {
}

// CreateRankLevel 创建RankLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (rankLevelService *RankLevelService) CreateRankLevel(rankLevel *war.RankLevel) (err error) {
	err = global.GVA_DB.Create(rankLevel).Error
	return err
}

// DeleteRankLevel 删除RankLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (rankLevelService *RankLevelService) DeleteRankLevel(rankLevel war.RankLevel) (err error) {
	err = global.GVA_DB.Delete(&rankLevel).Error
	return err
}

// DeleteRankLevelByIds 批量删除RankLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (rankLevelService *RankLevelService) DeleteRankLevelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.RankLevel{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRankLevel 更新RankLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (rankLevelService *RankLevelService) UpdateRankLevel(rankLevel war.RankLevel) (err error) {
	err = global.GVA_DB.Save(&rankLevel).Error
	return err
}

// GetRankLevel 根据id获取RankLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (rankLevelService *RankLevelService) GetRankLevel(id uint) (rankLevel war.RankLevel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&rankLevel).Error
	return
}

// GetRankLevelInfoList 分页获取RankLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (rankLevelService *RankLevelService) GetRankLevelInfoList(info warReq.RankLevelSearch) (list []war.RankLevel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.RankLevel{})
	var rankLevels []war.RankLevel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&rankLevels).Error
	return rankLevels, total, err
}

// 判断用户的军衔在哪一个等级
func UpdateRank(userId uint, exp int) error {
	fmt.Println("用户的id", userId, "用户经验", exp)
	//查询下一个军衔
	var rankLevel war.RankLevel
	err := global.GVA_DB.Where("experience < ?", exp).Order("experience DESC").First(&rankLevel).Error
	if err != nil {
		return err
	}
	//判断当前经验是否大于下一个军衔的经验
	nowExp := rankLevel.Experience
	fmt.Println("军衔的经验", nowExp, "军衔的ID", rankLevel.ID)

	err = global.GVA_DB.Model(&war.Member{}).Where("id = ?", userId).Update("rank_level_id", rankLevel.ID).Error
	if err != nil {
		return err
	}
	return nil
}
