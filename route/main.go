package route

import (
	"server-api/app/http/api/ping"
	"server-api/route/admin"

	"github.com/gin-gonic/gin"
)

// Register 注册所有的路由
// 这里，路由请按模块分开写，一个模块一个文件（建议一个模块中分多个函数来写子模块）
// 单个模块的路由注册使用私有方法，不对外暴露
func Register(router *gin.Engine) {
	// 静态文件
	router.Static("/storage", "storage/public")

	router.Any("/ping", ping.Controller{}.Ping)

	// 注册路由
	registerAuth(router)
	admin.Register(router)
}
