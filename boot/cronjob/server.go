package cronjob

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	jobapp "server-api/app/job"
	"server-api/global"
	"server-api/repository/platform"

	"github.com/save95/xerror"

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

func (s server) Start() error {
	global.Log.Infof("cronjob server starting...")

	// 注册定时脚本
	jobapp.Register(s)

	s.c.Start()
	global.Log.Infof("cronjob server started")
	return nil
}

func (s server) Register(spec string, cmd job.ICommandJob) {
	wrapper := job.NewCronJobWrapper(
		job.WrapWithContext(s.ctx),
		job.WrapWithLogger(global.Log),
		job.WrapWithFailedSaver(s.failedSaver),
	)

	eid, err := s.c.AddJob(spec, wrapper.FromCommandJob(cmd))
	if nil != err {
		global.Log.Errorf("cronjob register failed, err=%+v", err)
		return
	}

	name := strings.Trim(fmt.Sprintf("%T", cmd), "*")
	global.Log.Infof("cronjob register success, name=%s, entryID=%d", name, int(eid))
	return
}

func (s server) failedSaver(jobName string, in []string, err error) {
	db, derr := global.Database().Get("platform")
	if nil != derr {
		global.Log.Errorf("job failed saver get db failed, err=%+v", derr)
		return
	}

	argsBytes, _ := json.Marshal(in)

	record := &platform.FailedJob{
		JobName:     jobName,
		JobArgs:     string(argsBytes),
		Payload:     xerror.ParsePayload(err),
		Errors:      xerror.FormatStackTrace(err),
		Handled:     false,
		HandledAt:   nil,
		Compensated: false,
	}
	if err := db.Create(record).Error; nil != err {
		global.Log.Errorf("job failed saver failed, err=%+v", err)
		return
	}
}

func (s server) Shutdown() error {
	if s.c != nil {
		global.Log.Infof("cronjob server stop")
		s.c.Stop()
	}

	return nil
}
