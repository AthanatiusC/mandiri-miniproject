package model

type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	AccessLevel int64  `json:"access_level"`
	Status      string `json:"status"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}

type UserRequest struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	AccessLevel int64  `json:"access_level"`
	Status      string `json:"status"`
}

type UserResponse struct {
	Users []User `json:"users"`
	Page  int64  `json:"page"`
	Size  int64  `json:"size"`
}
