package admin

import (
	"server-api/global"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/middleware"
	"github.com/save95/go-pkg/http/types"
)

func Register(router *gin.Engine) {
	ra := router.Group(
		"/admin",
		middleware.RESTFul(global.ApiVersionLatest),
		middleware.JWTWith(&middleware.JWTOption{
			RoleConvert:     global.NewRole,
			RefreshDuration: 0, // 0-不自动刷新
			Secret:          []byte("go-quick-server-template"),
		}),
		middleware.Roles([]types.IRole{global.RoleAdmin}),
	)

	registerUser(ra)
}
