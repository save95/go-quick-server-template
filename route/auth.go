package route

import (
	"server-api/app/http/api/auth"
	"server-api/global"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/middleware"
	"github.com/save95/go-pkg/http/types"
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

		// 修改密码
		ra.PUT(
			"/passwords",
			middleware.Roles([]types.IRole{global.RoleUser}),
			api.ChangePwd,
		)
	}
}
