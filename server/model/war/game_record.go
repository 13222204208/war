// 自动生成模板GameRecord
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// GameRecord 结构体
type GameRecord struct {
      global.GVA_MODEL
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`
      RoomId  *int `json:"roomId" form:"roomId" gorm:"column:room_id;comment:;"`
      GameResult  *int `json:"gameResult" form:"gameResult" gorm:"column:game_result;comment:;"`
      Faction  *int `json:"faction" form:"faction" gorm:"column:faction;comment:所属阵营红队或蓝队;"`
      Round  *int `json:"round" form:"round" gorm:"column:round;comment:第几回合;"`
}


// TableName GameRecord 表名
func (GameRecord) TableName() string {
  return "war_game_record"
}

