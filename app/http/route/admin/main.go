package admin

import (
	"server-api/global"

	"github.com/save95/go-pkg/http/jwt/jwtstore"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/middleware"
)

func Register(router *gin.Engine) {
	ra := router.Group(
		"/admin",
		middleware.RESTFul(global.ApiVersionLatest),
		middleware.JWTStatefulWith(
			global.JWTOption(false),
			jwtstore.NewMultiRedisStore(global.SessionStoreClient), // 多地登录
		),
		middleware.WithRole(global.RoleAdmin),
	)

	registerUser(ra)
}
