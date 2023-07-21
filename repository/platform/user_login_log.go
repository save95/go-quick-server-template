package platform

import (
	"gorm.io/gorm"
)

type UserLoginLog struct {
	gorm.Model

	UserID    uint
	UserAgent string
	IP        string
	Referer   string
}
