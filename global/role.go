package global

import (
	"errors"

	"github.com/save95/go-pkg/http/types"
)

type Role uint8

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

func NewRoleFromGenre(genre uint8) (types.IRole, error) {
	for i := range rolesTitle {
		if uint8(i) == genre {
			return i, nil
		}
	}

	return nil, errors.New("unknown role")
}

func Role2Genre(role types.IRole) uint8 {
	switch role {
	case RoleUser:
		return uint8(RoleUser)
	case RoleAdmin:
		return uint8(RoleAdmin)
	default:
		return 0
	}
}
