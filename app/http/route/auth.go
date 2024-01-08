package route

import (
	"net/http"
	"time"

	"github.com/save95/go-pkg/http/jwt/jwtstore"

	"server-api/app/http/api/auth"
	"server-api/global"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/middleware"
)

// 注册鉴权路由
func registerAuth(router *gin.Engine) {
	sOpt := middleware.SessionOption{
		Path:     "/",
		MaxAge:   10 * time.Minute,
		Secure:   false,
		HttpOnly: true,
		SameSite: 0,
	}
	// 兼容测试环境前端调试跨域 cookie 问题
	if global.Env().IsTest() {
		sOpt.Secure = true
		sOpt.SameSite = http.SameSiteNoneMode
	}

	api := auth.Controller{}

	ra := router.Group(
		"/auth",
		middleware.SessionWithStore("ac", global.SessionRedisStore(sOpt), sOpt),
	)
	{
		// 获得验证码图片/音频
		ra.GET("/captcha", api.Captcha)
		// 创建 Token
		ra.POST("/tokens", middleware.RESTFul(global.ApiVersionLatest), api.Token)
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