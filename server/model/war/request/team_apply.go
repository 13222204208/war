package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
)

type TeamApplySearch struct {
	war.TeamApply
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

// 审批用户申请
type TeamApplyApproval struct {
	Status int `json:"status" form:"status"` //1同意，2拒绝
}
