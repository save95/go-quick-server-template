package middleware

import (
	"server-api/app/http/internal/helper"
	"server-api/global"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/save95/go-pkg/http/middleware"
)

func HTTPCache() gin.HandlerFunc {
	return middleware.HttpCache(
		//middleware.WithHttpCacheDebug(),
		middleware.WithHttpCacheLogger(global.Log),
		middleware.WithHttpCacheJWTOption(helper.JWTOption(false)),
		middleware.WithHttpCacheGlobalDuration(5*time.Minute),
		middleware.WithHttpCacheRedisStore(redis.NewClient(&redis.Options{
			Addr:     global.Config.HttpCache.Addr,
			Password: global.Config.HttpCache.Password,
			DB:       global.Config.HttpCache.DB,
		})),
		middleware.WithHttpCacheGlobalSkipFields("v"),
		middleware.WithHttpCacheRouteSkipFiledPolicy("/user/", true),
	)
}
