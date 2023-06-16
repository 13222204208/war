package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
)

type ComplaintSearch struct {
	war.Complaint
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

// 投诉信息
type ComplaintInfo struct {
	RoomId int `json:"roomId" form:"roomId"`
	//被投诉人
	Complainee string `json:"complainee" form:"complainee"`
}
