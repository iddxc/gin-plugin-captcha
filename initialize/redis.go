package initialize

import (
	"context"
	"fmt"

	"github.com/iddxc/gin-plugin-captcha/global"

	"github.com/go-redis/redis/v8"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Redis connect ping failed: ", err)
	} else {
		fmt.Println("Redis connect pong success: ", pong)
		global.GVA_REDIS = client
	}
}
