package service

import (
	"fmt"

	"github.com/AthanatiusC/mandiri-miniproject/user-service/model"
	"github.com/AthanatiusC/mandiri-miniproject/user-service/repository"
)

type Service struct {
	Repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) GetUsers() (*[]model.User, error) {
	return nil, nil
}

func (s *Service) UpdateUsers() (*model.User, error) {
	return nil, nil
}

func (s *Service) CreateUsers(user model.UserRequest) (*model.User, error) {
	tx, err := s.Repository.DB.Begin()
	if err != nil {
		return nil, err
	}

	existingUser, _ := s.Repository.GetUserByUsername(tx, user.Username)
	if existingUser != nil {
		return nil, fmt.Errorf("user already exist")
	}

	userModel := model.User{
		Username:    user.Username,
		AccessLevel: user.AccessLevel,
	}

	response, err := s.Repository.CreateUser(tx, userModel)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *Service) DeleteUsers() (*model.User, error) {
	return nil, nil
}
