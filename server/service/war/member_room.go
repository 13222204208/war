package war

import (
	"errors"
	"fmt"
	"time"

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
	//更新房间的实到人数
	err = global.GVA_DB.Model(&room).Where("id = ?", room.ID).Update("actual_num_players", gorm.Expr("actual_num_players + ?", 1)).Error

	return
}

// 身份证签到
func (memberRoomService *MemberRoomService) IdCardSign(idCard string) (err error) {
	//查询已经开始的房间
	var room war.Room
	err = global.GVA_DB.Where("status = ?", 2).First(&room).Error
	if err != nil {
		return errors.New("当前没有游戏正在进行")
	}

	//根据身份证查询用户id
	var member war.Member
	err = global.GVA_DB.Where("id_card = ?", idCard).First(&member).Error
	if err != nil {
		return errors.New("身份证没有查询到相关用户信息")
	}
	userId := member.ID
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
func (memberRoomService *MemberRoomService) GetMemberRoomListByUserId(userId uint) (res []warRes.GameRecordResponse, err error) {
	// 查询游戏记录获取用户的比分，比赛时间，比赛结果
	var games []war.GameRecord
	err = global.GVA_DB.
		Preload("Room").
		Where("user_id = ?", userId).
		Order("created_at DESC").
		Find(&games).Error
	if err != nil {
		return res, err
	}
	if len(games) == 0 {
		return res, errors.New("你还没有对局记录")
	}

	// 构造响应数据
	for _, game := range games {
		var gameRecord warRes.GameRecordResponse
		gameRecord.RoomId = game.RoomId
		gameRecord.GameTime = game.CreatedAt.Format("2006-01-02")
		if game.Faction == 1 {
			gameRecord.Score = fmt.Sprintf("%d:%d", game.Room.RedScore, game.Room.BlueScore)
		} else {
			gameRecord.Score = fmt.Sprintf("%d:%d", game.Room.BlueScore, game.Room.RedScore)
		}
		gameRecord.GameResult = game.GameResult
		gameRecord.GameType = game.GameType
		res = append(res, gameRecord)
	}

	return res, err
}

// 获取用户对局详情
func (memberRoomService *MemberRoomService) GetMemberRoomInfoByRoomId(roomId, userId int) (redAndBlue warRes.MemberRoomResponse, err error) {
	//查询你的阵营
	var memberRoom war.MemberRoom
	err = global.GVA_DB.Where("room_id = ? AND user_id = ?", roomId, userId).First(&memberRoom).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return redAndBlue, errors.New("你未加入房间")
		}
		return
	}
	redAndBlue.Faction = memberRoom.Faction

	//faction 1 为红方 2 为蓝方 分组返回双方信息
	var m []war.MemberRoom

	err = global.GVA_DB.Where("room_id = ? ", roomId).Preload("User").Find(&m).Error
	if err != nil {
		return
	}
	if len(m) > 0 {
		for _, v := range m {
			var MemberInfo warRes.MemberRoomInfo
			MemberInfo.Id = v.UserId
			MemberInfo.Avatar = v.User.Avatar
			MemberInfo.Nickname = v.User.Nickname
			MemberInfo.Kda = v.User.Kda
			MemberInfo.TeamName = v.User.TeamName
			if v.Faction == 1 {
				redAndBlue.Red = append(redAndBlue.Red, MemberInfo)
			} else {
				redAndBlue.Blue = append(redAndBlue.Blue, MemberInfo)
			}
		}
	}
	//根据房间id查询房间信息的游戏结束时间
	var room war.Room
	err = global.GVA_DB.Where("id = ?", roomId).First(&room).Error
	if err != nil {
		return
	}
	//获取游戏结束15分钟之后的时间
	redAndBlue.EndTime = room.EndTime.Add(time.Minute * 15).Format("2006-01-02 15:04:05")
	return
}

// 获取快速匹配的房间信息
func (memberRoomService *MemberRoomService) GetQuickMatchRoomInfo(userId uint) (roomInfo warRes.RoomResponse, err error) {
	//查询已经开始的房间
	room, err := GetMemberRoomByUserId(userId)
	if err != nil {
		return
	}
	fmt.Println("房间信息", room.ID)
	//faction 1 为红方 2 为蓝方 分组返回双方信息
	var red []war.MemberRoom
	var blue []war.MemberRoom
	var redResponse []warRes.MemberInfo
	var blueResponse []warRes.MemberInfo

	global.GVA_DB.Where("room_id = ? AND faction = ?", room.ID, 1).Preload("User").Find(&red)
	global.GVA_DB.Where("room_id = ? AND faction = ?", room.ID, 2).Preload("User").Find(&blue)
	fmt.Println("红方", red[0].User)
	redResponse = getMemberInfo(red)
	blueResponse = getMemberInfo(blue)

	roomInfo.Countdown = room.EndTime.Format("2006-01-02 15:04:05")
	roomInfo.Red = redResponse
	roomInfo.RoomId = room.ID
	roomInfo.Blue = blueResponse
	return
}

func getMemberInfo(members []war.MemberRoom) []warRes.MemberInfo {
	var result []warRes.MemberInfo
	for _, v := range members {
		info := warRes.MemberInfo{
			Avatar:   v.User.Avatar,
			Nickname: v.User.Nickname,
			TeamName: v.User.TeamName,
			Kda:      v.User.Kda,
		}
		result = append(result, info)
	}
	return result
}

// 查询用户是否有准备中的游戏
func GetMemberRoomByUserId(userId uint) (room war.Room, err error) {
	//查询已经开始的房间
	err = global.GVA_DB.Where("status = ?", 1).First(&room).Error
	if err != nil {
		return room, errors.New("当前没有游戏在准备中")
	}
	//查询用户是否已经加入了房间的记录
	var memberRoom war.MemberRoom
	err = global.GVA_DB.Where("user_id = ? AND room_id = ?", userId, room.ID).First(&memberRoom).Error
	if err != nil {
		return room, errors.New("你未加入房间")
	}
	return room, nil
}

// 根据房间id获取房间成员人数及阵营
func GetMemberRoomInfoByRoomId(roomId uint) (total, faction int, err error) {
	//统计房间人数及Faction 为1 的人数和为2的人数
	var memberRooms []war.MemberRoom
	err = global.GVA_DB.Where("room_id = ?", roomId).Find(&memberRooms).Error
	if err != nil {
		return
	}
	//平均分配faction的值，双方阵营只差1人
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
		faction = 2
	} else {
		faction = 1
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
