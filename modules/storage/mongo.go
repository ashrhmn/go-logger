package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

type Mongo struct {
	Client *mongo.Client
}

const uri = "mongodb://localhost:27017/test?w=majority"

func newMongo(lc fx.Lifecycle) Mongo {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {

			go func() {
				var result bson.M
				if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
					panic(err)
				}
				fmt.Println("Connected to MongoDB!")
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			if err = client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
			println("MongoDB disconnected!")
			return nil
		},
	})

	return Mongo{
		Client: client,
	}
}
