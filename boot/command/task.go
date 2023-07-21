package command

import (
	"server-api/app/job/example"

	"github.com/save95/go-pkg/job"
)

var tasks = map[string]job.ICommandJob{
	"example-simple": example.NewSimpleJob(),

	// todo 其他命令脚本
}
