package db

import (
	"context"

	"server-api/global"

	"github.com/go-redis/redis/v8"
	"github.com/save95/xerror"
)

func initRedis() error {
	if !global.Config.Redis.Enabled {
		global.Log.Debug("redis disabled, skip")
		return nil
	}

	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB,
	})
	if err := global.RedisClient.Ping(context.Background()).Err(); nil != err {
		return xerror.Wrap(err, "redis client connect failed")
	}

	global.Log.Debug("redis enabled, init success")

	return nil
}
