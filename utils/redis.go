package utils

import (
	"context"
	"time"

	"github.com/iddxc/gin-plugin-captcha/global"

	"github.com/google/uuid"
)

var ctx = context.Background()

func AcquireLock(lockName string, timeout int) (stat bool, identifier string) {
	identifier = uuid.NewString()
	lockName = "lock:" + lockName
	end := time.Now().Add(time.Duration(timeout))
	for time.Now().Before(end) {
		if global.GVA_REDIS.SetNX(ctx, lockName, identifier, time.Duration(timeout)).Val() {
			return true, identifier
		} else if global.GVA_REDIS.TTL(ctx, lockName).Val() < time.Duration(timeout) {
			global.GVA_REDIS.Expire(ctx, lockName, time.Duration(timeout))
		}
		time.Sleep(time.Millisecond)
	}
	return false, ""
}

func ReleaseLock(lockName string, identifier string) (stat bool) {
	pipe := global.GVA_REDIS.Pipeline()
	lockName = "lock:" + lockName
	for {
		if pipe.Get(ctx, lockName).Val() == identifier {
			pipe.Del(ctx, lockName)
			pipe.Exec(ctx)
			return true
		} else {
			break
		}
	}
	return false
}
