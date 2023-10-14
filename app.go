package main

import (
	"os"

	"github.com/ashrhmn/go-logger/channels"
	"github.com/ashrhmn/go-logger/consumer"
	"github.com/ashrhmn/go-logger/modules"
	"github.com/ashrhmn/go-logger/server"
	"github.com/ashrhmn/go-logger/types"
	"go.uber.org/fx"
)

func ConstructQueueConfig() types.QueueConfig {
	queueName := os.Getenv("QUEUE_NAME")
	if queueName == "" {
		queueName = "LogQueue"
	}
	rmqUrl := os.Getenv("RMQ_URL")
	if rmqUrl == "" {
		rmqUrl = "amqp://guest:guest@localhost:5672/"
	}
	return types.QueueConfig{
		QueueName: queueName,
		Url:       rmqUrl,
	}
}

func main() {
	fx.New(
		server.Module,
		consumer.Module,
		modules.Register,
		fx.NopLogger,
		fx.Provide(ConstructQueueConfig, channels.NewAppChannel),
	).Run()
}
