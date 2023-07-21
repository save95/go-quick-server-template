package example

import (
	"server-api/global"

	"github.com/save95/go-pkg/job"
)

type simpleJob struct {
}

func NewSimpleJob() job.ICommandJob {
	return &simpleJob{}
}

func (s simpleJob) Run(args ...string) error {
	global.Log.Debug("example simple job, only print. args=%s", args)

	return nil
}
