package job

import (
	"context"

	"server-api/app/job/example"
	"server-api/global"

	"github.com/save95/go-pkg/job"

	"github.com/robfig/cron/v3"
)

type server struct {
	ctx context.Context
	c   *cron.Cron
}

func NewJobServer(ctx context.Context) *server {
	return &server{
		ctx: ctx,
		c:   cron.New(),
	}
}

func (s *server) Start() error {
	if s.c == nil {
		return nil
	}

	wrapper := job.NewWrapper().WithLog(global.Log)
	global.Log.Infof("job server starting...")

	// 每10分钟，执行一次
	s.c.AddJob("*/10 * * * *", wrapper.Cron(example.NewSimpleJob()))

	s.c.Start()
	global.Log.Infof("job server started")
	return nil
}

func (s *server) Shutdown() error {
	if s.c != nil {
		global.Log.Infof("job server stop")
		s.c.Stop()
	}

	return nil
}
