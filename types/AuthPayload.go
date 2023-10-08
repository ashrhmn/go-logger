package types

import (
	"context"

	"slices"
	"time"

	"github.com/ashrhmn/go-logger/config"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthPayload struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username    string             `json:"username"`
	Email       string             `json:"email"`
	Permissions []string           `json:"permissions"`
}

func AuthPayloadFromToken(token string, authSessionCollection *mongo.Collection) (AuthPayload, error) {
	var authSession AuthSession
	err := authSessionCollection.FindOne(context.Background(),
		bson.D{{Key: "token", Value: token}},
	).Decode(&authSession)
	if err != nil {
		return AuthPayload{}, err
	}
	return AuthPayload{
		ID:          authSession.ID,
		Username:    authSession.Username,
		Email:       authSession.Email,
		Permissions: authSession.Permissions,
	}, nil
}

func (payload AuthPayload) GetValidTokens(authSessionCollection *mongo.Collection) ([]string, error) {

	cursor, err := authSessionCollection.Find(context.Background(), AuthSession{
		ID: payload.ID,
	})
	if err != nil {
		return nil, err
	}
	var results []AuthSession
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	var tokens []string
	for _, result := range results {
		if result.ExpiresAt > time.Now().Unix() {
			tokens = append(tokens, result.Token)
		}
	}

	return tokens, nil
}

func (payload AuthPayload) IsTokenValid(token string, authSessionCollection *mongo.Collection) (bool, error) {
	tokens, err := payload.GetValidTokens(authSessionCollection)
	if err != nil {
		return false, err
	}
	return slices.Contains(tokens, token), nil
}

func (payload AuthPayload) addTokenToStore(token string, authSessionCollection *mongo.Collection) error {
	authSession := AuthSession{
		ID:                primitive.NewObjectID(),
		Username:          payload.Username,
		Email:             payload.Email,
		Token:             token,
		Permissions:       payload.Permissions,
		CreatedAt:         time.Now().Unix(),
		ExpiresAt:         time.Now().Add(config.TokenExpiration).Unix(),
		Client:            "web",
		SelectedLogLevels: []string{"error"},
	}
	_, err := authSessionCollection.InsertOne(context.TODO(), authSession)
	if err != nil {
		return err
	}
	return nil
}

func (payload AuthPayload) RevokeToken(token string, authSessionCollection *mongo.Collection) error {
	_, err := authSessionCollection.DeleteOne(context.TODO(), AuthSession{
		ID:    payload.ID,
		Token: token,
	})
	if err != nil {
		return err
	}
	return nil
}

func (payload AuthPayload) RevokeAllTokens(authSessionCollection *mongo.Collection) error {
	_, err := authSessionCollection.DeleteMany(context.TODO(), AuthSession{
		ID: payload.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (payload AuthPayload) GenerateToken(authSessionCollection *mongo.Collection) (string, error) {
	token := uuid.New().String() + "." + uuid.New().String() + "." + uuid.New().String()
	err := payload.addTokenToStore(token, authSessionCollection)
	if err != nil {
		return "", err
	}
	return token, nil
}
