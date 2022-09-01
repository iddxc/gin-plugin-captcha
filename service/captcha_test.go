package service

import (
	"fmt"
	"gin-plugin-captcha/global"
	"gin-plugin-captcha/initialize"
	"testing"
	"time"
)

func init() {
	global.GVA_CONFIG.Redis.Addr = "127.0.0.1:6379"
	global.GVA_CONFIG.Redis.DB = 0
	global.GVA_CONFIG.Redis.Password = ""
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
func TestGetCaptcha(t *testing.T) {
	keyId, data, err := service.GetCaptcha("library", 5, 200, 80, 24)
	fmt.Println("library:", keyId, data, err)

	keyId, data, err = service.GetCaptcha("original", 5, 200, 80, 24)
	fmt.Println("original:", keyId, data, err)

	keyId, data, err = service.GetCaptcha("digit", 5, 200, 80, 24)
	fmt.Println("digit:", keyId, data, err)

	keyId, data, err = service.GetCaptcha("library+image", 5, 200, 80, 24)
	fmt.Println("library+image:", keyId, data, err)

	keyId, data, err = service.GetCaptcha("original+image", 5, 200, 80, 24)
	fmt.Println("original+image", keyId, data, err)

	keyId, data, err = service.GetCaptcha("digit+image", 5, 200, 80, 24)
	fmt.Println("digit+image:", keyId, data, err)
}

func TestGetImage(t *testing.T) {
	captcha := service.LibraryCaptcha.GetString()
	data, err := service.ImageCapture.GetImage(250, 80, 28, captcha)
	fmt.Println(data, err)
}

func TestAddCaptcha(t *testing.T) {
	stat := service.AddCaptcha("test", "admin", "badfewqref")
	fmt.Println(stat)
}

func TestVerify(t *testing.T) {
	// keyId:7807dc07-3529-4a8d-a561-e9b6ca7ee09c, captcha:RI%bc
	// 错误验证码
	msg, err := service.Verify("7807dc07-3529-4a8d-a561-e9b6ca7ee09c", "abced")
	fmt.Println(msg, err)
	// 错误keyId
	msg, err = service.Verify("7807dc07-3529-4a8d-a561-e9b6ca7ee19c", "abcedf")
	fmt.Println(msg, err)
	// 正确样例
	msg, err = service.Verify("7807dc07-3529-4a8d-a561-e9b6ca7ee09c", "abcedf")
	fmt.Println(msg, err)
	// 已使用过的验证码
	msg, err = service.Verify("7807dc07-3529-4a8d-a561-e9b6ca7ee09c", "abcedf")
	fmt.Println(msg, err)
}
