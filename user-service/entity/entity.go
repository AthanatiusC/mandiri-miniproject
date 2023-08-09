package entity

import "time"

type User struct {
	ID          int64
	Username    string
	AccessLevel int64
	Status      string
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
