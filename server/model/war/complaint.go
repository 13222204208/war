// 自动生成模板Complaint
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Complaint 结构体
type Complaint struct {
	global.GVA_MODEL
	RoomId      int    `json:"roomId" form:"roomId" gorm:"column:room_id;comment:;default:0;"`
	Complainant uint   `json:"complainant" form:"complainant" gorm:"column:complainant;comment:投诉人ID;default:0;"`
	Complainee  int    `json:"complainee" form:"complainee" gorm:"column:complainee;comment:被投诉人;default:0;"`
	Reason      string `json:"reason" form:"reason" gorm:"column:reason;comment:投诉原因;"`
	Status      *int   `json:"status" form:"status" gorm:"column:status;comment:投诉的状态，0未处理，1已处理，2拒绝;default:0;"`
}

// TableName Complaint 表名
func (Complaint) TableName() string {
	return "war_complaint"
}
