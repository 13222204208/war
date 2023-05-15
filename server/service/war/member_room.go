package war

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
	warRes "github.com/flipped-aurora/gin-vue-admin/server/model/war/response"
	"gorm.io/gorm"
)

type MemberRoomService struct {
}

// CreateMemberRoom 创建MemberRoom记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberRoomService *MemberRoomService) CreateMemberRoom(memberRoom *war.MemberRoom) (err error) {
	err = global.GVA_DB.Create(memberRoom).Error
	return err
}

// DeleteMemberRoom 删除MemberRoom记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberRoomService *MemberRoomService) DeleteMemberRoom(memberRoom war.MemberRoom) (err error) {
	err = global.GVA_DB.Delete(&memberRoom).Error
	return err
}

// DeleteMemberRoomByIds 批量删除MemberRoom记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberRoomService *MemberRoomService) DeleteMemberRoomByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.MemberRoom{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMemberRoom 更新MemberRoom记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberRoomService *MemberRoomService) UpdateMemberRoom(memberRoom war.MemberRoom) (err error) {
	err = global.GVA_DB.Save(&memberRoom).Error
	return err
}

// GetMemberRoom 根据id获取MemberRoom记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberRoomService *MemberRoomService) GetMemberRoom(id uint) (memberRoom war.MemberRoom, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&memberRoom).Error
	return
}

// GetMemberRoomInfoList 分页获取MemberRoom记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberRoomService *MemberRoomService) GetMemberRoomInfoList(info warReq.MemberRoomSearch) (list []war.MemberRoom, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.MemberRoom{})
	var memberRooms []war.MemberRoom
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&memberRooms).Error
	return memberRooms, total, err
}

// 用户签到
func (memberRoomService *MemberRoomService) Sign(userId uint) (err error) {
	//查询已经开始的房间
	var room war.Room
	err = global.GVA_DB.Where("status = ?", 2).First(&room).Error
	if err != nil {
		return errors.New("当前没有游戏正在进行")
	}
	//查询用户是否已经加入了房间的记录
	var memberRoom war.MemberRoom
	err = global.GVA_DB.Where("user_id = ? AND room_id = ?", userId, room.ID).First(&memberRoom).Error
	if err != nil {
		return errors.New("你未加入房间")
	}
	//判断用户是否已经签到
	if memberRoom.SignStatus == 1 {
		return errors.New("你已经签到")
	}
	//更新用户签到状态
	err = global.GVA_DB.Model(&memberRoom).Update("sign_status", 1).Error
	if err != nil {
		return
	}
	return
}

// 对局列表
func (memberRoomService *MemberRoomService) GetMemberRoomListByUserId(userId uint) (list []war.MemberRoom, err error) {
	err = global.GVA_DB.Where("user_id = ?", userId).Preload("Romm").Find(&list).Error
	return
}

// 获取用户对局详情
func (memberRoomService *MemberRoomService) GetMemberRoomInfoByRoomId(roomId int) (redAndBlue warRes.MemberRoomResponse, err error) {
	//faction 1 为红方 2 为蓝方 分组返回双方信息
	var red []war.MemberRoom
	var blue []war.MemberRoom
	err = global.GVA_DB.Where("room_id = ? AND faction = ?", roomId, 1).Preload("User").First(&red).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Where("room_id = ? AND faction = ?", roomId, 2).Preload("User").First(&blue).Error
	if err != nil {
		return
	}
	redAndBlue.Red = red
	redAndBlue.Blue = blue
	return
}

// 根据房间id获取房间成员人数及阵营
func GetMemberRoomInfoByRoomId(roomId uint) (total, faction int, err error) {
	//统计房间人数及Faction 为1 的人数和为2的人数
	var memberRooms []war.MemberRoom
	err = global.GVA_DB.Where("room_id = ?", roomId).Find(&memberRooms).Error
	if err != nil {
		return
	}
	var faction1, faction2 int
	for _, memberRoom := range memberRooms {
		if memberRoom.Faction == 1 {
			faction1++
		} else {
			faction2++
		}
	}
	total = len(memberRooms)
	if faction1 > faction2 {
		faction = 1
	} else {
		faction = 2
	}
	return total, faction, err
}

// 用户加入房间
func AddMemberRoom(userId, roomId uint, faction int) (err error) {
	//查询用户是否已经加入了房间的记录
	var memberRoom war.MemberRoom
	err = global.GVA_DB.Where("user_id = ? AND room_id = ?", userId, roomId).First(&memberRoom).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		memberRoom = war.MemberRoom{
			RoomId:  roomId,
			UserId:  userId,
			Faction: faction,
		}
		err = global.GVA_DB.Create(&memberRoom).Error
		if err != nil {
			return
		}
	} else {
		if memberRoom.ID != 0 {
			return errors.New("已经加入了房间")
		}
	}
	//房间人数加一
	err = AddMemberRoomNum(roomId)
	if err != nil {
		return
	}
	err = DeductUserMatch(userId, 1, "快速匹配游戏扣除场次")
	return
}

// 房间的人数加一
func AddMemberRoomNum(roomId uint) (err error) {
	var room war.Room
	err = global.GVA_DB.Where("id = ?", roomId).First(&room).Error
	if err != nil {
		return
	}
	room.NumPlayers++
	err = global.GVA_DB.Save(&room).Error
	return
}
