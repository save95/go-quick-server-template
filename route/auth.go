package route

import (
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
}
