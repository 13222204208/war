package response

// 返回红蓝双方信息
type MemberRoomResponse struct {
	//投诉结束时间
	EndTime string `json:"endTime"`
	//阵营
	Faction int `json:"faction"`
	//红方
	Red []MemberRoomInfo `json:"red"`
	//蓝方
	Blue []MemberRoomInfo `json:"blue"`
}

// 对局详情的用户信息
type MemberRoomInfo struct {
	Id uint `json:"id"`
	MemberInfo
}

// 返回房间信息
type RoomResponse struct {
	RoomId uint `json:"roomId"`
	//倒计时
	Countdown string `json:"countdown"`
	//红方
	Red []MemberInfo `json:"red"`
	//蓝方
	Blue []MemberInfo `json:"blue"`
}

type MemberInfo struct {
	Avatar   string  `json:"avatar"`
	Nickname string  `json:"nickname"`
	Kda      float64 `json:"kda"`
	TeamName string  `json:"teamName"`
}

// 对局列表
type GameRecordResponse struct {
	//房间Id
	RoomId uint `json:"roomId"`
	//比赛结果
	GameResult int `json:"gameResult"`
	//比赛类型
	GameType int `json:"gameType"`
	//比赛时间
	GameTime string `json:"gameTime"`
	//比分
	Score string `json:"score"`
}
