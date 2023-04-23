package global

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/save95/go-pkg/http/middleware"
)

func CORSConfig() cors.Config {
	return cors.Config{
		AllowOriginFunc: func(origin string) bool {
			if !Env().IsProd() {
				return true
			}

			// todo
			//return origin == "https://xxxx.com"
			return true
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "User-Agent", "Cookie", "Authorization",
			"X-Auth-Token", "X-Token", "X-Requested-With",
			// https://www.npmjs.com/package/huge-uploader
			"uploader-chunk-number", "uploader-chunks-total", "uploader-file-id",
		},
		AllowCredentials: true,
		ExposeHeaders: []string{
			"Authorization", "Content-MD5",
			// 分页响应头
			"Link", "X-More-Resource", "X-Pagination-Info", "X-Total-Count", "X-Error-Code",
		},
		MaxAge: 12 * time.Hour,
	}
}

func JWTOption(refresh bool) *middleware.JWTOption {
	opt := &middleware.JWTOption{
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
