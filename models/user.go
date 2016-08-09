package models

import "time"

type User struct {
	UserId       int64     `json:"user_id"`
	Email        string    `json:"email"`
	Token        string    `json:"-"`
	Secret       string    `json:"-"`
	Expiration   time.Time `json:"-"`
	PasswordSalt string    `json:"-"`
	PasswordHash string    `json:"-"`
	CreatedOn    time.Time `json:"created_on"`
}
