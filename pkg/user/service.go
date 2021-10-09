package user

import (
	"posty/pkg/entities"
	"posty/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	InsertUser(user *entities.RequestUser) (*entities.User, error)
	FetchUser(id primitive.ObjectID) (*entities.User, error)
	// UpdateBook(book *entities.Book) (*entities.Book, error)
	// RemoveBook(ID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertUser(requestUser *entities.RequestUser) (*entities.User, error) {
	var user entities.User

	user.Email = requestUser.Email
	user.Password = utils.Hash(requestUser.Password)
	user.Name = requestUser.Name

	return s.repository.CreateUser(&user)
}

func (s *service) FetchUser(id primitive.ObjectID) (*entities.User, error) {
	return s.repository.ReadUser(id)
}

// func (s *service) UpdateBook(book *entities.Book) (*entities.Book, error) {
// 	return s.repository.UpdateBook(book)
// }
// func (s *service) RemoveBook(ID string) error {
// 	return s.repository.DeleteBook(ID)
// }
