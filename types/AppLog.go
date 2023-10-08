package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AppLog struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Level     string             `json:"level,omitempty" bson:"level,omitempty"`
	Note      string             `json:"note,omitempty" bson:"note,omitempty"`
	Timestamp int64              `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Context   string             `json:"context,omitempty" bson:"context,omitempty"`
	Payload   interface{}        `json:"payload,omitempty" bson:"payload,omitempty"`
}
