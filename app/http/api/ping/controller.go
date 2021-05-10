package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/restful"
)

type Controller struct {
}

func (c Controller) Ping(ctx *gin.Context) {
	res := service{}.Ping()

	rru := restful.NewResponse(ctx)
	rru.Retrieve(res)
}
