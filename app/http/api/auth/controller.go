package auth

import (
	"server-api/global"
	"server-api/service/lang"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/constant"
	"github.com/save95/go-pkg/http/restful"
)

type Controller struct {
}

func (c Controller) Token(ctx *gin.Context) {
	ru := restful.NewResponse(
		ctx,
		restful.WithErrorMsgHandle(global.LangKey, lang.Handle()),
	)

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

func (c Controller) ChangePwd(ctx *gin.Context) {
	ru := restful.NewResponse(
		ctx,
		restful.WithErrorMsgHandle(global.LangKey, lang.Handle()),
	)

	var in changePwdRequest
	if err := ctx.ShouldBindJSON(&in); nil != err {
		ru.WithError(err)
		return
	}

	if err := newService().ChangePwd(ctx, &in); nil != err {
		ru.WithError(err)
		return
	}

	ru.WithMessage("success")
}
