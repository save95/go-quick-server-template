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
		middleware.JWTWith(global.JWTOption()),
		middleware.Roles([]types.IRole{global.RoleAdmin}),
	)

	registerUser(ra)
}
