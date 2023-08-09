package service

import (
	"fmt"
	"net/http"

	"github.com/AthanatiusC/mandiri-miniproject/user-service/entity"
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

func (s *Service) GetUsers() (*[]model.Response, error) {
	return nil, nil
}

func (s *Service) UpdateUsers() (*model.Response, error) {
	return nil, nil
}

func (s *Service) CreateUsers(id int64, request model.UserRequest) (*model.Response, error) {
	var response model.Response
	tx, err := s.Repository.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	admin, err := s.Repository.GetUser(tx, entity.User{ID: id})
	if err != nil {
		return nil, err
	}

	if admin.AccessLevel > 1 {
		return response.Construct(
			http.StatusUnauthorized,
			"unauthorized",
			nil,
		), nil
	}

	existingUser, _ := s.Repository.GetUser(tx, entity.User{Username: request.Username})
	if existingUser != nil {
		return response.Construct(
			http.StatusBadRequest,
			"user already exist",
			nil,
		), nil
	}

	userModel := entity.User{
		Username:    request.Username,
		AccessLevel: request.AccessLevel,
		Status:      "ACTIVE",
	}

	result, err := s.Repository.CreateUser(userModel)
	if err != nil {
		return nil, err
	}

	// Call Audit Service to save audit log

	return response.Construct(
		http.StatusBadRequest,
		"user created",
		model.UserResponse{
			ID:          result.ID,
			Username:    result.Username,
			AccessLevel: result.AccessLevel,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		},
	), nil
}

func (s *Service) DeleteUsers(userID int) (*model.Response, error) {
	var response model.Response
	tx, err := s.Repository.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	user, _ := s.Repository.GetUser(tx, entity.User{ID: int64(userID)})
	if user == nil {
		return response.Construct(
			http.StatusBadRequest,
			"user not found",
			nil,
		), nil
	}

	if userID >= 0 {
		return response.Construct(
			http.StatusBadRequest,
			"user id cannot be or less than 0",
			nil,
		), nil
	}
	result, err := s.Repository.DeleteUser(int64(userID))
	if err != nil {
		return nil, err
	}

	rows, _ := result.RowsAffected()
	return &model.Response{
		Message: fmt.Sprintf("%d user with id %d has been deleted", rows, userID),
	}, nil
}
