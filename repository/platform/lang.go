package platform

import "gorm.io/gorm"

type Lang struct {
	gorm.Model

	Code int
	ZhCN string `gorm:"column:zh_CN"`
	ZhHK string `gorm:"column:zh_HK"`
	EnUS string `gorm:"column:en_US"`
}
