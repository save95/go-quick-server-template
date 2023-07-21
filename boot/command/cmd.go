package command

import (
	"context"
	"time"

	"server-api/global"

	"github.com/save95/go-pkg/job"

	"github.com/save95/xerror"
)

type cmd struct {
	ctx     context.Context
	timeout time.Duration

	name string
	task job.ICommandJob
}

func NewCommand(ctx context.Context, name string, timeout int) *cmd {
	task, ok := tasks[name]
	if !ok {
		global.Log.Warningf("command server not register task")
	}

	return &cmd{
		ctx:     ctx,
		timeout: time.Second * time.Duration(timeout),
		name:    name,
		task:    task,
	}
}

func (c *cmd) Execute(args ...string) {
	if c.task == nil {
		global.Log.Warningf("command is nil, skip. task=%s, args=%s", c.name, args)
		return
	}

	if err := c.task.Run(args...); nil != err {
		if xe, ok := err.(xerror.XError); ok {
			err = xe.Unwrap()
		}
		global.Log.Errorf("command failed, task=%s, args=%s, err=%+v", c.name, args, err)
		return
	}

	global.Log.Infof("command done, task=%s, args=%s", c.name, args)
}
