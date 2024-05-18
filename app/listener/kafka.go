package listener

import (
	"server-api/app/listener/internal/example"
	"server-api/global/kafka/cg"
	"server-api/global/kafka/topic"

	"github.com/save95/go-pkg/listener/kafkaconsumer"
)

func KafkaRegister(s kafkaconsumer.IRegister) {
	s.Register(cg.ExampleRecorder, example.KafkaConsumer, topic.ExampleData)

	// 注册其他消费者

}

// KafkaRelease 释放资源
func KafkaRelease() error {

	return nil
}
