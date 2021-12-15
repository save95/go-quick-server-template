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
