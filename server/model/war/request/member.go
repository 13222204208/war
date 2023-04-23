package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
)

type MemberSearch struct {
	war.Member
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

// 微信小程序换取openid
type WechatCode struct {
	Code string `json:"code" form:"code"`
}

// 会员增加或修改场次
type MemberMatch struct {
	Match     uint `json:"match" form:"match"`
	UserId    uint `json:"userId" form:"userId"`
	MatchType uint `json:"matchType" form:"matchType"`
}
