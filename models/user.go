package models

import "time"

type User struct {
	UserID       int64     `json:"user_id"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	Secret       string    `json:"secret"`
	PasswordSalt string    `json:"password_salt"`
	PasswordHash string    `json:"password_hash"`
	CreatedOn    time.Time `json:"created_on"`
}
