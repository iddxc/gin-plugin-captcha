package service

import (
	"fmt"
	"gin-plugin-captcha/global"
	"gin-plugin-captcha/initialize"
	"testing"
	"time"
)

func init() {
	global.GVA_CONFIG.Redis.Addr = "120.26.196.124:6379"
	global.GVA_CONFIG.Redis.DB = 0
	global.GVA_CONFIG.Redis.Password = "520qp025"
	initialize.Redis()
}

var service = new(CaptchaService)

func TestCreateCaptchaToLibrary(t *testing.T) {
	for i := 0; i < 30; i++ {
		captcha := service.OriginalCaptcha.GetString(6)
		fmt.Println(i, captcha)
		time.Sleep(time.Second)
		service.LibraryCaptcha.Set(captcha)
	}
}
func TestGetCaptchaString(t *testing.T) {
	captcha := service.LibraryCaptcha.GetString()
	fmt.Println(captcha)
}

func TestGetImage(t *testing.T) {
	captcha := service.LibraryCaptcha.GetString()
	data, err := service.ImageCapture.GetImage(250, 80, 28, captcha)
	fmt.Println(data, err)
}
