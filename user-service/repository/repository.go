package repository

import (
	"database/sql"

	"github.com/AthanatiusC/mandiri-miniproject/user-service/entity"
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

func (r *Repository) GetUsers(tx *sql.Tx, filter entity.User) (*[]entity.User, error) {
	qf := CreateQueryFilter(filter)
	query := sq.StatementBuilder.Select("id, username, access_level, status, updated_at, created_at").From("users").Where(qf).
		PlaceholderFormat(sq.Dollar).
		RunWith(tx)

	rows, err := query.Query()
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(
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
		users = append(users, user)
	}
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &users, nil
}

func (r *Repository) GetUser(tx *sql.Tx, filter entity.User) (*entity.User, error) {
	qf := CreateQueryFilter(filter)
	query := sq.StatementBuilder.Select("id, username, access_level, status, updated_at, created_at").From("users").Where(qf).
		PlaceholderFormat(sq.Dollar).
		RunWith(tx)
	var user entity.User
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

func (r *Repository) UpdateUser() (*entity.User, error) {
	return nil, nil
}

func (r *Repository) CreateUser(user entity.User) (*entity.User, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	query := sq.StatementBuilder.Insert("users").Columns(
		"username",
		"access_level",
		"status",
	).Values(
		user.Username,
		user.AccessLevel,
		user.Status,
	).Suffix("RETURNING id, updated_at, created_at").
		PlaceholderFormat(sq.Dollar).
		RunWith(tx)

	err = query.QueryRow().Scan(
		&user.ID,
		&user.UpdatedAt,
		&user.CreatedAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &entity.User{
		ID:          user.ID,
		Username:    user.Username,
		AccessLevel: user.AccessLevel,
		Status:      user.Status,
		UpdatedAt:   user.UpdatedAt,
		CreatedAt:   user.CreatedAt,
	}, nil
}

func (r *Repository) DeleteUser(id int64) (sql.Result, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	result, err := sq.StatementBuilder.Delete("users").Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(tx).Exec()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return result, nil
}

func CreateQueryFilter(filter entity.User) sq.Eq {
	queryFilter := make(sq.Eq)
	if filter.ID != 0 {
		queryFilter["id"] = filter.ID
	}
	if filter.Username != "" {
		queryFilter["username"] = filter.Username
	}
	if filter.AccessLevel != 0 {
		queryFilter["access_level"] = filter.AccessLevel
	}
	if filter.Status != "" {
		queryFilter["status"] = filter.Status
	}
	return queryFilter
}
