package db

import (
	"context"
	"strings"
	"time"

	"server-api/global"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	"github.com/save95/xerror"
)

func initCache() error {
	if !global.Config.Cache.Enabled {
		global.Log.Debug("cache disabled, skip")
		return nil
	}

	// 获得不同驱动的存储
	var stored store.StoreInterface
	switch global.Config.Cache.Drive {
	case "redis":
		cnf := global.Config.Cache.Redis
		if len(cnf.Addr) == 0 || !strings.Contains(cnf.Addr, ":") {
			return xerror.New("cache redis config not exist")
		}

		stored = store.NewRedis(redis.NewClient(&redis.Options{
			Addr:     global.Config.Cache.Redis.Addr,
			Password: global.Config.Cache.Redis.Password,
			DB:       global.Config.Cache.Redis.DB,
		}), nil)
	default:
		return xerror.New("cache drive not support")
	}

	cacheManager := cache.New(stored)

	// 设置测试缓存
	if err := cacheManager.Set(context.Background(), "cacheMangerTest", "test cache", &store.Options{
		Expiration: 10 * time.Minute, // Override default value of 10 seconds defined in the store
	}); nil != err {
		return xerror.Wrap(err, "cache manager failed")
	}

	global.CacheManager = cacheManager
	global.Log.Debug("cache manger init ... success")
	return nil
}
