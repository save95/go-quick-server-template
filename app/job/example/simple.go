package example

import (
	"fmt"

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

	params := global.NewCMDArgs(args...)
	version := params.Get("ver", "version")
	isTest := params.GetBool("test")
	fmt.Printf("version=%s\n", version)
	fmt.Printf("is test=%v\n", isTest)

	return nil
}
