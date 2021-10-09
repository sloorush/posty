package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name     string             `json:"name" binding:"required,min=2" bson:"name"`
	Email    string             `json:"email" binding:"required" bson:"email"`
	Password string             `json:"password" binding:"required" bson:"password"`
}

type RequestUser struct {
	Email    string
	Name     string
	Password string
}
