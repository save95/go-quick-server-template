package listener

import (
	"context"

	"server-api/app/listener/example"
	"server-api/global"

	"github.com/save95/go-pkg/listener"
)

type server struct {
	ctx       context.Context
	consumers []listener.IConsumer
}

func NewListenerServer(ctx context.Context) *server {
	return &server{
		ctx:       ctx,
		consumers: make([]listener.IConsumer, 0),
	}
}

func (s *server) Start() error {
	if s.consumers == nil {
		return nil
	}

	global.Log.Infof("listener server starting...")

	s.consumers = append(s.consumers, example.NewSimpleConsumer())

	for _, consumer := range s.consumers {
		consumer := consumer
		go func() {
			_ = consumer.Consume()
		}()
	}

	global.Log.Infof("listener server started")
	return nil
}

func (s *server) Shutdown() error {
	//if s.c != nil {
	//	global.Log.Infof("listener server stop")
	//	s.c.Stop()
	//}

	return nil
}
