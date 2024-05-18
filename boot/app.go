package boot

import (
	"context"
	"server-api/boot/internal/cronjob"
	"server-api/boot/internal/http"
	"server-api/boot/internal/listener"
	"server-api/boot/internal/watcher"
	"server-api/global"
	"server-api/service/lang"

	"github.com/fsnotify/fsnotify"

	"github.com/pkg/errors"

	"github.com/save95/go-pkg/application"
	"github.com/save95/xerror"
)

func initialize(cnf global.BootConfig) error {
	// 加载配置
	if err := global.ParseConfig(cnf.ConfigFilename); nil != err {
		return errors.Wrap(err, "parser config file failed")
	}

	// 初始化日志
	if err := global.InitLogger(cnf.LogCategory()); err != nil {
		return errors.Wrap(err, "init logger failed")
	}

	// 初始化db
	if err := global.InitDataBase(); err != nil {
		return errors.Wrap(err, "init db connect failed")
	}

	// 初始化语言包
	if err := lang.Init(); nil != err {
		return xerror.Wrap(err, "lang init failed")
	}

	return nil
}

func Boot(cnf global.BootConfig) error {
	if err := initialize(cnf); nil != err {
		return errors.Wrap(err, "initialize failed")
	}

	ctx := context.Background()
	// 注册 app
	app := application.NewManager(global.Log)

	// 注册 配置文件监听器
	if global.Config.App.WatchConfigEnabled {
		global.Log.Debugf("watch config file charge enabled")
		localname, _ := global.GetConfigFilename(cnf.ConfigFilename)
		app.Register(watcher.NewFileServer(ctx, localname, func(ev fsnotify.Event) error {
			// 配置文件被修改，则更新全局配置
			if ev.Op == fsnotify.Write {
				return global.ParseConfig(cnf.ConfigFilename)
			}
			return nil
		}))
	}

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
