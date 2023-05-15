package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/war"

// 返回红蓝双方信息
type MemberRoomResponse struct {
	//红方
	Red []war.MemberRoom `json:"red"`
	//蓝方
	Blue []war.MemberRoom `json:"blue"`
}
