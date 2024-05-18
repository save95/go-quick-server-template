package helper

import (
	"server-api/global"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/save95/go-pkg/http/jwt"
	"github.com/save95/go-pkg/http/middleware"
)

// JWTOption JWT 相关配置
func JWTOption(refresh bool) *jwt.Option {
	opt := &jwt.Option{
		RoleConvert:     global.NewRole,
		RefreshDuration: 0, // 0-不自动刷新
		Secret:          []byte(global.Config.App.Secret),
	}

	refreshDuration := time.Duration(0)
	if refresh {
		refreshDuration = 12 * time.Hour
	}

	opt.RefreshDuration = refreshDuration

	return opt
}

// SessionRedisStore 分布式 session 存储
func SessionRedisStore(opt middleware.SessionOption) sessions.Store {
	store, err := redis.NewStore(
		int(opt.MaxAge.Minutes()), // 有效时间，分钟
		"tcp",
		global.Config.Redis.Addr,
		global.Config.Redis.Password,
		[]byte(global.Config.App.Secret),
	)
	if nil != err {
		global.Log.Errorf("session redis store failed: %+v", err)
	}

	return store
}
