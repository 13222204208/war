// 自动生成模板Room
package war

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Room 结构体
type Room struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"`
	//第几个房间
	RoomNum    int64 `json:"roomNum" form:"roomNum" gorm:"column:room_num;comment:第几个房间;default:0"`
	MinPlayers *int  `json:"minPlayers" form:"minPlayers" gorm:"column:min_players;comment:游戏开始的最少人数;default:6"`
	MaxPlayers *int  `json:"maxPlayers" form:"maxPlayers" gorm:"column:max_players;comment:游戏房间最多人数;default:24"`
	Countdown  *int  `json:"countdown" form:"countdown" gorm:"column:countdown;comment:倒计时还有多少分钟;default:15"`
	NumPlayers int   `json:"numPlayers" form:"numPlayers" gorm:"column:num_players;comment:房间目前的人数;default:0"`
	//实到人数
	ActualNumPlayers int       `json:"actualNumPlayers" form:"actualNumPlayers" gorm:"column:actual_num_players;comment:实到人数;default:0"`
	EndTime          time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:倒计时结束时间;default:null"`
	GameOverTime     time.Time `json:"gameOverTime" form:"gameOverTime" gorm:"column:game_over_time;comment:游戏结束时间;default:null"`
	//红方分数
	RedScore int `json:"redScore" form:"redScore" gorm:"column:red_score;comment:红方分数;default:0"`
	//蓝方分数
	BlueScore int `json:"blueScore" form:"blueScore" gorm:"column:blue_score;comment:蓝方分数;default:0"`
	Status    int `json:"status" form:"status" gorm:"column:status;comment:房间状态1 准备中， 2 游戏中 3已结束;default:0"`
}

// TableName Room 表名
func (Room) TableName() string {
	return "war_room"
}

// 保存游戏记录信息
type SaveGameRecord struct {
	MemberRoom     []MemberRoom `json:"memberRoom"`
	Faction        int          `json:"faction"`
	WinExperience  int          `json:"winExperience"`
	LoseExperience int          `json:"loseExperience"`
	WinScore       int          `json:"winScore"`
	LoseScore      int          `json:"loseScore"`
}
