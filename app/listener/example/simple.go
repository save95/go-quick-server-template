package example

import (
	"context"
	"time"

	"server-api/global"

	"github.com/save95/go-pkg/listener"
	"github.com/save95/go-pkg/queue"
)

type simpleConsumer struct {
}

func NewSimpleConsumer() listener.IConsumer {
	return &simpleConsumer{}
}

func (s simpleConsumer) Consume() error {
	global.Log.Info("[queue] simple consumer, start")
	defer global.Log.Info("[queue] simple consumer, end")

	ctx := context.Background()

	// 用户激活 队列
	cnf := &queue.RedisQueueConfig{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password,
	}
	queueName := "simpleQueue"
	queued := queue.NewSimpleRedis(cnf, queueName)

	for {
		str, err := queued.Pop(ctx)
		if nil != err {
			global.Log.Warningf("get queue item failed: [%s]: %+v", queueName, err)
			continue
		}

		if len(str) == 0 {
			time.Sleep(5 * time.Second)
			continue
		}

		global.Log.Infof("[queue] simple consumer receive: %s", str)
	}
}
