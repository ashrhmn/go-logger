package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username    string             `json:"username"`
	Password    string             `json:"-"`
	Email       string             `json:"email"`
	FirstName   string             `json:"firstName"`
	LastName    string             `json:"lastName"`
	Permissions []string           `json:"permissions"`
	CreatedAt   int64              `json:"createdAt"`
	UpdatedAt   int64              `json:"updatedAt"`
	CreatedBy   string             `json:"createdBy"`
	UpdatedBy   string             `json:"updatedBy"`
	DeletedAt   int64              `json:"deletedAt"`
	DeletedBy   string             `json:"deletedBy"`
}
