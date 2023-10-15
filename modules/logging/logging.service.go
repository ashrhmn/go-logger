package logging

import (
	"context"
	"fmt"

	"github.com/ashrhmn/go-logger/modules/storage"
	"github.com/ashrhmn/go-logger/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (ls LoggingService) InsertLog(logData types.AppLog) error {
	_, err := ls.mongoCollection.LogCollection.InsertOne(context.Background(), logData)
	if err != nil {
		return err
	}
	return nil
}

func (ls LoggingService) GetLogs(
	logLevels []string,
	from int64,
	to int64,
	limit int64,
	offset int64,
) ([]types.AppLog, int64, bool, error) {
	var logs []types.AppLog
	var count int64
	var hasMore bool
	filter := bson.D{
		{
			Key:   "level",
			Value: bson.D{{Key: "$in", Value: logLevels}},
		},
		{
			Key:   "timestamp",
			Value: bson.D{{Key: "$gte", Value: from}},
		},
		{
			Key:   "timestamp",
			Value: bson.D{{Key: "$lte", Value: to}},
		},
	}
	count, err := ls.mongoCollection.LogCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		return logs, count, hasMore, err
	}
	cursor, err := ls.mongoCollection.LogCollection.Find(
		context.Background(),
		filter,
		&options.FindOptions{
			Limit: &limit,
			Skip:  &offset,
			Sort:  bson.D{{Key: "timestamp", Value: -1}},
		},
	)
	if err != nil {
		return logs, count, hasMore, err
	}
	err = cursor.All(context.Background(), &logs)
	if err != nil {
		return logs, count, hasMore, err
	}
	hasMore = int(count) > len(logs)
	return logs, count, hasMore, nil
}
