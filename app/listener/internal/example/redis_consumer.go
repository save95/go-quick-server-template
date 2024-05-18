package example

import (
	"server-api/global"

	"github.com/save95/go-pkg/listener/singleconsumer"

	"github.com/save95/go-pkg/queue"
)

var cnf = &queue.RedisQueueConfig{
	Addr:     global.Config.Redis.Addr,
	Password: global.Config.Redis.Password,
}

func RedisConsumer() singleconsumer.IConsumer {
	return singleconsumer.NewRedisConsumer(
		singleconsumer.WithRedisConsumerLogger(global.Log),
		singleconsumer.WithRedisConsumerHandle(cnf, "", func(val string) error {
			global.Log.Infof("[queue] simple consumer receive: %s", val)
			return nil
		}))
}
