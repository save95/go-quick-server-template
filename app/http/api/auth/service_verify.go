package auth

import (
	"server-api/global"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var (
	captchaSessionKey = "authCaptchaId"
)

type verifyService struct {
}

func (v verifyService) MakeCaptcha(ctx *gin.Context) (string, []byte, error) {
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

	return mimeType, bin, nil
}
