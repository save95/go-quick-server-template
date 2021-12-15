package db

import (
	"server-api/global"

	"github.com/go-redis/redis/v8"
)

func initRedis() error {
	if !global.Config.Redis.Enabled {
		global.Log.Debug("redis disabled, skip")
		return nil
	}

	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.Db,
	})
	global.Log.Debug("redis enabled, init success")

	return nil
}
