package repository

import (
	"database/sql"

	"github.com/AthanatiusC/mandiri-miniproject/user-service/model"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) GetUsers(tx *sql.Tx, filter model.User) (*model.User, error) {
	return nil, nil
}

func (r *Repository) GetUserByUsername(tx *sql.Tx, username string) (*model.User, error) {
	defer tx.Commit()
	query := sq.StatementBuilder.Select("id, username, access_level, status, updated_at, created_at").From("users").Where(sq.Eq{"username": username}).
		PlaceholderFormat(sq.Dollar).
		RunWith(tx)
	var user model.User
	err := query.QueryRow().Scan(
		&user.ID,
		&user.Username,
		&user.AccessLevel,
		&user.Status,
		&user.UpdatedAt,
		&user.CreatedAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdateUser() (*model.User, error) {
	return nil, nil
}

func (r *Repository) CreateUser(tx *sql.Tx, user model.User) (*model.User, error) {
	defer tx.Commit()
	query := sq.StatementBuilder.Insert("users").Columns(
		"username",
		"access_level",
		"status",
		"updated_at",
		"created_at",
	).Values(
		user.Username,
		user.AccessLevel,
		user.Status,
		user.UpdatedAt,
		user.CreatedAt,
	).Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		RunWith(tx)
	err := query.QueryRow().Scan(&user.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &model.User{
		ID:          user.ID,
		Username:    user.Username,
		AccessLevel: user.AccessLevel,
		Status:      user.Status,
		UpdatedAt:   user.UpdatedAt,
		CreatedAt:   user.CreatedAt,
	}, nil
}

func (r *Repository) DeleteUser() (*model.User, error) {
	return nil, nil
}
