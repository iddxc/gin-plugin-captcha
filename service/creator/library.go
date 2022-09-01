package creator

import (
	"context"
	"gin-plugin-captcha/global"
)

type LibraryCaptcha struct{}

var ctx = context.Background()

func (l *LibraryCaptcha) Set(value ...string) {
	pipe := global.GVA_REDIS.Pipeline()
	pipe.SAdd(ctx, "captcha:library", value)
	pipe.Exec(ctx)
}

// func (l *LibraryCaptcha) GetString() (captcha string, err error) {
// 	length := global.GVA_REDIS.LLen(ctx, "captcha:library").Val()
// 	rand.Seed(time.Now().Unix())
// 	randNum := rand.Int63n(length)
// 	captcha, err = global.GVA_REDIS.LIndex(ctx, "captcha:library", randNum).Result()
// 	if captcha == "" || err != nil {
// 		return "", err
// 	}
// 	return captcha, nil
// }

func (l *LibraryCaptcha) GetString() (captcha string) {
	captcha, _ = global.GVA_REDIS.SRandMember(ctx, "captcha:library").Result()
	if captcha == "" {
		return ""
	}
	return captcha
}
