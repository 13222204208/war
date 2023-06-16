package copilot

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/business"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
)

// 根据code获取openid
func GetOpenId(code string) (openid string, err error) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	appid := global.GVA_CONFIG.Mini.Appid
	secret := global.GVA_CONFIG.Mini.Secret
	fmt.Println("appid:", appid)
	fmt.Println("secret:", secret)
	cfg := &config.Config{
		AppID:     appid,
		AppSecret: secret,
		Cache:     memory,
	}
	min := wc.GetMiniProgram(cfg)
	a := min.GetAuth()
	r, err := a.Code2Session(code)

	if err != nil {
		return openid, err
	} else {
		return r.OpenID, err
	}
}

// 获取二维码
func GetQrCode(page, scene string) (url string, err error) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	appid := global.GVA_CONFIG.Mini.Appid
	secret := global.GVA_CONFIG.Mini.Secret
	fmt.Println("appid:", appid)
	fmt.Println("secret:", secret)
	cfg := &config.Config{
		AppID:     appid,
		AppSecret: secret,
		Cache:     memory,
	}
	min := wc.GetMiniProgram(cfg)
	a := min.GetQRCode()
	var coderParams qrcode.QRCoder
	coderParams.Page = page
	coderParams.Scene = scene

	r, err := a.GetWXACodeUnlimit(coderParams)
	if err != nil {
		return url, err
	} else {
		//二维码二进制数据转为图片
		qrcodeFilePath := "uploads/file/" + scene + ".jpg"
		err = ioutil.WriteFile(qrcodeFilePath, r, 0666)
		if err != nil {
			return url, err
		}
		return qrcodeFilePath, err
	}
}

// 获取微信小程序用户的手机号
func GetWeChatPhone(code string) (phone string, err error) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	appid := global.GVA_CONFIG.Mini.Appid
	secret := global.GVA_CONFIG.Mini.Secret
	fmt.Println("appid:", appid)
	fmt.Println("secret:", secret)
	cfg := &config.Config{
		AppID:     appid,
		AppSecret: secret,
		Cache:     memory,
	}
	min := wc.GetMiniProgram(cfg)
	b := min.GetBusiness()
	var p business.GetPhoneNumberRequest
	p.Code = code
	info, err := b.GetPhoneNumber(&p)
	if err != nil {
		fmt.Println("获取微信小程序用户手机号失败", err)
		return "", errors.New("获取微信小程序用户手机号失败")
	}
	phone = info.PurePhoneNumber
	return
}
