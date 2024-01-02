package boot

import (
	"context"

	"server-api/boot/config"
	"server-api/boot/cronjob"
	"server-api/boot/db"
	"server-api/boot/http"
	"server-api/boot/listener"
	"server-api/boot/logger"
	"server-api/global"

	"github.com/pkg/errors"
	"github.com/save95/go-pkg/application"
)

func initialize(cnf global.InitConfig) error {
	// 加载配置
	if err := config.Init(cnf.ConfigFilename); nil != err {
		return errors.Wrap(err, "parser config file failed")
	}

	// 初始化日志
	if err := logger.Init(); err != nil {
		return errors.Wrap(err, "init logger failed")
	}

	// 初始化db
	if err := db.Connect(); err != nil {
		return errors.Wrap(err, "init db connect failed")
	}

	return nil
}

func Boot(cnf global.InitConfig) error {
	if err := initialize(cnf); nil != err {
		return errors.Wrap(err, "initialize failed")
	}

	ctx := context.Background()
	// 注册 app
	app := application.NewManager(global.Log)
	// 注册配置文件监听器
	app.Register(config.NewWatchServer())
	// 注册应用服务
	for _, server := range cnf.RegisterServers {
		switch server {
		case global.InitServerTypeWeb:
			app.Register(http.NewHttpServer(ctx))
		case global.InitServerTypeCronjob:
			app.Register(cronjob.NewCronjobServer(ctx))
		case global.InitServerTypeListener:
			app.Register(listener.NewListenerServer(ctx))
		}
	}

	app.Run()

	return nil
}
