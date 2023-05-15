package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type QuestionSearch struct{
    war.Question
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
