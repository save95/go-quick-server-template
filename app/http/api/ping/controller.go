package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/restful"
)

type Controller struct {
}

// Ping 健康检查
func (c Controller) Ping(ctx *gin.Context) {
	res := service{}.Ping()

	rru := restful.NewResponse(ctx)
	rru.WithBody(res.Message)
}

// Endpoint 接受数据
func (c Controller) Endpoint(ctx *gin.Context) {
	rru := restful.NewResponse(ctx)
	rru.WithMessage("success")
}
