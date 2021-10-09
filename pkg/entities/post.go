package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Caption   string             `json:"caption" binding:"required,min=2" bson:"caption"`
	UserID    primitive.ObjectID `json:"userid"`
	Image     string             `json:"image" binding:"required,min=2" bson:"image"`
	Timestamp time.Time          `json:"timestamp"`
}

type RequestPost struct {
	Caption  string
	Image    string
	UserID   string
	Password string
}
