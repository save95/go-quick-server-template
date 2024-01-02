package platform

import (
	"time"

	"gorm.io/gorm"
)

type WechatUser struct {
	gorm.Model

	UserID uint
	BindAt *time.Time

	AppID      string
	OpenID     string
	UnionID    string
	Nickname   string
	Gender     uint
	Province   string
	City       string
	Country    string
	HeadImgURL string
	Privilege  string

	PhoneNumber      string
	PurePhoneNumber  string
	PhoneCountryCode string
}
