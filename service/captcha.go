package service

import (
	"context"
	"gin-plugin-captcha/global"
	"gin-plugin-captcha/service/creator"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CaptchaService struct {
	creator.LibraryCaptcha
	creator.OriginalCaptcha
	creator.ImageCapture
}

var ctx = context.Background()

func (s *CaptchaService) GetCaptcha(mode string, length, Width, Height int, FontSize float64) (keyId, data string, err error) {
	if length <= 0 || length > 6 {
		length = 6
	}
	mode = strings.ToLower(mode)
	mode = strings.ReplaceAll(mode, " ", "")
	captcha := ""
	switch mode {
	case "library":
		data = s.LibraryCaptcha.GetString()
		captcha = data
	case "original":
		data = s.OriginalCaptcha.GetString(length)
		captcha = data
	case "digit":
		data = s.OriginalCaptcha.GetDigit(length)
		captcha = data
	case "library+image":
		captcha = s.LibraryCaptcha.GetString()
		data, err = s.ImageCapture.GetImage(Width, Height, FontSize, captcha)
	case "original+image":
		captcha = s.OriginalCaptcha.GetString(length)
		data, err = s.ImageCapture.GetImage(Width, Height, FontSize, captcha)
	case "digit+image":
		captcha = s.OriginalCaptcha.GetDigit(length)
		data, err = s.ImageCapture.GetImage(Width, Height, FontSize, captcha)
	default:
		data = s.OriginalCaptcha.GetString(length)
		captcha = data
	}
	keyId = uuid.NewString()
	global.GVA_REDIS.Set(ctx, "captcha:items:"+keyId, captcha, time.Minute*10)
	return
}

func (s *CaptchaService) Verify(keyId, captcha string) (string, error) {
	val, err := global.GVA_REDIS.Get(ctx, "captcha:items:"+keyId).Result()
	captcha = strings.ToLower(captcha)
	val = strings.ToLower(val)
	if err != nil {
		return "验证码错误请重试", err
	}
	if val != captcha {
		return "验证码错误请重试", nil
	}
	global.GVA_REDIS.Del(ctx, "captcha:items:"+keyId)
	return "ok", nil
}

func (s *CaptchaService) AddCaptcha(captcha ...string) bool {
	pipe := global.GVA_REDIS.Pipeline()
	tempSlice := make([]string, 10)
	for _, v := range captcha {
		if len(v) > 6 {
			v = v[:6]
		}
		tempSlice = append(tempSlice, v)
	}
	pipe.SAdd(ctx, "captcha:library", tempSlice)
	pipe.Exec(ctx)
	return true
}
