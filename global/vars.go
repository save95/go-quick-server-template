package global

import (
	"github.com/save95/xlog"
	"gorm.io/gorm"
)

var (
	Config projectConfig
	Log    xlog.XLogger
)

var (
	DbPlatform *gorm.DB
)
