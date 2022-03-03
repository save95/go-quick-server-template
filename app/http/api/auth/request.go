package auth

import (
	"server-api/global"

	"github.com/save95/xerror"
)

type createTokenRequest struct {
	Genre    global.Role `json:"genre"`
	Account  string      `json:"account"`
	Password string      `json:"password"`
}

func (c createTokenRequest) Validate() error {
	if c.Genre != global.RoleAdmin && c.Genre != global.RoleUser {
		return xerror.New("该帐号角色未开通登录权限")
	}

	if len(c.Account) == 0 || len(c.Password) == 0 {
		return xerror.New("请填写正确的登录信息")
	}

	if len(c.Password) < 6 {
		return xerror.New("密码 错误")
	}

	return nil
}

type changePwdRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func (c changePwdRequest) Validate() error {
	if len(c.OldPassword) == 0 {
		return xerror.New("原密码 不能为空")
	}
	if len(c.NewPassword) < 6 {
		return xerror.New("新密码 不能少于6个字符")
	}
	if c.NewPassword == c.OldPassword {
		return xerror.New("原密码 和 新密码 不能一致")
	}

	return nil
}
