package models

import "time"

type User struct {
	ID	        int64
	Name        string
	Email       string
	Password	[]byte
	CreatedAt   time.Time
	UpdateAt    time.Time
	DeletedAt   time.Time
}
