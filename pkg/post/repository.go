package post

import (
	"context"
	"posty/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreatePost(post *entities.Post) (*entities.Post, error)
	ReadPost(id primitive.ObjectID) (*entities.Post, error)
	ReadAllPostsbyUser(id primitive.ObjectID) (*[]entities.Post, error)
}
type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreatePost(post *entities.Post) (*entities.Post, error) {
	post.ID = primitive.NewObjectID()
	post.Timestamp = primitive.NewObjectID().Timestamp()
	_, err := r.Collection.InsertOne(context.Background(), post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *repository) ReadPost(id primitive.ObjectID) (*entities.Post, error) {
	var post entities.Post
	err := r.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&post)
	if err != nil {
		return nil, err
	}

	// fmt.Println(post)

	return &post, nil
}

func (r *repository) ReadAllPostsbyUser(id primitive.ObjectID) (*[]entities.Post, error) {
	var posts []entities.Post

	cursor, err := r.Collection.Find(context.Background(), bson.M{"userid": id})

	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var post entities.Post
		_ = cursor.Decode(&post)
		posts = append(posts, post)
	}
	return &posts, nil
}
