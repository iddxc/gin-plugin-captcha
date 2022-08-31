package captcha

import (
	"gin-plugin-captcha/global"
	"gin-plugin-captcha/initialize"

	"github.com/gin-gonic/gin"
)

type captchaPlugin struct{}

func CreateCaptchaPlugin(Addr, Password string, DB, Mode int) *captchaPlugin {
	global.GVA_CONFIG.Redis.Addr = Addr
	global.GVA_CONFIG.Redis.Password = Password
	global.GVA_CONFIG.Redis.DB = DB
	initialize.Redis()
	return &captchaPlugin{}
}

func (*captchaPlugin) Register(group *gin.RouterGroup) {
	// router.RouterGroupApp.InitChatRoomRouter(group)
}

func (*captchaPlugin) RouterPath() string {
	return "captcha"
}
