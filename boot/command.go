package boot

import (
	"context"

	"server-api/boot/command"
	"server-api/global"

	"github.com/pkg/errors"
)

func Command(cnf global.InitConfig) error {
	if err := initialize(cnf); nil != err {
		return errors.Wrap(err, "initialize failed")
	}

	conf := cnf.CMDConfig
	ctx := context.Background()
	command.NewCommand(ctx, conf.Name, conf.Timeout).Execute(conf.Args...)

	global.Log.Info("Command done")
	return nil
}
