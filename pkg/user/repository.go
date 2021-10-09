package user

import (
	"context"
	"posty/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	ReadUser(id primitive.ObjectID) (*entities.User, error)
	// UpdateUser(user *entities.User) (*entities.User, error)
	// DeleteUser(ID string) error
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreateUser(user *entities.User) (*entities.User, error) {
	user.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) ReadUser(id primitive.ObjectID) (*entities.User, error) {
	var user entities.User
	err := r.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	// fmt.Println(user)

	return &user, nil
}
