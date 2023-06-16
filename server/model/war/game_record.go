// 自动生成模板GameRecord
package war

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GameRecord 结构体
type GameRecord struct {
	global.GVA_MODEL
	UserId     uint `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`
	RoomId     uint `json:"roomId" form:"roomId" gorm:"column:room_id;comment:;"`
	GameResult int  `json:"gameResult" form:"gameResult" gorm:"column:game_result;comment:;"`
	Faction    int  `json:"faction" form:"faction" gorm:"column:faction;comment:所属阵营红队或蓝队;"`
	Score      int  `json:"score" form:"score" gorm:"column:score;comment:得分;"`
	Experience int  `json:"experience" form:"experience" gorm:"column:experience;comment:经验;"`
	Round      *int `json:"round" form:"round" gorm:"column:round;comment:第几回合;"`
	//比赛类型
	GameType int `json:"gameType" form:"gameType" gorm:"column:game_type;comment:比赛类型 1普通比赛，2战队赛;default:1;"`
	//会员信息
	User Member `json:"user" form:"user" gorm:"foreignKey:UserId;references:ID;comment:会员信息;"`
	Room Room   `json:"room" form:"room" gorm:"foreignKey:RoomId;references:ID;comment:房间信息;"`
}

// TableName GameRecord 表名
func (GameRecord) TableName() string {
	return "war_game_record"
}
