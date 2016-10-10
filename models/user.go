package models

import "time"

type User struct {
	UserID       int64     `json:"user_id"`
	Email        string    `json:"email"`
	Token        string    `json:"-"`
	PasswordHash []byte    `json:"-"`
	Password     string    `json:"password"`
	CreatedOn    time.Time `json:"created_on"`
}
