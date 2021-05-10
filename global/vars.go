package global

import (
	"github.com/jinzhu/gorm"
	"github.com/save95/xlog"
)

var (
	Config projectConfig
	Log    xlog.XLogger

	// db
	DbPlatform *gorm.DB
)
