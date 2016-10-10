package models

import "time"

type User struct {
	UserID       int64     `json:"user_id"`
	Email        string    `json:"email"`
	Token        string    `json:"-"`
	PasswordHash []byte    `json:"-"`
	CreatedOn    time.Time `json:"created_on"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
