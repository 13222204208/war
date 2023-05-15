package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type MemberRoomSearch struct{
    war.MemberRoom
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
