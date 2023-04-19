package copilot

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
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
