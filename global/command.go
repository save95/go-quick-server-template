package global

import (
	"fmt"

	"github.com/save95/go-pkg/job"
	"github.com/save95/xerror"
)

type ICronJobRegister interface {
	Register(spec string, cmd job.ICommandJob)
}

var _commandTasks = map[string]job.ICommandJob{
	//"example-simple": example.NewSimpleJob(),
}

type ICommandStore interface {
	Register(name string, cmd job.ICommandJob) error
	Get(name string) (job.ICommandJob, error)
}

type commandStore struct {
}

func CommandStore() ICommandStore {
	return &commandStore{}
}

func (cs commandStore) Register(name string, cmd job.ICommandJob) error {
	if len(name) == 0 || cmd == nil {
		return xerror.New("command task register params error")
	}

	if _, ok := _commandTasks[name]; ok {
		return xerror.New(fmt.Sprintf("%s command task duplicate registration", name))
	}

	_commandTasks[name] = cmd
	return nil
}

func (cs commandStore) Get(name string) (job.ICommandJob, error) {
	c, ok := _commandTasks[name]
	if !ok {
		return nil, xerror.New(fmt.Sprintf("%s command task not registered", name))
	}

	return c, nil
}
