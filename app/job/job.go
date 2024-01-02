package job

import (
	"server-api/app/job/example"
	"server-api/global"
)

func Register(r global.ICronJobRegister) {
	// 每10分钟，执行一次
	r.Register("*/10 * * * *", example.NewSimpleJob())

	// todo 注册其它定时任务

}
