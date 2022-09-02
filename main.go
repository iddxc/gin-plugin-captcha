package captcha

import (
	"github.com/iddxc/gin-plugin-captcha/global"
	"github.com/iddxc/gin-plugin-captcha/initialize"
	"github.com/iddxc/gin-plugin-captcha/router"

	"github.com/gin-gonic/gin"
)

type captchaPlugin struct{}

func CreateCaptchaPlugin(Addr, Password, FontFile string, DB int) *captchaPlugin {
	global.GVA_CONFIG.Redis.Addr = Addr
	global.GVA_CONFIG.Redis.Password = Password
	global.GVA_CONFIG.Redis.DB = DB
	global.GVA_CONFIG.FontFile = FontFile
	initialize.Redis()
	return &captchaPlugin{}
}

func (*captchaPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitCaptchaApiRouter(group)
}

func (*captchaPlugin) RouterPath() string {
	return "captcha"
}
