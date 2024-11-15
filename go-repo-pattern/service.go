package main

import (
	"github.com/zeindevs/go-repo-pattern/mailer"
	"github.com/zeindevs/go-repo-pattern/storage"
)

type UserService struct {
	storage *storage.SqlStorage
	mailer  mailer.Client
}

func NewUserService(storage *storage.SqlStorage, mailer mailer.Client) *UserService {
	return &UserService{
		storage: storage,
		mailer:  mailer,
	}
}

func (s *UserService) GetOneByID(id uint) (*storage.User, error) {
	user, err := s.storage.GetOneByID(id)
	return user, err
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Create a new user in the database and sends a welcome email
func (s *UserService) Create(req *CreateUserRequest) error {
	user := &storage.User{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := s.storage.CreateUser(user); err != nil {
		return err
	}

	s.mailer.Send(mailer.UserWelcomeTemplates, req.Name, req.Email, req, false)

	return nil
}
