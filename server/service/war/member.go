package war

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
	warRes "github.com/flipped-aurora/gin-vue-admin/server/model/war/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/copilot"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MemberService struct {
}

// CreateMember 创建Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) CreateMember(member *war.Member) (err error) {
	err = global.GVA_DB.Create(member).Error
	return err
}

func (memberService *MemberService) ProcessExcelFile(filePath string) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("错误一", err)
		return err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("会员资料1")
	if err != nil {
		fmt.Println("错误三", err)
		return err
	}
	//保存数据到会员表
	var member []war.Member
	for _, row := range rows {
		m := war.Member{}
		for k, colCell := range row {
			// fmt.Print(k, "列", colCell, "\t")
			CopyMember(k, colCell, &m)
		}
		//添加m到member
		fmt.Println("m数据", m)

		member = append(member, m)

	}
	err = global.GVA_DB.Create(&member).Error
	fmt.Println("结果", err)
	return err
}

// CopyMember 复制Member记录
func CopyMember(k int, colCell string, m *war.Member) (err error) {
	g := new(int)
	switch k {
	case 1:
		if colCell == "男" {
			*g = 1
		} else {
			*g = 2
		}
		m.Gender = g
	case 2:
		m.Phone = colCell
	case 0, 3, 4, 6, 7:
		switch k {
		case 0:
			m.Name = colCell
		case 3:
			colCell = strconv.Itoa(GetMemberLevel(colCell))
			memberLevelId, err := strconv.ParseUint(colCell, 10, 32)
			if err != nil {
				fmt.Println("错误五", err)
			}
			m.MemberLevelId = uint(memberLevelId)
		case 4:
			m.Score, err = strconv.Atoi(colCell)
		case 6:
			m.IdCard = colCell
		case 7:
			//colCell转为uint
			age, err := strconv.ParseUint(colCell, 10, 32)
			if err != nil {
				fmt.Println("错误六", err)
				return err
			}
			match := uint(age)
			m.Match = &match
		}
		if err != nil {
			fmt.Println("错误六", err)
			return err
		}
	case 5:
		match, err := strconv.ParseUint(colCell, 10, 32)
		if err != nil {
			fmt.Println("错七", err)
			return err
		}
		matchVal := uint(match)
		m.Match = &matchVal
	}
	return err
}

// 判断会员的等级
func GetMemberLevel(s string) (level int) {
	switch s {
	case "职业玩家卡":
		level = 5
	case "竞技玩家卡":
		level = 6
	case "进阶玩家卡":
		level = 7
	case "军警卡":
		level = 8
	case "畅玩玩家卡":
		level = 9
	case "体验玩家卡":
		level = 10
	case "新手入门卡":
		level = 11
	}
	return
}

// DeleteMember 删除Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) DeleteMember(member war.Member) (err error) {
	err = global.GVA_DB.Delete(&member).Error
	//删除关联的战队成员信息
	if err == nil {
		err = global.GVA_DB.Where("user_id = ?", member.ID).Delete(&war.TeamMember{}).Error
		//战队的人数减一D
		if err == nil {
			err = global.GVA_DB.Model(&war.Team{}).Where("id = ?", member.TeamID).Update("team_member_num", gorm.Expr("team_member_num - ?", 1)).Error
		}
	}

	return err
}

// DeleteMemberByIds 批量删除Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) DeleteMemberByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]war.Member{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMember 更新Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) UpdateMember(member war.Member) (err error) {
	err = global.GVA_DB.Save(&member).Error
	return err
}

// GetMember 根据id获取Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) GetMember(id uint) (member war.Member, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&member).Error
	return
}

// GetMemberInfoList 分页获取Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) GetMemberInfoList(info warReq.MemberSearch) (list []war.Member, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&war.Member{})
	var members []war.Member
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	//如果手机号不为空
	if info.Phone != "" {
		db = db.Where("phone = ?", info.Phone)
	}
	//如果姓名不为空
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// err = db.Limit(limit).Offset(offset).Find(&members).Error
	err = db.Table("war_member").Preload("MemberLevel").Order("id DESC").Limit(limit).Offset(offset).Find(&members).Error
	return members, total, err
}

// 登陆
// Login logs in a member with a given code and phone number
func (memberService *MemberService) Login(l warReq.WechatLogin) (member war.Member, err error) {

	err = global.GVA_DB.Where("phone = ?", l.Phone).First(&member).Error
	fmt.Println("登陆err错误", err)

	if err == gorm.ErrRecordNotFound {
		//如果头像和昵称为空
		if l.Avatar == "" && l.Nickname == "" {
			return member, errors.New("头像和昵称不能为空")
		}

		openid, openidErr := copilot.GetOpenId(l.Code)
		if openidErr != nil {
			return member, openidErr
		}
		if openid == "" {
			return member, errors.New("没有获取到openid")
		}
		member = war.Member{
			Openid:   openid,
			Phone:    l.Phone,
			Avatar:   l.Avatar,
			Nickname: l.Nickname,
		}
		err = global.GVA_DB.Create(&member).Error
		if err != nil {
			return
		}
	}
	member, err = memberService.GetMemberInfo(member.ID)
	return
}

// 会员修改信息
func (memberService *MemberService) UpdateMemberInfo(userID uint, member war.Member) (err error) {
	err = global.GVA_DB.Model(&war.Member{}).Where("id = ?", userID).Updates(member).Error
	return err
}

// 获取会员资料
func (memberService *MemberService) GetMemberInfo(userID uint) (member war.Member, err error) {
	var m war.Member
	err = global.GVA_DB.Where("id = ?", userID).Preload("RankLevel").Preload("MemberLevel").First(&m).Error
	if err != nil {
		//判断数据是否存在
		if err == gorm.ErrRecordNotFound {
			return m, errors.New("用户不存在")
		}
		return
	}
	if m.Exp == 0 {
		m.Exp = 0
		m.RankLevel.Experience = 21
		return m, err
	}

	//查询当前军衔等级的经验和下一个等级的经验之差
	exp1, exp2, err := GetMemberRankLevel(m.RankLevel.ID, m.Exp)
	if err != nil {
		return m, err
	} else {
		m.Exp = exp1
		m.RankLevel.Experience = exp2
	}

	return m, err
}

// 计算当前军衔等级的经验和下一个等级的经验
func GetMemberRankLevel(levelId uint, exp int) (exp1, exp2 int, err error) {
	var rankLevel war.RankLevel
	err = global.GVA_DB.Where("id = ?", levelId).First(&rankLevel).Error
	if err != nil {
		return
	}

	var r war.RankLevel
	err = global.GVA_DB.Where("experience > ?", exp).Order("experience ASC").First(&r).Error
	if err != nil {
		return
	}

	fmt.Println("当前军衔等级的经验和下一个等级的经验", rankLevel.Experience, r.Experience)
	fmt.Println("我的当前经验", exp)
	exp1 = exp - rankLevel.Experience
	exp2 = r.Experience - rankLevel.Experience
	return
}

// 会员增加或减少场次
func (memberService *MemberService) AddOrUpdateMemberMatch(userID, match, matchType uint) (err error) {
	global.GVA_LOG.Info("场次类型", zap.Any("matchType", matchType))
	if matchType == 1 {
		err = AddUserMatch(userID, match, "后台增加场次")
		if err != nil {
			return err
		}
	} else if matchType == 2 {
		err = DeductUserMatch(userID, match, "后台减少场次")
		if err != nil {
			return err
		}
	} else {
		return errors.New("场次类型必须为1或者2")
	}
	return err
}

// 获取我的kda
func (memberService *MemberService) GetMyKda(userID uint) (response warRes.MyKdaResponse, err error) {
	//当前更新时间
	response.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	//统计游戏记录中的game_result 等于1和2各自的次数
	var gameRecord []war.GameRecord
	err = global.GVA_DB.Model(&war.GameRecord{}).Where("user_id = ?", userID).Find(&gameRecord).Error
	if err != nil {
		return response, err
	}
	var win, lose int
	if len(gameRecord) > 0 {
		for _, v := range gameRecord {
			if v.GameResult == 1 {
				win++
			} else {
				lose++
			}
		}
	}
	response.Win = win
	response.Lose = lose
	if len(gameRecord) == 0 {
		response.Kda = 0
	} else {
		kda := float64(win) / float64(win+lose)
		//只取小数点后两位
		response.Kda, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", kda), 64)
		err = UpdateMyKda(userID, response.Kda)
	}
	return response, err
}

// 更新我的kda
func UpdateMyKda(userID uint, kda float64) (err error) {
	//更新会员的kda
	err = global.GVA_DB.Model(&war.Member{}).Where("id = ?", userID).Update("kda", kda).Error
	if err != nil {
		return err
	}

	//查询战队成员表中的kda
	var teamMember war.TeamMember
	err = global.GVA_DB.Model(&war.TeamMember{}).Where("user_id = ?", userID).First(&teamMember).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	teamMember.DamageRatio = kda
	err = global.GVA_DB.Model(&war.TeamMember{}).Where("user_id = ?", userID).Update("damage_ratio", kda).Error
	return err
}

// 我的战斗信息
// GetMyBattleInfo 获取我的战斗信息
func (memberService *MemberService) GetMyBattleInfo(userID uint) (response warRes.MyBattleResponse, err error) {
	var gameRecord []war.GameRecord
	if err := global.GVA_DB.Model(&war.GameRecord{}).Where("user_id = ?", userID).Find(&gameRecord).Error; err != nil {
		return response, err
	}

	response.Count = len(gameRecord)
	fmt.Println("我的战斗记录", gameRecord)
	for _, v := range gameRecord {
		if v.GameResult == 1 {
			response.Win++
		} else {
			response.Lose++
		}
	}

	if response.Count == 0 {
		response.Kda = 0
	} else {
		response.Kda = float64(response.Win) / float64(response.Count)
		response.Kda, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", response.Kda), 64)
	}

	return response, nil
}

// 计算kda
func CalculateKda(userID uint) (kda float64, err error) {
	//统计游戏记录中的game_result 等于1和2各自的次数
	var gameRecord []war.GameRecord
	err = global.GVA_DB.Model(&war.GameRecord{}).Where("user_id = ?", userID).Find(&gameRecord).Error
	if err != nil {
		return 0, err
	}
	var win, lose int
	if len(gameRecord) > 0 {
		for _, v := range gameRecord {
			if v.GameResult == 1 {
				win++
			} else {
				lose++
			}
		}
	}
	kda = float64(win) / float64(win+lose)
	//只取小数点后两位
	kda, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", kda), 64)
	return kda, err
}

// 获取会员手机号
func (memberService *MemberService) GetMemberPhone(code string) (phone string, err error) {
	phone, err = copilot.GetWeChatPhone(code)
	return phone, err
}

// 会员增加场次
func AddUserMatch(userID, match uint, remark string) (err error) {
	var member war.Member
	err = global.GVA_DB.Where("id = ?", userID).First(&member).Error
	if err != nil {
		return err
	} else {
		*member.Match += match
		err = global.GVA_DB.Save(&member).Error
		if err != nil {
			return err
		}
	}
	var record war.MatchRecord
	record.UserId = userID
	record.MatchNum = match
	record.MatchType = 1
	record.Remark = remark
	err = global.GVA_DB.Create(&record).Error
	return err
}

// 个人详情
func (memberService *MemberService) GetMemberDetail(userID uint) (response warRes.MemberResponse, err error) {
	//查询用户的kda
	var member war.Member
	err = global.GVA_DB.Where("id = ?", userID).First(&member).Error
	if err != nil {
		return response, errors.New("用户不存在")
	}
	response.Avatar = member.Avatar
	response.Nickname = member.Nickname
	response.DamageRatio = member.Kda

	//判断有没有战队
	if *member.TeamID > 0 {
		var tm war.TeamMember
		err = global.GVA_DB.Where("user_id = ?", userID).Preload("TeamInfo").Preload("TeamRoleInfo").First(&tm).Error
		if err != nil {
			return response, err
		}
		response.TeamName = tm.TeamInfo.Name
		response.TeamLogo = tm.TeamInfo.Logo
		response.Role = tm.TeamRoleInfo.Role
	}

	//查询用户的kda排名
	rank, err := GetMemberKdaRank(userID)
	if err != nil {
		return response, err
	}
	response.WinRateRank = rank
	//获取个人装备
	response.Equipments, err = GetMemberEquip(userID)
	return response, err
}

// 查询用户kda的排名越大的越靠前
func GetMemberKdaRank(userID uint) (rank int, err error) {
	var members []war.Member
	err = global.GVA_DB.Order("kda desc").Find(&members).Error
	if err != nil {
		return rank, err
	}
	for i, v := range members {
		if v.ID == userID {
			rank = i + 1
			break
		}
	}
	return rank, err
}

// 查询个人装备
func GetMemberEquip(userID uint) (m []map[string]interface{}, err error) {
	var response []warRes.EquipmentResponse
	//查询用户装备
	var userEquip []war.UserEquipment
	err = global.GVA_DB.Where("user_id = ?", userID).Preload("Category").Preload("Equipment").Find(&userEquip).Error
	if err != nil {
		return m, err
	}
	for _, v := range userEquip {
		var equip warRes.EquipmentResponse
		equip.Name = v.Equipment.Name
		equip.CategoryName = v.Category.Name
		equip.Icon = v.Equipment.Icon
		response = append(response, equip)
	}
	//装备分类分组
	m, err = GetEquipCategory(response)
	return m, err
}

// 装备分类分组
func GetEquipCategory(response []warRes.EquipmentResponse) (m []map[string]interface{}, err error) {

	categories := make(map[string][]map[string]string)

	for _, equipment := range response {
		category := equipment.CategoryName
		item := map[string]string{
			"name": equipment.Name,
			"icon": equipment.Icon,
		}

		if _, ok := categories[category]; !ok {
			categories[category] = []map[string]string{}
		}

		categories[category] = append(categories[category], item)
	}

	result := []map[string]interface{}{}

	for categoryName, children := range categories {
		category := map[string]interface{}{
			"categoryName": categoryName,
			"children":     children,
		}

		result = append(result, category)
	}
	return result, err
}

// 按年或月查询游戏记录并排行
func (memberService *MemberService) GetMemberRank(t string) (res []warRes.MemberRankResponse, err error) {
	// 根据用户积分排行
	records, err := GetMemberRecordRank(t)
	if err != nil {
		return res, err
	}
	response := make([]warRes.MemberRankResponse, len(records))
	// 获取用户排名
	for rank, r := range records {
		//如果战队名称不为空，则查询出战队logo
		var logo string
		if r.User.TeamName != "" {
			var team war.Team
			err = global.GVA_DB.Where("name = ?", r.User.TeamName).First(&team).Error
			if err != nil {
				return res, err
			}
			logo = team.Logo
		}
		res := warRes.MemberRankResponse{
			UserId:   r.UserId,
			Nickname: r.User.Nickname,
			Avatar:   r.User.Avatar,
			TeamName: r.User.TeamName,
			TeamLogo: logo,
			Kda:      r.User.Kda,
			Score:    r.User.Score,
			Rank:     rank + 1,
		}

		// var userEquip []war.UserEquipment
		// global.GVA_DB.Where("user_id = ?", r.UserId).Preload("Equipment").Find(&userEquip)
		// equipments := make([]warRes.EquipmentIcon, 0, len(userEquip))
		// for _, v := range userEquip {
		// 	equipments = append(equipments, warRes.EquipmentIcon{Icon: v.Equipment.Icon})
		// }
		//我的装备详情
		var eq *EquipmentService
		equipments, err := eq.Detail(r.UserId)
		fmt.Println("获取我的装备错误", err)
		res.Equipments = equipments
		response[rank] = res
	}
	return response, nil
}

// 根据状态查询用户游戏记录
func GetMemberRecordRank(queryType string) (records []war.GameRecord, err error) {
	var startTime, endTime time.Time
	now := time.Now()
	// 本年的起始时间和结束时间
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	endOfYear := startOfYear.AddDate(1, 0, 0).Add(-time.Second)

	// 本月的起始时间和结束时间
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)

	// 判断查询的时间范围
	if queryType == "year" {
		startTime = startOfYear
		endTime = endOfYear
	} else if queryType == "month" {
		startTime = startOfMonth
		endTime = endOfMonth
	} else {
		// 查询所有时间
		startTime = time.Time{} // 时间的零值，即不限制开始时间
		endTime = now
	}

	// 查询积分排行
	err = global.GVA_DB.Model(&war.GameRecord{}).
		Select("user_id, SUM(score) as total_score").
		Where("created_at >= ? AND created_at <= ?", startTime, endTime).
		Group("user_id").
		Order("total_score DESC").
		Preload("User").
		Find(&records).Error
	if err != nil {
		return records, err
	}
	return records, nil
}
