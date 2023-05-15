// 自动生成模板Answer
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Answer 结构体
type Answer struct {
      global.GVA_MODEL
      QuestionId  *int `json:"questionId" form:"questionId" gorm:"column:question_id;comment:题目ID;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:答案内容;size:2000;"`
      IsCorrect  *int `json:"isCorrect" form:"isCorrect" gorm:"column:is_correct;comment:答案是否正确 1 正确，2错误;"`
}


// TableName Answer 表名
func (Answer) TableName() string {
  return "war_answer"
}

