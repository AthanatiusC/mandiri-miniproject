package repository

import (
	"database/sql"

	"github.com/AthanatiusC/mandiri-miniproject/user-service/model"
	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetUsers() (*model.User, error) {
	return nil, nil
}

func (r *Repository) UpdateUser() (*model.User, error) {
	return nil, nil
}

func (r *Repository) CreateUser() (*model.User, error) {
	return nil, nil
}

func (r *Repository) DeleteUser() (*model.User, error) {
	return nil, nil
}
