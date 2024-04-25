package logger

import (
	"server-api/global"

	"github.com/save95/go-pkg/framework/logger"
	"github.com/save95/xlog"
)

func Init(category string) error {
	global.Log = NewLogger(category)
	global.Log.Debugf("configs: %+v", global.Config)
	return nil
}

func NewLogger(category string) xlog.XLogger {
	path := "storage/logs"
	if len(global.Config.Log.Dir) > 0 {
		path = global.Config.Log.Dir
	}

	var log xlog.XLogger
	switch global.Config.Log.Format {
	case "json":
		log = logger.NewLogger(path, category, xlog.DailyStack, logger.WithFormat(logger.LogFormatJson))
	default:
		log = logger.NewLogger(path, category, xlog.DailyStack)
	}

	log.SetStdPrint(global.Config.Log.StdPrint)
	if len(global.Config.Log.Level) > 0 {
		log.SetLevelByString(global.Config.Log.Level)
	}

	return log
}
