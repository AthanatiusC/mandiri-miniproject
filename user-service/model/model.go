package model

import "time"

type User struct {
	ID          int64
	Username    string
	AccessLevel int64
	Status      string
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

type UserRequest struct {
	ID          int64  `json:"id"`
	Username    string `json:"username" binding:"required"`
	AccessLevel int64  `json:"access_level" binding:"required"`
	Status      string `json:"status"`
}

type Pagination struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
}
