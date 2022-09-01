package creator

import (
	"math/rand"
	"time"
)

type OriginalCaptcha struct{}

const (
	characterBytes  = "@#$%^&*"
	digitBytes      = "0123456789"
	lowLetterBytes  = "abcdefghijklmnopqrstuvwxyz"
	highLetterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	keyQuery        = characterBytes + digitBytes + lowLetterBytes + highLetterBytes
)

func (o *OriginalCaptcha) GetString(length int) string {
	rand.Seed(time.Now().Unix())
	captcha := ""
	for i := 0; i < length; i++ {
		index := rand.Intn(len(keyQuery))
		captcha += string(keyQuery[index])
	}
	return captcha
}

func (o *OriginalCaptcha) GetDigit(length int) string {
	rand.Seed(time.Now().Unix())
	captcha := ""
	for i := 0; i < length; i++ {
		index := rand.Intn(len(keyQuery))
		captcha += string(digitBytes[index])
	}
	return captcha
}
