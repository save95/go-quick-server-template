package example

import (
	"context"
	"server-api/global"

	"github.com/save95/go-pkg/httpsqs"

	"github.com/save95/go-pkg/listener/singleconsumer"
)

func HttpSQSConsumer() singleconsumer.IConsumer {
	return singleconsumer.NewHttpSQSConsumer(
		singleconsumer.WithHttpSQSConsumerLogger(global.Log),
		singleconsumer.WithHttpSQSConsumerRetry(3),
		singleconsumer.WithHttpSQSConsumerHandler(&httpSQSConsumerHandler{}),
	)
}

type httpSQSConsumerHandler struct{}

func (h *httpSQSConsumerHandler) QueueName() string {
	return "test"
}

func (h *httpSQSConsumerHandler) GetClient() (httpsqs.IClient, error) {
	return httpsqs.NewClient(&httpsqs.Config{
		Addr:     "",
		Password: "",
		Timeout:  10,
	}), nil
}

func (h *httpSQSConsumerHandler) OnBefore(ctx context.Context) error {
	return nil
}

func (h *httpSQSConsumerHandler) Handle(ctx context.Context, data string, pos int64) error {
	return nil
}

func (h *httpSQSConsumerHandler) OnFailed(ctx context.Context, data string, err error) {
}
