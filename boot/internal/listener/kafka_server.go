package listener

import (
	"context"
	"server-api/app/listener"
	"server-api/repository/platform"

	"github.com/save95/go-pkg/application"
	"github.com/save95/go-pkg/listener/kafkaconsumer"
	"github.com/save95/xerror"

	"server-api/global"
)

type kafkaServer struct {
	ctx  context.Context
	name string
	svr  application.IApplication
}

func NewKafkaListenerServer(ctx context.Context, name string) application.IApplication {
	return &kafkaServer{
		ctx:  ctx,
		name: name,
	}
}

func (s *kafkaServer) Start() error {
	svr := kafkaconsumer.New(
		global.Config.Listener.Kafka.Addrs,
		kafkaconsumer.WithContext(s.ctx),
		kafkaconsumer.WithLogger(global.Log),
		kafkaconsumer.WithConsumeGroupFailedHandler(s.cgFailedHandler),
	)

	// 注册服务
	listener.KafkaRegister(svr)

	if svr.CountListener() == 0 {
		global.Log.Infof("listener server no register listener, skip")
		return nil
	}

	global.Log.Infof("listener server starting, %d consumer ...", svr.CountListener())

	if err := svr.Start(); nil != err {
		global.Log.Errorf("listener server start error, %s", err.Error())
		return err
	}

	// 标记在线
	flagErr := s.setRunStateFlag()

	global.Log.Infof("listener server started. flag(%v)", flagErr == nil)
	s.svr = svr
	return nil
}

func (s *kafkaServer) cgFailedHandler(consumerGroup, topic string, msg []byte, err error) {
	db, derr := global.Database().Get("platform")
	if nil != derr {
		global.Log.Errorf("job failed saver get db failed, err=%+v", derr)
		return
	}

	record := &platform.FailedListener{
		ConsumeGroup:  consumerGroup,
		Topic:         topic,
		Msg:           string(msg),
		FailedPayload: xerror.ParsePayload(err),
		FailedReason:  xerror.FormatStackTrace(err),
	}
	if err := db.Create(record).Error; nil != err {
		global.Log.Errorf("job failed saver failed, err=%+v", err)
		return
	}
}

func (s *kafkaServer) getRunningFlagKey() string {
	return global.GetServerRunningFlagKey("listener", s.name)
}

func (s *kafkaServer) setRunStateFlag() error {
	// 设置标记缓存
	key := s.getRunningFlagKey()
	if err := global.RedisClient.Set(s.ctx, key, "1", 0).Err(); nil != err {
		return xerror.Wrap(err, "listener-running state flag failed")
	}
	return nil
}

func (s *kafkaServer) Shutdown() error {
	defer global.Log.Infof("listener server stoped")

	// 释放资源
	if err := listener.KafkaRelease(); err != nil {
		global.Log.Errorf("kafka release failed, %+v", err)
	}

	if s.svr != nil {
		if err := s.svr.Shutdown(); err != nil {
			return err
		}
	}

	// 清理标记缓存
	key := s.getRunningFlagKey()
	_ = global.RedisClient.Del(s.ctx, key).Err()

	return nil
}
