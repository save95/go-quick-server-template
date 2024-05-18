package logger

import (
	"github.com/save95/go-pkg/framework/logger"
	"github.com/save95/xlog"
)

const defaultLogPath = "storage/logs"

func New(cnf LogConfig, category string) xlog.XLogger {
	path := defaultLogPath
	if len(cnf.Dir) > 0 {
		path = cnf.Dir
	}

	var log xlog.XLogger
	switch cnf.Format() {
	case "json":
		log = logger.NewLogger(path, category, xlog.DailyStack, logger.WithFormat(logger.LogFormatJson))
	default:
		log = logger.NewLogger(path, category, xlog.DailyStack)
	}

	log.SetStdPrint(cnf.StdPrint)
	if len(cnf.Level) > 0 {
		log.SetLevelByString(cnf.Level)
	}

	return log
}
