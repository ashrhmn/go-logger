package consumer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ashrhmn/go-logger/channels"
	"github.com/ashrhmn/go-logger/modules/logging"
	"github.com/ashrhmn/go-logger/types"
	"github.com/gofiber/fiber/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/fx"
)

type Consumer struct {
	start func()
	stop  func()
}

func newConsumer(
	queueConfig types.QueueConfig,
	loggingService logging.LoggingService,
	appChannel channels.AppChannel,
) Consumer {
	return Consumer{
		start: func() {
			connection, err := amqp.Dial(queueConfig.Url)
			if err != nil {
				panic(err)
			}
			defer connection.Close()
			channel, err := connection.Channel()
			if err != nil {
				panic(err)
			}
			defer channel.Close()
			_, err = channel.QueueDeclare(
				queueConfig.QueueName,
				false,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				panic(err)
			}
			messages, err := channel.Consume(
				queueConfig.QueueName,
				"",
				true,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				panic(err)
			}
			for msg := range messages {
				var logData types.AppLog
				err := json.Unmarshal(msg.Body, &logData)
				if err != nil {
					log.Error("Error unmarshalling message")
					continue
				}
				logData.ID = primitive.NewObjectID()
				logData.Timestamp = time.Now().Unix()
				err = loggingService.InsertLog(logData)
				if err != nil {
					log.Error("Error inserting log")
					continue
				}
				go func() {
					appChannel.LogStream <- logData
				}()
			}
		},
		stop: func() {
			log.Info("Stopping consuming messages")
		},
	}
}

func start(lc fx.Lifecycle, c Consumer) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go c.start()
			return nil
		},
		OnStop: func(context.Context) error {
			c.stop()
			return nil
		},
	})

}

var Module = fx.Module("consumer", fx.Provide(
	newConsumer,
), fx.Invoke(
	start,
))
