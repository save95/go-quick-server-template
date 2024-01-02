package admin

import (
	"server-api/app/http/api/admin/user"

	"github.com/gin-gonic/gin"
)

func registerUser(router gin.IRouter) {
	api := user.Controller{}
	ra := router.Group("/users")
	{
		// 列表
		ra.GET("", api.Paginate)
		// 新建
		ra.POST("", api.Create)
		// 修改
		ra.PUT("/:id", api.Modify)
	}
}
