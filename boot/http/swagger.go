package http

import (
	_ "server-api/docs"
	"server-api/global"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func registerSwagger(router *gin.Engine) {
	if global.Config.Server.Swagger.Enabled {
		global.Log.Info("gin-swagger enabled")

		// 注册 swagger 路由
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
