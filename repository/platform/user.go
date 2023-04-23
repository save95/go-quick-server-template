package platform

import (
	"time"

	"server-api/global"

	"github.com/save95/go-pkg/http/types"
	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey;autoIncrement;not null"`

	Genre       uint8  `gorm:"not null;uniqueIndex:udx_account"`
	Account     string `gorm:"not null;size:32;uniqueIndex:udx_account"`
	Nickname    string `gorm:"size:32;default:''"`
	Avatar      string `gorm:"not null;size:128;default:''"`
	Password    string `gorm:"not null;size:128"`
	State       int8   `gorm:"not null;default:0"`
	LastLoginAt *time.Time
	LastLoginIP string `gorm:"size:32"`
	DriverNo    string `gorm:"not null;size:32;default:''"`

	CreatedAt time.Time      `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"not null;default:current_timestamp on update current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted_at"`
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
