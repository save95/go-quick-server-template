package route

import (
	"net/http"
	"time"

	"server-api/app/http/api/auth"
	"server-api/global"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/middleware"
	"github.com/save95/go-pkg/http/types"
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
		middleware.Session("ac", "go-quick-server", sOpt),
	)
	{
		// 创建 Token
		ra.POST("/tokens", middleware.RESTFul(global.ApiVersionLatest), api.Token)
		// 获得验证码图片/音频
		ra.GET("/captcha", api.Captcha)

		// 修改密码
		ra.PUT(
			"/passwords",
			middleware.RESTFul(global.ApiVersionLatest),
			middleware.Roles([]types.IRole{global.RoleUser}),
			api.ChangePwd,
		)
	}
}
