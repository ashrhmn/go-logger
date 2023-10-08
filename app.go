package main

import (
	"github.com/ashrhmn/go-logger/consumer"
	"github.com/ashrhmn/go-logger/modules"
	"github.com/ashrhmn/go-logger/server"
	"github.com/ashrhmn/go-logger/types"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		server.Module,
		consumer.Module,
		modules.Register,
		fx.NopLogger,
		fx.Provide(
			func() types.QueueConfig {
				return types.QueueConfig{
					QueueName: "LogQueue",
					Url:       "amqp://guest:guest@localhost:5672/",
				}
			},
		),
	).Run()
}
