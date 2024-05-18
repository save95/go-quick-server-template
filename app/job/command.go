package job

import (
	"server-api/app/job/internal/example"

	"github.com/save95/go-pkg/job"
)

func CMDRegister(r job.ICommandRegister) {
	r.Register("example-simple", example.NewSimpleJob())

	// todo 注册其它命令

}
