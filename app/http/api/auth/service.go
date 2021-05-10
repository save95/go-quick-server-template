package auth

import (
	"context"
	"server-api/repository/platform/dao"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/jwt"
	"github.com/save95/go-pkg/http/types"
	"github.com/save95/go-pkg/utils/userutil"
	"github.com/save95/xerror"
)

type service struct {
}

func newService() *service {
	return &service{}
}

func (s service) Login(ctx context.Context, in *createTokenRequest) (*tokenEntity, error) {
	udao := dao.NewUser()
	user, err := udao.FirstByAccount(uint8(in.Genre), in.Account)
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

	// 生成JWT TOKEN
	claims, err := jwt.NewClaims(types.User{
		ID:    user.ID,
		Name:  user.Account,
		Roles: user.Roles(),
	})
	if err != nil {
		return nil, xerror.Wrap(err, "生成 Token 时失败")
	}

	token, err := claims.ToTokenString()
	if err != nil {
		return nil, xerror.Wrap(err, "生成 Token 时失败")
	}

	// 更新最后登陆时间
	now := time.Now()
	user.LastLoginAt = &now
	user.LastLoginIp = ctx.(*gin.Context).ClientIP()
	user.UpdatedAt = now
	if err := udao.Save(user); nil != err {
		return nil, xerror.Wrap(err, "登陆失败")
	}

	// todo 写登陆日志

	return &tokenEntity{
		AccessToken:  token,
		Roles:        user.RoleTitles(),
		Introduction: "",
		ID:           user.ID,
		Avatar:       user.Avatar,
		Name:         user.Account,
	}, nil
}
