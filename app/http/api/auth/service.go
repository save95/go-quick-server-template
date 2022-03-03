package auth

import (
	"context"
	"time"

	"server-api/global"
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

	user, err := dao.NewUser().FirstByAccount(uint8(in.Genre), in.Account)
	if err != nil {
		return nil, xerror.Wrap(err, "账号或密码错误")
	}

	// 账号无效
	if user.State != 1 {
		return nil, xerror.New("账号或密码错误")
	}

	// 检查密码
	if !userutil.NewHasher().Check(in.Password, user.Password) {
		return nil, xerror.New("账号或密码错误")
	}

	return s.makeToken(ctx, user)
}

func (s *service) makeToken(ctx context.Context, user *platform.User) (*tokenEntity, error) {

	// 生成JWT TOKEN
	token := jwt.NewToken(types.User{
		ID:      user.ID,
		Account: user.Account,
		Name:    user.Nickname,
		Roles:   user.Roles(),
	}).WithSecret([]byte(global.Config.App.Secret))

	tokenStr, err := token.ToString()
	if err != nil {
		return nil, xerror.Wrap(err, "生成 Token 时失败")
	}

	// 更新最后登陆时间
	now := time.Now()
	user.LastLoginAt = &now
	user.LastLoginIP = ctx.(*gin.Context).ClientIP()
	user.UpdatedAt = now
	if err := dao.NewUser().Save(user); nil != err {
		return nil, xerror.Wrap(err, "登陆失败")
	}

	// 写登陆日志
	httpRequest := ctx.(*gin.Context).Request
	_ = dao.NewUserLoginLog().Create(&platform.UserLoginLog{
		UserID:    user.ID,
		UserAgent: httpRequest.UserAgent(),
		IP:        user.LastLoginIP,
		Referer:   httpRequest.Referer(),
	})

	return &tokenEntity{
		AccessToken:  tokenStr,
		Roles:        user.RoleTitles(),
		Introduction: "",
		ID:           user.ID,
		Avatar:       user.Avatar,
		Name:         user.Account,
	}, nil
}

func (s *service) ChangePwd(ctx context.Context, in *changePwdRequest) error {
	if err := in.Validate(); nil != err {
		return xerror.WithXCodeMessage(xcode.RequestParamError, err.Error())
	}

	htx, err := types.ParserHttpContext(ctx)
	if nil != err {
		return xerror.Wrap(err, "http context parser failed")
	}

	user, err := dao.NewUser().First(htx.User().ID)
	if nil != err {
		return xerror.Wrap(err, "用户信息错误")
	}

	if !userutil.NewHasher().Check(in.OldPassword, user.Password) {
		return xerror.New("原密码 错误")
	}

	password, err := userutil.NewHasher().Sum(in.NewPassword)
	if nil != err {
		return xerror.Wrap(err, "生成密码失败")
	}
	user.Password = password

	return dao.NewUser().Save(user)
}
