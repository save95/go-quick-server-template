package auth

import (
	"context"
	"server-api/global"
	"server-api/global/ecode"
	"server-api/repository/platform/dao"
	"time"

	"github.com/save95/go-utils/otputil"

	"github.com/save95/xerror"
	"github.com/save95/xerror/xcode"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var (
	captchaSessionKey  = "authCaptchaId"
	tfaTokenSessionKey = "2faToken"
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

func (v verifyService) Bind2FA(ctx context.Context, code string) error {
	if len(code) == 0 {
		return xerror.New("验证码错误")
	}

	u, err := global.MustParseUser(ctx)
	if nil != err {
		return err
	}

	user, err := dao.NewVWUser().First(u.GetID())
	if nil != err {
		return xerror.Wrap(err, "获取用户信息失败")
	}

	if user.TFABindAt != nil {
		return xerror.New("已绑定 2FA 令牌，无法继续绑定")
	}

	//// 如果未生成 secret，则主动生成
	//if len(user.TFASecret) == 0 {
	//	user.TFASecret = otputil.NewGoogleAuth().GenSecret(16)
	//}

	// 验证 code
	if !otputil.NewGoogleAuth().Verify(user.TFASecret, code) {
		return xerror.New("验证码错误")
	}

	now := time.Now()
	user.TFABindAt = &now

	if err := dao.NewUser().Save(user.ToUser()); nil != err {
		return xerror.Wrap(err, "绑定失败")
	}

	return nil
}

func (v verifyService) Rest2FA(ctx context.Context, code string) error {
	if len(code) == 0 {
		return xerror.New("验证码错误")
	}

	u, err := global.MustParseUser(ctx)
	if nil != err {
		return err
	}

	user, err := dao.NewVWUser().First(u.GetID())
	if nil != err {
		return xerror.Wrap(err, "获取用户信息失败")
	}

	if len(user.TFASecret) > 0 {
		return xerror.New("已绑定 2FA 令牌，无法继续绑定")
	}

	// 验证 code
	if !otputil.NewGoogleAuth().Verify(user.TFASecret, code) {
		return xerror.New("验证码错误")
	}

	secret := otputil.NewGoogleAuth().GenSecret(16)
	if err := dao.NewUser().Rest2FA(user.ID, secret); nil != err {
		return xerror.Wrap(err, "绑定失败")
	}

	return nil
}

func (v verifyService) TokenBy2FA(ctx context.Context, in *tfaRequest) (*tokenEntity, error) {
	if err := in.Validate(); nil != err {
		return nil, xerror.WithXCodeMessage(xcode.RequestParamError, err.Error())
	}

	session := sessions.Default(ctx.(*gin.Context))
	vs := session.Get(tfaTokenSessionKey)
	val := vs.(tfaValue)
	if val.Token != in.FAToken {
		return nil, xerror.WithXCode(xcode.Unauthorized)
	}

	user, err := dao.NewVWUser().FirstByAccount(val.Account, "UserRoles")
	if err != nil {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorAuthParams)
	}

	// 验证 code
	if !otputil.NewGoogleAuth().Verify(user.TFASecret, in.Code) {
		return nil, xerror.New("验证码错误")
	}

	// 清除验证码
	session.Delete(tfaTokenSessionKey)
	_ = session.Save()

	return newService().makeToken(ctx, user)
}
