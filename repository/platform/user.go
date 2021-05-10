package platform

import (
	"server-api/global"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/save95/go-pkg/http/types"
)

type User struct {
	gorm.Model

	Genre       uint8
	Account     string
	Avatar      string
	Password    string
	State       int8
	LastLoginAt *time.Time
	LastLoginIp string
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
