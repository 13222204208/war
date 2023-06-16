package war

import (
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/war"
	warReq "github.com/flipped-aurora/gin-vue-admin/server/model/war/request"
	warRes "github.com/flipped-aurora/gin-vue-admin/server/model/war/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MemberApi struct {
}

var memberService = service.ServiceGroupApp.WarServiceGroup.MemberService

// CreateMember 创建Member
// @Tags Member
// @Summary 创建Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Member true "创建Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /member/createMember [post]
func (memberApi *MemberApi) CreateMember(c *gin.Context) {
	var member war.Member
	err := c.ShouldBindJSON(&member)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := memberService.CreateMember(&member); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// 导入excel
func (memberApi *MemberApi) ImportExcel(c *gin.Context) {

	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	filepath := "uploads/file/" + "excelImport.xlsx"
	_ = c.SaveUploadedFile(header, filepath)

	if err := memberService.ProcessExcelFile(filepath); err != nil {
		global.GVA_LOG.Error("导入失败!", zap.Error(err))
		response.FailWithMessage("导入失败", c)
		return
	} else {
		response.OkWithMessage("导入成功", c)
	}
}

// DeleteMember 删除Member
// @Tags Member
// @Summary 删除Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Member true "删除Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /member/deleteMember [delete]
func (memberApi *MemberApi) DeleteMember(c *gin.Context) {
	var member war.Member
	err := c.ShouldBindJSON(&member)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := memberService.DeleteMember(member); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMemberByIds 批量删除Member
// @Tags Member
// @Summary 批量删除Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /member/deleteMemberByIds [delete]
func (memberApi *MemberApi) DeleteMemberByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := memberService.DeleteMemberByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMember 更新Member
// @Tags Member
// @Summary 更新Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body war.Member true "更新Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /member/updateMember [put]
func (memberApi *MemberApi) UpdateMember(c *gin.Context) {
	var member war.Member
	err := c.ShouldBindJSON(&member)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := memberService.UpdateMember(member); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMember 用id查询Member
// @Tags Member
// @Summary 用id查询Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query war.Member true "用id查询Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /member/findMember [get]
func (memberApi *MemberApi) FindMember(c *gin.Context) {
	var member war.Member
	err := c.ShouldBindQuery(&member)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remember, err := memberService.GetMember(member.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remember": remember}, c)
	}
}

// GetMemberList 分页获取Member列表
// @Tags Member
// @Summary 分页获取Member列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query warReq.MemberSearch true "分页获取Member列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /member/getMemberList [get]
func (memberApi *MemberApi) GetMemberList(c *gin.Context) {
	var pageInfo warReq.MemberSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := memberService.GetMemberInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// 会员登陆
func (memberApi *MemberApi) Login(c *gin.Context) {
	var l warReq.WechatLogin
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(l.Phone) != 11 {
		response.FailWithMessage("手机号不正确", c)
		return
	}
	if user, err := memberService.Login(l); err != nil {
		global.GVA_LOG.Error("登陆失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		memberApi.TokenNext(c, user)
	}
}

// 会员修改资料
func (memberApi *MemberApi) UpdateMemberInfo(c *gin.Context) {
	var member war.Member
	err := c.ShouldBindJSON(&member)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if err := memberService.UpdateMemberInfo(userID, member); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// TokenNext 登录以后签发jwt
func (memberApi *MemberApi) TokenNext(c *gin.Context, user war.Member) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		ID: user.ID,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	fmt.Println("我的token", token)
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(warRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}
	response.OkWithDetailed(warRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)

}

// 获取会员资料
func (memberApi *MemberApi) GetMemberInfo(c *gin.Context) {
	userID := utils.GetUserID(c)
	fmt.Println("用户Id", userID)
	if user, err := memberService.GetMemberInfo(userID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"user": user}, c)
	}
}

// 会员增加或修改场次
func (memberApi *MemberApi) AddOrUpdateMemberMatch(c *gin.Context) {
	var match warReq.MemberMatch
	err := c.ShouldBindJSON(&match)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	global.GVA_LOG.Info("match:", zap.Any("match", match))
	userID := match.UserId
	if err := memberService.AddOrUpdateMemberMatch(userID, match.Match, match.MatchType); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// 个人详情
func (memberApi *MemberApi) GetMemberDetail(c *gin.Context) {
	userId := c.Query("userId")
	var userID uint
	if userId != "" {
		// Convert userId to uint data type
		id, err := strconv.ParseUint(userId, 10, 32)
		if err != nil {
			response.FailWithMessage("Invalid userId", c)
			return
		}
		userID = uint(id)
	} else {
		userID = utils.GetUserID(c)
	}

	if user, err := memberService.GetMemberDetail(userID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"user": user}, c)
	}
}

// 获取我的kda
func (memberApi *MemberApi) GetMyKda(c *gin.Context) {
	userID := utils.GetUserID(c)
	if kda, err := memberService.GetMyKda(userID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"info": kda}, c)
	}
}

// 获取我的战斗信息
func (memberApi *MemberApi) GetMyBattleInfo(c *gin.Context) {
	userID := utils.GetUserID(c)
	if matchInfo, err := memberService.GetMyBattleInfo(userID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"info": matchInfo}, c)
	}
}

// 用户排行
func (memberApi *MemberApi) GetMemberRank(c *gin.Context) {
	var t warReq.RankType
	err := c.ShouldBindQuery(&t)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rank, err := memberService.GetMemberRank(t.Type); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"rank": rank}, c)
	}
}

// 获取会员手机号
func (memberApi *MemberApi) GetMemberPhone(c *gin.Context) {
	var code warReq.WechatLogin
	err := c.ShouldBindJSON(&code)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if phone, err := memberService.GetMemberPhone(code.Code); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"phone": phone}, c)
	}
}
