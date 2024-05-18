package listener

import (
	"server-api/app/listener/internal/example"

	"github.com/save95/go-pkg/listener/singleconsumer"
)

func SingleConsumerRegister(r singleconsumer.IRegister) {
	r.Register(example.RedisConsumer())
	//r.Register(example.HttpSQSConsumer())

	// todo 注册其它消费者

}

// SingleConsumerRelease 释放资源
func SingleConsumerRelease() error {

	return nil
}
