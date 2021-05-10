package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/constant"
	"github.com/save95/go-pkg/http/restful"
)

type Controller struct {
}

func (c Controller) Token(ctx *gin.Context) {
	ru := restful.NewResponse(ctx)

	var in createTokenRequest
	if err := ctx.ShouldBindJSON(&in); nil != err {
		ru.WithError(err)
		return
	}

	token, err := newService().Login(ctx, &in)
	if err != nil {
		ru.WithError(err)
		return
	}

	ru.SetHeader(constant.HttpTokenHeaderKey, token.AccessToken)
	ru.Post(token)
}
