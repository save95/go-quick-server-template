package boot

import (
	"context"

	jobapp "server-api/app/job"
	"server-api/boot/command"
	"server-api/global"

	"github.com/pkg/errors"
)

func Command(cnf global.InitConfig) error {
	if err := initialize(cnf); nil != err {
		return errors.Wrap(err, "initialize failed")
	}

	// 注册所有命令
	jobapp.RegisterCmd()

	// 执行命令
	conf := cnf.CMDConfig
	ctx := context.Background()
	command.NewCommand(ctx, conf.Timeout).Execute(conf.Name, conf.Args...)

	global.Log.Info("Command done")
	return nil
}
