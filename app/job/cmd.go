package job

import (
	"server-api/app/job/example"
	"server-api/global"

	"github.com/save95/go-pkg/job"
)

func RegisterCmd() {
	register("example-simple", example.NewSimpleJob())

	// todo 注册其它命令

}

func register(name string, cmd job.ICommandJob) {
	cmdStore := global.CommandStore()
	if err := cmdStore.Register(name, cmd); nil != err {
		global.Log.Errorf("register command task failed: %s", err)
	}
}
