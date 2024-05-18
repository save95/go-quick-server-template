package listener

import (
	"context"
	"server-api/app/listener"

	"github.com/save95/go-pkg/listener/singleconsumer"

	"server-api/global"

	"github.com/save95/go-pkg/application"
)

type singleServer struct {
	ctx context.Context
	svr application.IApplication
}

func NewListenerServer(ctx context.Context) application.IApplication {
	return &singleServer{
		ctx: ctx,
	}
}

func (s *singleServer) Start() error {

	svr := singleconsumer.New(s.ctx)

	// 注册服务
	listener.SingleConsumerRegister(svr)

	if svr.CountConsumers() == 0 {
		global.Log.Infof("listener server no register consumer, skip")
		return nil
	}

	global.Log.Infof("listener server starting, %d consumer ...", svr.CountConsumers())

	if err := svr.Start(); nil != err {
		global.Log.Errorf("listener server start error, %s", err.Error())
		return err
	}

	global.Log.Info("listener server started")
	s.svr = svr
	return nil
}

func (s *singleServer) Shutdown() error {
	defer global.Log.Infof("listener server stoped")

	// 释放资源
	if err := listener.SingleConsumerRelease(); err != nil {
		global.Log.Errorf("single listner release failed, %+v", err)
	}

	if s.svr != nil {
		_ = s.svr.Shutdown()
	}

	return nil
}
