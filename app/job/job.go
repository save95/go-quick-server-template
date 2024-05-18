package job

import (
	"server-api/app/job/internal/example"

	"github.com/save95/go-pkg/job"
)

// CronRegister 定时任务注册
func CronRegister(r job.ICronjobRegister) {
	// 每10分钟，执行一次
	r.Register("*/10 * * * *", example.NewSimpleJob())

	// todo 注册其它定时任务

}

// Release 释放资源
func Release() error {

	return nil
}
