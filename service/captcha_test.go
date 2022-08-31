package service

import (
	"gin-plugin-captcha/global"
	"gin-plugin-captcha/initialize"
)

func init() {
	global.GVA_CONFIG.Redis.Addr = "127.0.0.1:6379"
	global.GVA_CONFIG.Redis.DB = 0
	global.GVA_CONFIG.Redis.Password = ""
	initialize.Redis()
}

var service = new(ChatRoomService)
