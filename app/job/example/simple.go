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
	global.Log.Debugf("example simple job, only print. args=%#v", args)

	params := global.NewCMDArgs(args...)
	version := params.Get("ver", "version")
	isTest := params.GetBool("test")
	fmt.Printf("example simple job args: version=%s\n", version)
	fmt.Printf("example simple job args: is test=%v\n", isTest)

	return nil
}
