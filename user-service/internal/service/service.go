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

func (s *Service) GetUsers(accessLevel int, request model.UserRequest) (*model.Response, error) {
	var response model.Response
	tx, err := s.Repository.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	users, err := s.Repository.GetUsers(tx, entity.User{
		ID:          request.ID,
		Username:    request.Username,
		AccessLevel: request.AccessLevel,
		Status:      request.Status,
	})
	if err != nil {
		return nil, err
	}

	var userResponses []model.UserResponse
	for _, user := range *users {
		if accessLevel <= int(user.AccessLevel) {
			userResponses = append(userResponses, model.UserResponse{
				ID:          user.ID,
				Username:    user.Username,
				AccessLevel: user.AccessLevel,
				Status:      user.Status,
				CreatedAt:   user.CreatedAt,
				UpdatedAt:   user.UpdatedAt,
			})
		}
	}

	return response.Construct(
		http.StatusOK,
		"success",
		userResponses,
	), nil
}

func (s *Service) UpdateUsers() (*model.Response, error) {
	return nil, nil
}

func (s *Service) CreateUsers(accessLevel int, request model.UserRequest) (*model.Response, error) {
	var response model.Response
	tx, err := s.Repository.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	if accessLevel > int(request.AccessLevel) {
		return response.Construct(
			http.StatusUnauthorized,
			"insufficent access level",
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

func (s *Service) DeleteUsers(accessLevel int, userID int) (*model.Response, error) {
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

	if accessLevel <= int(user.AccessLevel) {
		return response.Construct(
			http.StatusBadRequest,
			"insufficent access level",
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
