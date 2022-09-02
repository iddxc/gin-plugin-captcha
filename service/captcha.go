package service

import (
	"context"
	"strings"
	"time"

	"github.com/iddxc/gin-plugin-captcha/global"
	"github.com/iddxc/gin-plugin-captcha/service/creator"

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
		captcha = s.LibraryCaptcha.GetString()
	case "original":
		captcha = s.OriginalCaptcha.GetString(length)
	case "digit":
		captcha = s.OriginalCaptcha.GetDigit(length)
	case "library_image":
		captcha = s.LibraryCaptcha.GetString()
		data, err = s.ImageCapture.GetImage(Width, Height, FontSize, captcha)
	case "original_image":
		captcha = s.OriginalCaptcha.GetString(length)
		data, err = s.ImageCapture.GetImage(Width, Height, FontSize, captcha)
	case "digit_image":
		captcha = s.OriginalCaptcha.GetDigit(length)
		data, err = s.ImageCapture.GetImage(Width, Height, FontSize, captcha)
	default:
		captcha = s.OriginalCaptcha.GetString(length)
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
		r := []rune(v)
		if len(r) > 6 {
			v = string(r[:6])
		}
		tempSlice = append(tempSlice, v)
	}
	pipe.SAdd(ctx, "captcha:library", tempSlice)
	pipe.Exec(ctx)
	return true
}
