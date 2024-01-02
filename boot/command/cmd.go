package command

import (
	"context"
	"time"

	"github.com/save95/go-pkg/job"

	"server-api/global"

	"github.com/save95/xerror"
)

type cmd struct {
	ctx     context.Context
	timeout time.Duration

	tasks map[string]job.ICommandJob
}

func NewCommand(ctx context.Context, timeout int) *cmd {
	return &cmd{
		ctx:     ctx,
		timeout: time.Second * time.Duration(timeout),
	}
}

func (c *cmd) Execute(name string, args ...string) {
	task, err := global.CommandStore().Get(name)
	if err != nil {
		global.Log.Warningf("command server not register task, skip. task=%s, args=%s", name, args)
		return
	}
	if task == nil {
		global.Log.Warningf("command task is nil, skip. task=%s, args=%s", name, args)
		return
	}

	if err := task.Run(args...); nil != err {
		if xe, ok := err.(xerror.XError); ok {
			err = xe.Unwrap()
		}
		global.Log.Errorf("command failed, task=%s, args=%s, err=%+v", name, args, err)
		return
	}

	global.Log.Infof("command done, task=%s, args=%s", name, args)
}
