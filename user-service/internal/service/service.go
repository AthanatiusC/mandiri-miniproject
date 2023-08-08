package service

import "github.com/AthanatiusC/mandiri-miniproject/user-service/repository"

type User interface {
	GetUsers() (interface{}, error)
	UpdateUser() (interface{}, error)
	CreateUser() (interface{}, error)
	DeleteUser() (interface{}, error)
}

type Service struct {
	Repository repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Repository: *repository,
	}
}

func (s *Service) GetUsers() (interface{}, error) {
	return s.Repository.GetUsers()
}
