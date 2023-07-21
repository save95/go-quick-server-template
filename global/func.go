package global

import (
	"context"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/jwt"
	"github.com/save95/go-pkg/http/middleware"
	"github.com/save95/go-pkg/http/types"
	"github.com/save95/xerror"
)

// JWTOption JWT 相关配置
func JWTOption(refresh bool) *jwt.Option {
	opt := &jwt.Option{
		RoleConvert:     NewRole,
		RefreshDuration: 0, // 0-不自动刷新
		Secret:          []byte(Config.App.Secret),
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
		Config.Redis.Addr,
		Config.Redis.Password,
		[]byte(Config.App.Secret),
	)
	if nil != err {
		Log.Errorf("session redis store failed: %+v", err)
	}

	return store
}

// MustParseUser 从上下文中解析授权用户，否则报错
func MustParseUser(ctx context.Context) (*types.User, error) {
	htx, err := types.MustParseHttpContext(ctx)
	if nil != err {
		return nil, err
	}

	return htx.User(), nil
}

// ParseUser 从上下文中解析授权用户
func ParseUser(ctx context.Context) *types.User {
	user, err := MustParseUser(ctx)
	if nil != err {
		Log.Warningf("ParseUser: parse failed, err=%+v", err)
		return &types.User{}
	}

	return user
}

func MustParseAPPHeader(ctx context.Context) (*APPHeader, error) {
	gtx, ok := ctx.(*gin.Context)
	if !ok {
		return nil, xerror.New("parse gtx failed")
	}

	var h APPHeader
	// 优先从 http header 解析
	if err := gtx.ShouldBindHeader(&h); nil != err {
		return nil, xerror.Wrap(err, "parse global header failed from http header")
	}

	//if h == nil {
	//	// 兼容：如果从 header 获取失败，则从 query 中获取
	//	if err := gtx.ShouldBindQuery(h); nil != err {
	//		return nil, xerror.Wrap(err, "parse global header failed from query string")
	//	}
	//}

	return &h, nil
}

func ParseAPPHeader(ctx context.Context) *APPHeader {
	h, err := MustParseAPPHeader(ctx)
	if err != nil {
		Log.Warningf("ParseAppHeader: parse failed, err=%+v", err)
		return nil
	}

	return h
}
