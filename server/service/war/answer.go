package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type AnswerService struct {
}

// CreateAnswer 创建Answer记录
// Author [piexlmax](https://github.com/piexlmax)
func (answerService *AnswerService) CreateAnswer(answer *war.Answer) (err error) {
	err = global.GVA_DB.Create(answer).Error
	return err
}

// DeleteAnswer 删除Answer记录
// Author [piexlmax](https://github.com/piexlmax)
func (answerService *AnswerService)DeleteAnswer(answer war.Answer) (err error) {
	err = global.GVA_DB.Delete(&answer).Error
	return err
}

// DeleteAnswerByIds 批量删除Answer记录
// Author [piexlmax](https://github.com/piexlmax)
func (answerService *AnswerService)DeleteAnswerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.Answer{},"id in ?",ids.Ids).Error
	return err
}

// UpdateAnswer 更新Answer记录
// Author [piexlmax](https://github.com/piexlmax)
func (answerService *AnswerService)UpdateAnswer(answer war.Answer) (err error) {
	err = global.GVA_DB.Save(&answer).Error
	return err
}

// GetAnswer 根据id获取Answer记录
// Author [piexlmax](https://github.com/piexlmax)
func (answerService *AnswerService)GetAnswer(id uint) (answer war.Answer, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&answer).Error
	return
}

// GetAnswerInfoList 分页获取Answer记录
// Author [piexlmax](https://github.com/piexlmax)
func (answerService *AnswerService)GetAnswerInfoList(info warReq.AnswerSearch) (list []war.Answer, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&war.Answer{})
    var answers []war.Answer
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&answers).Error
	return  answers, total, err
}
