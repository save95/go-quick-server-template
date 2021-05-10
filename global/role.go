package global

import (
	"errors"

	"github.com/save95/go-pkg/http/types"
)

type Role uint

const (
	RoleUser  Role = iota // 用户
	RoleAdmin             // 管理员
)

var rolesTitle = map[Role]string{
	RoleUser:  "user",
	RoleAdmin: "admin",
}

func (r Role) String() string {
	if title, ok := rolesTitle[r]; ok {
		return title
	}

	return "unknown"
}

func NewRole(str string) (types.IRole, error) {
	for i := range rolesTitle {
		if rolesTitle[i] == str {
			return i, nil
		}
	}

	return nil, errors.New("unknown role")
}
