package auth

import (
	"net/http"

	"server-api/global"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/save95/go-pkg/constant"
	"github.com/save95/go-pkg/http/restful"
	"github.com/save95/xerror"
)

var (
	captchaSessionKey = "authCaptchaId"
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

	if len(in.Code) == 0 {
		ru.WithError(xerror.New("请填写 验证码"))
		return
	}

	if global.Config.App.AuthCaptchaEnabled {
		// 获取验证码
		session := sessions.Default(ctx)
		captchaId := session.Get(captchaSessionKey)
		if captchaId == nil {
			ru.WithError(xerror.New("验证码 已过期，请刷新验证码"))
			return
		}

		// 校验验证码
		if !base64Captcha.VerifyCaptchaAndIsClear(captchaId.(string), in.Code, false) {
			ru.WithError(xerror.New("验证码错误"))
			return
		}
		// 清除验证码
		session.Delete(captchaSessionKey)
		_ = session.Save()
	}

	token, err := newService().Login(ctx, &in)
	if err != nil {
		ru.WithError(err)
		return
	}

	ru.SetHeader(constant.HttpTokenHeaderKey, token.AccessToken)
	ru.Post(token)
}

func (c Controller) Captcha(ctx *gin.Context) {
	// 未开启授权认证的验证码功能，直接抛错
	if global.Config.App.AuthCaptchaEnabled {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// 生成验证码
	captchaConfig := global.GetCaptchaConfig()
	config := captchaConfig.ConfigCharacter
	captchaId, digitCap := base64Captcha.GenerateCaptcha("", config)

	// 通过 session 存储验证码 id
	session := sessions.Default(ctx)
	session.Set(captchaSessionKey, captchaId)
	_ = session.Save()

	bin := digitCap.BinaryEncoding()
	var mimeType string
	if _, ok := digitCap.(*base64Captcha.Audio); ok {
		mimeType = base64Captcha.MimeTypeCaptchaAudio
	} else {
		mimeType = base64Captcha.MimeTypeCaptchaImage
	}

	// 直接显示验证码
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.Data(http.StatusOK, mimeType, bin)
}

func (c Controller) ChangePwd(ctx *gin.Context) {
	ru := restful.NewResponse(ctx)

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
