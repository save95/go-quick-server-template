package platform

import (
	"time"

	"server-api/global"
	"server-api/repository/types/platformtypes"

	"github.com/save95/xerror"

	"github.com/save95/go-pkg/http/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	CheckedGenre uint8
	Account      string
	Nickname     string
	Gender       platformtypes.Gender
	AvatarURL    string
	Password     string
	State        int8
	LastLoginAt  *time.Time
	LastLoginIP  string
	DriverNo     string

	UserRoles []*UserRole `gorm:"foreignKey:UserID"`
}

func (u User) Roles() ([]types.IRole, []string, error) {
	if u.UserRoles == nil || len(u.UserRoles) == 0 {
		return nil, nil, xerror.New("user roles is empty")
	}

	roles := []types.IRole{
		global.RoleUser,
	}
	titles := []string{
		global.RoleUser.String(),
	}

	for _, role := range u.UserRoles {
		r, err := global.NewRoleFromGenre(role.Genre)
		if nil != err {
			return nil, nil, err
		}
		roles = append(roles, r)
		titles = append(titles, r.String())
	}

	return roles, titles, nil
}

func (u User) CurrentRole() types.IRole {
	r, err := global.NewRoleFromGenre(u.CheckedGenre)
	if nil == err {
		return r
	}

	return global.RoleUser
}
