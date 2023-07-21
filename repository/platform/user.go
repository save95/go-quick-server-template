package platform

import (
	"time"

	"server-api/global"

	"github.com/save95/go-pkg/http/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Genre       uint8
	Account     string
	Nickname    string
	AvatarURL   string
	Password    string
	State       int8
	LastLoginAt *time.Time
	LastLoginIP string
	DriverNo    string
}

func (u User) Roles() []types.IRole {
	roles := []types.IRole{
		global.RoleUser,
		global.Role(u.Genre),
	}

	return roles
}

func (u User) RoleTitles() []string {
	var titles []string
	roles := u.Roles()
	for i := range roles {
		titles = append(titles, roles[i].String())
	}
	return titles
}
