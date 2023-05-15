// 自动生成模板MemberRoom
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MemberRoom 结构体
type MemberRoom struct {
	global.GVA_MODEL
	RoomId  uint `json:"roomId" form:"roomId" gorm:"column:room_id;comment:房间ID;"`
	UserId  uint `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`
	Faction int  `json:"faction" form:"faction" gorm:"column:faction;comment:所属的阵营;"`
	//签到状态
	SignStatus int `json:"signStatus" form:"signStatus" gorm:"column:sign_status;comment:签到状态;"`
	//房间信息
	Room Room `json:"room" form:"room" gorm:"foreignKey:RoomId;references:ID;comment:房间信息;"`
	//会员信息
	User Member `json:"user" form:"user" gorm:"foreignKey:UserId;references:ID;comment:会员信息;"`
}

// TableName MemberRoom 表名
func (MemberRoom) TableName() string {
	return "war_member_room"
}
