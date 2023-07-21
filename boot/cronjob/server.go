package cronjob

import (
	"context"
	"fmt"
	"strings"

	"server-api/app/job/example"
	"server-api/global"

	"github.com/robfig/cron/v3"
	"github.com/save95/go-pkg/application"
	"github.com/save95/go-pkg/job"
)

type server struct {
	ctx context.Context
	c   *cron.Cron
}

func NewCronjobServer(ctx context.Context) application.IApplication {
	return &server{
		ctx: ctx,
		c:   cron.New(),
	}
}

func (s *server) Start() error {
	global.Log.Infof("cronjob server starting...")

	// 每10分钟，执行一次
	s.register("*/10 * * * *", example.NewSimpleJob())

	// todo 其他定时脚本

	s.c.Start()
	global.Log.Infof("cronjob server started")
	return nil
}

func (s *server) register(spec string, cmd job.ICommandJob) {
	wrapper := job.NewCronJobWrapper(
		job.WrapWithContext(s.ctx),
		job.WrapWithLogger(global.Log),
	)

	eid, err := s.c.AddJob(spec, wrapper.FromCommandJob(example.NewSimpleJob()))
	if nil != err {
		global.Log.Errorf("cronjob register failed, err=%+v", err)
		return
	}

	name := strings.Trim(fmt.Sprintf("%T", cmd), "*")
	global.Log.Infof("cronjob register success, name=%s, entryID=%d", name, int(eid))
	return
}

func (s *server) Shutdown() error {
	if s.c != nil {
		global.Log.Infof("cronjob server stop")
		s.c.Stop()
	}

	return nil
}
