package platform

import (
	"time"

	"gorm.io/gorm"
)

type AlipayUser struct {
	gorm.Model

	BindUserID uint
	BindAt     *time.Time

	AppID     string
	OpenID    string
	UnionID   string
	Nickname  string
	AvatarURL string

	Mobile string // 手机号
}
