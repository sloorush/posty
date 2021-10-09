package post

import (
	"posty/pkg/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	InsertPost(post *entities.Post) (*entities.Post, error)
	FetchPost(id primitive.ObjectID) (*entities.Post, error)
	FetchAllPostsByUser(id primitive.ObjectID) (*[]entities.Post, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertPost(post *entities.Post) (*entities.Post, error) {
	return s.repository.CreatePost(post)
}

func (s *service) FetchPost(id primitive.ObjectID) (*entities.Post, error) {
	return s.repository.ReadPost(id)
}

func (s *service) FetchAllPostsByUser(id primitive.ObjectID) (*[]entities.Post, error) {
	return s.repository.ReadAllPostsbyUser(id)
}
