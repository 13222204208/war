package war

import (
	"errors"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/copilot"
	"gorm.io/gorm"
)

type RoomService struct {
}

// CreateRoom 创建Room记录
// Author [piexlmax](https://github.com/piexlmax)
func (roomService *RoomService) CreateRoom(room *war.Room) (err error) {
	err = global.GVA_DB.Create(room).Error
	return err
}

// DeleteRoom 删除Room记录
// Author [piexlmax](https://github.com/piexlmax)
func (roomService *RoomService) DeleteRoom(room war.Room) (err error) {
	err = global.GVA_DB.Delete(&room).Error
	return err
}

// DeleteRoomByIds 批量删除Room记录
// Author [piexlmax](https://github.com/piexlmax)
func (roomService *RoomService) DeleteRoomByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.Room{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRoom 更新Room记录
// Author [piexlmax](https://github.com/piexlmax)
func (roomService *RoomService) UpdateRoom(room war.Room) (err error) {
	err = global.GVA_DB.Save(&room).Error
	if err == nil {
		if room.Status == 3 {
			return
		}
		//计算分数，胜方加30分，败方加3分
		//获取房间信息
		var memberRoom []war.MemberRoom
		err = global.GVA_DB.Where("room_id = ?", room.ID).Find(&memberRoom).Error
		if err != nil {
			return err
		}
		var faction int
		loseScore := room.RedScore*global.GVA_CONFIG.Setting.WinScore + room.BlueScore*global.GVA_CONFIG.Setting.LoseScore
		winScore := room.BlueScore*global.GVA_CONFIG.Setting.WinScore + room.RedScore*global.GVA_CONFIG.Setting.LoseScore
		if room.RedScore > room.BlueScore {
			faction = 1
		} else {
			faction = 2
		}

		//给获胜的增加经验20，失败的12
		var g war.SaveGameRecord
		g.MemberRoom = memberRoom
		g.Faction = faction
		g.WinScore = winScore
		g.LoseScore = loseScore
		g.WinExperience = global.GVA_CONFIG.Setting.WinExperience
		g.LoseExperience = global.GVA_CONFIG.Setting.LoseExperience
		err = SaveGameRecord(&g)
		if err == nil {
			//更新用户的kda
			for _, v := range memberRoom {
				userId := v.UserId
				kda, err := CalculateKda(userId)
				if err != nil {
					return err
				}
				//更新用户的kda
				err = global.GVA_DB.Model(&war.Member{}).Where("id = ?", userId).Update("kda", kda).Error
				if err != nil {
					return err
				}
			}
		}
	}
	return err
}

// GetRoom 根据id获取Room记录
// Author [piexlmax](https://github.com/piexlmax)
func (roomService *RoomService) GetRoom(id uint) (room war.Room, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&room).Error
	return
}

// GetRoomInfoList 分页获取Room记录
// Author [piexlmax](https://github.com/piexlmax)
func (roomService *RoomService) GetRoomInfoList(info warReq.RoomSearch) (list []war.Room, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.Room{})
	var rooms []war.Room
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("id DESC").Find(&rooms).Error
	return rooms, total, err
}

// 快速匹配游戏
func (roomService *RoomService) QuickMatch(userId uint) (err error) {
	//判断用户的游戏场次是否足够
	ok := IsUserMatchFull(userId, 1)
	if !ok {
		return errors.New("场次不足")
	}
	//判断是否有空闲的房间
	roomId, err := GetFreeRoom()
	if err != nil {
		return err
	}
	//先判断房间人数及应该分配到哪个阵营
	total, faction, err := GetMemberRoomInfoByRoomId(roomId)
	if err != nil {
		return err
	}
	//如果房间人数小于24人，直接加入
	if total < 24 {
		err = AddMemberRoom(userId, roomId, faction)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("房间已满")
	}

}

// 获取空闲的房间
func GetFreeRoom() (roomId uint, err error) {
	var room war.Room
	err = global.GVA_DB.Where("status = ?", 1).First(&room).Error
	//如果没有则创建一个房间
	if err != nil {
		num, err := CountRoom()
		if err != nil {
			return 0, err
		}
		//name 等于当前日期加上 num+1
		name := time.Now().Format("2006-01-02") + "-" + strconv.Itoa(int(num+1))
		//15分钟之后的时间
		endTime := time.Now().Add(time.Minute * 15)
		room := war.Room{
			Name:    name,
			RoomNum: num + 1,
			EndTime: endTime,
			Status:  1,
		}
		err = global.GVA_DB.Create(&room).Error
		if err != nil {
			return 0, err
		}
		roomId = room.ID
	} else {
		if room.ID > 0 {
			roomId = room.ID
		}
	}
	return roomId, nil
}

// 判断是否有游戏中的房间和统计今天创建了多少房间
func CountRoom() (countRoom int64, err error) {
	//判断如果有游戏中的房间，刚无法创建房间
	var count int64
	err = global.GVA_DB.Model(&war.Room{}).Where("status = ?", 2).Count(&count).Error
	if err != nil {
		return
	}
	if count > 0 {
		return 0, errors.New("有游戏中的房间")
	}

	//统计今天创建了多少个房间
	err = global.GVA_DB.Model(&war.Room{}).Where("created_at BETWEEN ? AND ?", time.Now().Format("2006-01-02 00:00:00"), time.Now().Format("2006-01-02 23:59:59")).Count(&count).Error
	if err != nil {
		return
	}
	return count, nil
}

// 更改游戏状态，并更改倒计时结束时间
func (roomService *RoomService) StartGame(roomId uint, status int) (err error) {
	var room war.Room
	err = global.GVA_DB.Where("id = ?", roomId).First(&room).Error
	if err != nil {
		return err
	} else {
		if room.Status == 2 {
			return errors.New("游戏已开始")
		}
		if room.Status == 3 {
			return errors.New("游戏已结束")
		}
	}

	room.Status = status
	room.EndTime = time.Now()
	room.GameOverTime = time.Now().Add(time.Minute * 60)
	err = global.GVA_DB.Save(&room).Error
	return err
}

// 获取对战海报二维码
func (roomService *RoomService) GetRoomQrCode(roomId int) (qrCode string, err error) {
	var room war.Room
	err = global.GVA_DB.Where("id = ?", roomId).First(&room).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("房间不存在")
		}
		return "", err
	}
	//获取二维码pages/index/index
	qrCode, err = copilot.GetQrCode("pages/index/index", "room_id="+strconv.Itoa(roomId))
	if err != nil {
		return "", err
	}
	return qrCode, nil
}

// 倒计时结束 删除或开始房间
func (roomService *RoomService) Countdown() (err error) {
	var room war.Room
	err = global.GVA_DB.Where("status = ?", 1).First(&room).Error
	if err != nil {
		if err.Error() == "record not found" {
			return errors.New("没有准备中的房间")
		} else {
			return err
		}
	}
	//如果房间人数小于6人
	if room.NumPlayers < 6 {
		//删除房间
		err = global.GVA_DB.Delete(&room).Error
		if err != nil {
			return err
		}
	} else {
		//开始游戏
		err = roomService.StartGame(room.ID, 2)
		if err != nil {
			return err
		}
	}
	return
}
