package auth

import (
	"net/http"
	"server-api/global"
	"server-api/global/ecode"
	"server-api/service/lang"

	"github.com/save95/xerror"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/constant"
	"github.com/save95/go-pkg/http/restful"
)

type Controller struct {
}

func (c Controller) Captcha(ctx *gin.Context) {
	// 未开启授权认证的验证码功能，直接抛错
	if !global.Config.App.AuthCaptchaEnabled {
		global.Log.Error("captcha bad request, because config.app.AuthCaptchaEnabled=false")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	mimeType, bin, err := verifyService{}.MakeCaptcha(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// 直接显示验证码
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.Data(http.StatusOK, mimeType, bin)
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
		if xerror.IsXCode(err, ecode.ErrorAuthUse2FA) {
			ru.WithErrorData(err, xerror.ParsePayload(err))
			return
		}
		ru.WithError(err)
		return
	}

	ru.SetHeader(constant.HttpTokenHeaderKey, token.AccessToken)
	ru.Post(token)
}

func (c Controller) Logout(ctx *gin.Context) {
	err := newService().Logout(ctx)

	ru := restful.NewResponse(
		ctx,
		restful.WithErrorMsgHandle(global.LangKey, lang.Handle()),
	)
	ru.Delete(err)
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
