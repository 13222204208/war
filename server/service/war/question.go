package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
)

type QuestionService struct {
}

// CreateQuestion 创建Question记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionService *QuestionService) CreateQuestion(question *war.Question) (err error) {
	err = global.GVA_DB.Create(question).Error
	return err
}

// DeleteQuestion 删除Question记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionService *QuestionService)DeleteQuestion(question war.Question) (err error) {
	err = global.GVA_DB.Delete(&question).Error
	return err
}

// DeleteQuestionByIds 批量删除Question记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionService *QuestionService)DeleteQuestionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.Question{},"id in ?",ids.Ids).Error
	return err
}

// UpdateQuestion 更新Question记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionService *QuestionService)UpdateQuestion(question war.Question) (err error) {
	err = global.GVA_DB.Save(&question).Error
	return err
}

// GetQuestion 根据id获取Question记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionService *QuestionService)GetQuestion(id uint) (question war.Question, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&question).Error
	return
}

// GetQuestionInfoList 分页获取Question记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionService *QuestionService)GetQuestionInfoList(info warReq.QuestionSearch) (list []war.Question, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&war.Question{})
    var questions []war.Question
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&questions).Error
	return  questions, total, err
}

//获取所有的问题和答案
func (questionService *QuestionService)GetQuestionAndAnswer() (list []war.Question, err error) {
	err = global.GVA_DB.Preload("Answer").Find(&list).Error
	return
}