package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthSession struct {
	ID                primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Token             string             `json:"token"`
	Username          string             `json:"username"`
	Email             string             `json:"email"`
	Permissions       []string           `json:"permissions"`
	ExpiresAt         int64              `json:"expiresAt"`
	CreatedAt         int64              `json:"createdAt"`
	Client            string             `json:"client"`
	SelectedLogLevels []string           `json:"selectedLogLevels"`
}
