package model

import (
	"time"
)

type UserRequest struct {
	ID          int64  `json:"id"`
	Username    string `json:"username" binding:"required"`
	AccessLevel int64  `json:"access_level" binding:"required"`
	Status      string `json:"status"`
}

type UserResponse struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	AccessLevel int64     `json:"access_level"`
	Status      string    `json:"status"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type Response struct {
	Code       int         `json:"-"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page int64 `json:"page,omitempty"`
	Size int64 `json:"size,omitempty"`
}

func (Response) Construct(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Data:    data,
		Message: message,
	}
}
