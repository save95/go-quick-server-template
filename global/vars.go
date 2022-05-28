package global

import (
	"github.com/eko/gocache/v2/cache"
	"github.com/go-redis/redis/v8"
	"github.com/save95/go-utils/locker"
	"github.com/save95/xlog"
)

var version = "v1.1.0"

var (
	Config projectConfig
	Log    xlog.XLogger
)

var (
	RedisClient *redis.Client
)

var (
	Locker       locker.ILocker
	CacheManager *cache.Cache
)
