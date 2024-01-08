package auth

import (
	"context"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"strings"
	"time"

	"github.com/save95/go-pkg/http/jwt/jwtstore"

	"github.com/save95/go-utils/otputil"

	"github.com/hashicorp/go-uuid"

	"github.com/gin-contrib/sessions"
	"github.com/mojocn/base64Captcha"

	"server-api/global"
	"server-api/global/ecode"
	"server-api/repository/platform"
	"server-api/repository/platform/dao"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/jwt"
	"github.com/save95/go-pkg/http/types"
	"github.com/save95/go-utils/userutil"
	"github.com/save95/xerror"
	"github.com/save95/xerror/xcode"
)

type service struct {
}

func newService() *service {
	return &service{}
}

func (s service) Login(ctx context.Context, in *createTokenRequest) (*tokenEntity, error) {
	if err := in.Validate(); nil != err {
		return nil, xerror.WithXCodeMessage(xcode.RequestParamError, err.Error())
	}

	// 验证码校验
	if err := s.verifyCaptcha(ctx, in.Code); nil != err {
		return nil, err
	}

	user, err := dao.NewVWUser().FirstByAccount(in.Account, "UserRoles")
	if err != nil {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorAuthParams)
	}

	// 账号无效
	if user.State != 1 {
		return nil, xerror.WithXCode(ecode.ErrorAuthParams)
	}

	// 检查密码
	if !userutil.NewHasher().Check(in.Password, user.Password) {
		return nil, xerror.WithXCode(ecode.ErrorAuthParams)
	}

	// 2FA 验证
	if err := s.verify2FA(ctx, user, in); nil != err {
		return nil, err
	}

	return s.makeToken(ctx, user)
}

func (s service) verifyCaptcha(ctx context.Context, code string) error {
	// 未开启验证码，直接返回
	if !global.Config.App.AuthCaptchaEnabled {
		return nil
	}

	// 获取验证码
	session := sessions.Default(ctx.(*gin.Context))
	captchaId := session.Get(captchaSessionKey)
	if captchaId == nil {
		return xerror.WithXCode(ecode.ErrorAuthCodeExpired)
	}

	// 校验验证码
	if !base64Captcha.VerifyCaptchaAndIsClear(captchaId.(string), code, false) {
		return xerror.WithXCode(ecode.ErrorAuthCodeInvalid)
	}

	// 清除验证码
	session.Delete(captchaSessionKey)
	_ = session.Save()
	return nil
}

func (s service) verify2FA(ctx context.Context, user *platform.VWUser, in *createTokenRequest) error {
	// 未开启 2FA 验证，直接返回
	if !global.Config.App.Auth2FAEnabled {
		return nil
	}

	id, _ := uuid.GenerateUUID()
	faTokenArgs := []string{
		in.Account,
		in.Password,
		global.Config.App.Secret,
		id,
	}
	faToken := fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(faTokenArgs, "$$$"))))

	session := sessions.Default(ctx.(*gin.Context))
	gob.Register(tfaValue{})
	session.Set(tfaTokenSessionKey, tfaValue{
		Account: in.Account,
		Token:   faToken,
	})
	_ = session.Save()

	// 2FA
	// 如果没有绑定，则返回绑定二维码地址
	tfaURL := ""
	if user.TFABindAt == nil {
		otp := otputil.NewGoogleAuth()
		// 如果已开启 2FA，但用户未设置，则自动生成 secret
		if len(user.TFASecret) == 0 {
			tfaSecret := otp.GenSecret(16)
			if err := dao.NewUser().Rest2FA(user.ID, tfaSecret); nil == err {
				user.TFASecret = tfaSecret
			}
		}
		// 已存在 secret，则生成二维码地址，并引导绑定
		if len(user.TFASecret) > 0 {
			tfaURL = otp.GetQRCodeContent(user.Account, global.Config.App.Name, user.TFASecret)
		}
	}

	// 返回特定错误码
	return xerror.WithXCode(ecode.ErrorAuthUse2FA).WithFields(map[string]interface{}{
		"token":   faToken,
		"bindUrl": tfaURL,
	})
}

func (s service) makeToken(ctx context.Context, user *platform.VWUser) (*tokenEntity, error) {

	roles, roleTitles, err := user.Roles()
	if nil != err {
		return nil, err
	}

	// token 有效时长
	duration := 1 * 24 * time.Hour
	// 多地登陆
	store := jwtstore.NewMultiRedisStore(global.SessionStoreClient)
	//// 单一登陆
	//store := jwtstore.NewSingleRedisStore(global.SessionStoreClient)
	// 生成JWT TOKEN
	token := jwt.NewStatefulToken(types.User{
		ID:      user.ID,
		Account: user.Account,
		Name:    user.Nickname,
		Roles:   roles,
	}, store).
		WithIssuer(global.Config.App.ID).
		WithSecret([]byte(global.Config.App.Secret)).
		WithDuration(duration)

	tokenStr, err := token.ToString()
	if err != nil {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorAuthFailed)
	}

	// 更新最后登陆时间
	now := time.Now()
	user.LastLoginAt = &now
	user.LastLoginIP = ctx.(*gin.Context).ClientIP()
	user.UpdatedAt = now
	if err := dao.NewUser().Save(user.ToUser()); nil != err {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorAuthFailed)
	}

	// 写登陆日志
	httpRequest := ctx.(*gin.Context).Request
	header := global.ParseAPPHeader(ctx)
	_ = dao.NewUserLoginLog().Create(&platform.UserLoginLog{
		UserID:    user.ID,
		UserAgent: httpRequest.UserAgent(),
		IP:        user.LastLoginIP,
		Referer:   httpRequest.Referer(),

		UTMSource:   header.UTMSource(),
		UTMMedium:   header.UTMMedium(),
		UTMCampaign: header.UTMCampaign(),
		UTMTerm:     header.UTMTerm(),
		UTMContent:  header.UTMContent(),
	})

	// 2FA
	tfaBind := user.TFABindAt != nil
	tfaURL := ""
	if global.Config.App.Auth2FAEnabled && !tfaBind {
		otp := otputil.NewGoogleAuth()
		// 如果已开启 2FA，但用户未设置，则自动生成 secret
		if len(user.TFASecret) == 0 {
			tfaSecret := otp.GenSecret(16)
			if err := dao.NewUser().Rest2FA(user.ID, tfaSecret); nil == err {
				user.TFASecret = tfaSecret
			}
		}
		// 已存在 secret，则生成二维码地址，并引导绑定
		if len(user.TFASecret) > 0 {
			tfaURL = otp.GetQRCodeContent(user.Account, global.Config.App.Name, user.TFASecret)
		}
	}

	return &tokenEntity{
		AccessToken: tokenStr,
		ExpireTime:  int64(duration.Seconds()),
		Profile: &profileEntity{
			ID:          user.ID,
			Name:        user.ShowName(),
			AvatarURL:   user.ShowAvatarURL(),
			CurrentRole: user.CurrentRole().String(),
			Roles:       roleTitles,
			TFABind:     tfaBind,
			TFAURL:      tfaURL,
		},
	}, nil
}

func (s service) Logout(ctx context.Context) error {
	owner := global.ParseUser(ctx)
	if owner.GetID() == 0 {
		return nil
	}

	// 清除 token
	if err := jwtstore.NewMultiRedisStore(global.SessionStoreClient).Clean(owner.GetID()); nil != err {
		return xerror.WrapWithXCode(err, ecode.ErrorHandleFailed)
	}

	return nil
}

func (s service) ChangePwd(ctx context.Context, in *changePwdRequest) error {
	if err := in.Validate(); nil != err {
		return xerror.WithXCodeMessage(xcode.RequestParamError, err.Error())
	}

	owner := global.ParseUser(ctx)
	if owner.GetID() == 0 {
		return xerror.WithXCode(xcode.Unauthorized)
	}

	user, err := dao.NewVWUser().First(owner.ID)
	if nil != err {
		return xerror.WrapWithXCode(err, ecode.ErrorRequestData)
	}

	if !userutil.NewHasher().Check(in.OldPassword, user.Password) {
		return xerror.New("原密码 错误")
	}

	password, err := userutil.NewHasher().Sum(in.NewPassword)
	if nil != err {
		return xerror.WrapWithXCode(err, ecode.ErrorHandleFailed)
	}
	user.Password = password

	if err := dao.NewUser().Save(user.ToUser()); nil != err {
		return xerror.WrapWithXCode(err, ecode.ErrorHandleFailed)
	}

	// 注销登陆
	_ = s.Logout(ctx)
	return nil
}
