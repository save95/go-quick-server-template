package example

import (
	"server-api/global"

	"github.com/save95/go-pkg/job"
)

type simpleJob struct {
}

func NewSimpleJob() job.IJob {
	return &simpleJob{}
}

func (s simpleJob) Run() error {
	global.Log.Debug("example simple job, only print")

	return nil
}
