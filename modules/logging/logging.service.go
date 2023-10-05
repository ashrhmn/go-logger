package logging

import (
	"context"
	"fmt"

	"github.com/ashrhmn/go-logger/modules/storage"
	"go.mongodb.org/mongo-driver/bson"
)

type LoggingService struct {
	mongo storage.Mongo
}

func newLoggingService(mongo storage.Mongo) LoggingService {
	return LoggingService{
		mongo: mongo,
	}
}

func (ls LoggingService) Log() {

	cursor, err := ls.mongo.Client.Database("users").Collection("test").Find(
		context.Background(),
		bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "email", Value: "hdupoy1@blogtalkradio.com"}},
			bson.D{{Key: "first_name", Value: "Cathryn"}},
		}}},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	var results []map[string]interface{}
	if err = cursor.All(context.Background(), &results); err != nil {
		fmt.Println(err)
		return
	}

	for _, result := range results {
		fmt.Println(result["email"])
	}

}
