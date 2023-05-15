// 自动生成模板Question
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Question 结构体
type Question struct {
      global.GVA_MODEL
      Content  string `json:"content" form:"content" gorm:"column:content;comment:题目内容;size:1000;"`
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:题目类型;"`
      //题目的答案
      Answer []*Answer `json:"answer" form:"answer" gorm:"foreignKey:QuestionId;references:ID;comment:题目的答案;"`
}


// TableName Question 表名
func (Question) TableName() string {
  return "war_question"
}

