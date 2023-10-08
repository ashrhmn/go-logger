package logging

import (
	"context"
	"fmt"

	"github.com/ashrhmn/go-logger/modules/storage"
	"github.com/ashrhmn/go-logger/types"
	"go.mongodb.org/mongo-driver/bson"
)

type LoggingService struct {
	mongo           storage.Mongo
	mongoCollection storage.MongoCollection
}

func newLoggingService(mongo storage.Mongo, mongoCollection storage.MongoCollection) LoggingService {
	return LoggingService{
		mongo:           mongo,
		mongoCollection: mongoCollection,
	}
}

func (ls LoggingService) GetLogLevels() ([]string, error) {

	levels := []string{}

	dbLevels, err := ls.mongoCollection.LogCollection.Distinct(context.Background(), "level", bson.D{})
	if err != nil {
		fmt.Println(err)
		return levels, err
	}
	for _, level := range dbLevels {
		levelStr, ok := level.(string)
		if ok {
			levels = append(levels, levelStr)
		} else {
			fmt.Println("level is not a string", level)
		}
	}
	return levels, nil
}

func (ls LoggingService) GetSelectedLogLevel(token string) ([]string, error) {
	var authSession types.AuthSession
	err := ls.mongoCollection.AuthSessionCollection.FindOne(context.Background(), bson.D{{
		Key:   "token",
		Value: token,
	}}).Decode(&authSession)
	if err != nil {
		return []string{}, err
	}
	return authSession.SelectedLogLevels, nil
}

func (ls LoggingService) UpdateSelectedLogLevel(token string, logLevels []string) error {
	_, err := ls.mongoCollection.AuthSessionCollection.UpdateOne(
		context.Background(),
		bson.D{{Key: "token", Value: token}},
		bson.D{{Key: "$set", Value: bson.D{{Key: "selectedLogLevels", Value: logLevels}}}},
	)
	if err != nil {
		return err
	}
	return nil
}
