package api

import (
	"captcha/model"
	"captcha/model/response"
	"captcha/service"

	"strconv"

	"github.com/gin-gonic/gin"
)

type CaptchaApi struct{}

func (a *CaptchaApi) Get(c *gin.Context) {
	length := c.DefaultQuery("length", "6")
	l, err := strconv.Atoi(length)
	if err != nil {
		response.FailWithMessage("length 为非发参数", c)
		return
	}
	mode := c.DefaultQuery("mode", "original")
	width := c.DefaultQuery("width", "200")
	w, err := strconv.Atoi(width)
	if err != nil {
		response.FailWithMessage("width 为非发参数", c)
		return
	}

	height := c.DefaultQuery("height", "80")
	h, err := strconv.Atoi(height)
	if err != nil {
		response.FailWithMessage("height 为非发参数", c)
		return
	}

	fontSize := c.DefaultQuery("fontSize", "24")
	fs, err := strconv.Atoi(fontSize)
	if err != nil {
		response.FailWithMessage("fontSize 为非发参数", c)
		return
	}
	keyId, data, err := service.ServiceGroupApp.CaptchaService.GetCaptcha(mode, l, w, h, float64(fs))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cap := model.Captcha{
		KeyId: keyId,
		Src:   data,
	}
	response.OkWithData(cap, c)
}

func (a *CaptchaApi) Verify(c *gin.Context) {
	var cap model.CaptchaReq
	err := c.BindJSON(&cap)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	msg, err := service.ServiceGroupApp.CaptchaService.Verify(cap.KeyId, cap.Captcha)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if msg != "ok" {
		response.FailWithMessage(msg, c)
		return
	}
	response.Ok(c)
}

func (a *CaptchaApi) AddCaptcha(c *gin.Context) {
	var capList model.CaptchaListReq
	err := c.BindJSON(&capList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	service.ServiceGroupApp.CaptchaService.AddCaptcha(capList.CaptchaList...)
	response.Ok(c)
}
