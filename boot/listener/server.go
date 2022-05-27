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
	var c []listener.IConsumer
	if global.Config.Listener.Enabled {
		c = make([]listener.IConsumer, 0)
	}

	return &server{
		ctx:       ctx,
		consumers: c,
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
