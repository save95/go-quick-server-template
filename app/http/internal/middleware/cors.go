package middleware

import (
	"server-api/global"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/middleware"
)

func CORS() gin.HandlerFunc {
	return middleware.CORS(
		middleware.WithCORSAllowOriginFunc(func(origin string) bool {
			if !global.Env().IsProd() {
				return true
			}

			// todo cors domain
			//return origin == "https://xxxx.com"
			return true
		}),
		middleware.WithCORSAllowHeaders("X-Custom-Key"),
		middleware.WithCORSExposeHeaders("X-Custom-Key"),
		middleware.WithCORSMaxAge(24*time.Hour),
	)
}
