package logger

import (
	"server-api/global"

	"github.com/save95/go-pkg/framework/logger"
	"github.com/save95/xlog"
)

func Init() error {
	global.Log = NewLogger()
	global.Log.Debugf("configs: %+v", global.Config)
	return nil
}

func NewLogger() xlog.XLogger {
	path := "storage/logs"
	if len(global.Config.Log.Dir) > 0 {
		path = global.Config.Log.Dir
	}

	var log xlog.XLogger
	switch global.Config.Log.Format {
	case "json":
		log = logger.NewLogger(path, global.Config.Log.Category, xlog.DailyStack, logger.WithFormat(logger.LogFormatJson))
	default:
		log = logger.NewLogger(path, global.Config.Log.Category, xlog.DailyStack)
	}

	log.SetStdPrint(global.Config.Log.StdPrint)
	if len(global.Config.Log.Level) > 0 {
		log.SetLevelByString(global.Config.Log.Level)
	}

	return log
}
