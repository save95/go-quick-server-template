package global

import (
	"github.com/eko/gocache/v2/cache"
	"github.com/go-redis/redis/v8"
	"github.com/save95/go-utils/locker"
	"github.com/save95/xlog"
	"gorm.io/gorm"
)

var (
	Config projectConfig
	Log    xlog.XLogger
)

var (
	DbPlatform *gorm.DB

	RedisClient *redis.Client
)

var (
	Locker       locker.ILocker
	CacheManager *cache.Cache
)
