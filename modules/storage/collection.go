package storage

import "go.mongodb.org/mongo-driver/mongo"

type MongoCollection struct {
	UserCollection        *mongo.Collection
	AuthSessionCollection *mongo.Collection
	LogCollection         *mongo.Collection
}

func newMongoCollection(mongo Mongo) MongoCollection {
	userCollection := mongo.Client.Database("go-logger").Collection("users")
	authSessionCollection := mongo.Client.Database("go-logger").Collection("auth_sessions")
	logCollection := mongo.Client.Database("go-logger").Collection("logs")
	return MongoCollection{
		UserCollection:        userCollection,
		AuthSessionCollection: authSessionCollection,
		LogCollection:         logCollection,
	}
}
