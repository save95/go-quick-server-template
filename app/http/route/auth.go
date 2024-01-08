package route

import (
	"github.com/save95/go-pkg/http/jwt/jwtstore"

	"server-api/app/http/api/auth"
	"server-api/global"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/middleware"
)

// 注册鉴权路由
func registerAuth(router *gin.Engine) {
	api := auth.Controller{}

	ra := router.Group(
		"/auth",
		middleware.RESTFul(global.ApiVersionLatest),
	)
	{
		// 创建 Token
		ra.POST("/tokens", api.Token)
	}

	ra2 := router.Group(
		"/auth",
		middleware.RESTFul(global.ApiVersionLatest),
		middleware.JWTStatefulWith(
			global.JWTOption(false),
			jwtstore.NewMultiRedisStore(global.SessionStoreClient), // 多地登录
		),
		middleware.WithRole(global.RoleUser),
	)
	{
		// 退出登陆
		ra2.DELETE("/tokens", api.Logout)
		// 修改密码
		ra2.PUT("/passwords", api.ChangePwd)
	}
}
