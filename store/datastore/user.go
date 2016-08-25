package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetUserList(where string) ([]models.User, error) {
	var (
		users []models.User
		user  models.User
	)

	query := fmt.Sprintf("%s %s", getUserListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&user.UserId,
			&user.Email,
			&user.Token,
			&user.Secret,
			&user.PasswordSalt,
			&user.PasswordHash,
			&user.CreatedOn,
		)
		users = append(users, user)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return users, nil
}

func GetUserCount() (*int, error) {
	var count int

	row := store.DB.QueryRow(getUserCountQuery)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetUser(userId int64) (*models.User, error) {
	var user models.User

	row := store.DB.QueryRow(getUserQuery, userId)
	err := row.Scan(
		&user.UserId,
		&user.Email,
		&user.Token,
		&user.Secret,
		&user.PasswordSalt,
		&user.PasswordHash,
		&user.CreatedOn,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user models.User) (*models.User, error) {
	var created models.User

	row := store.DB.QueryRow(
		createUserQuery,
		user.Email,
		user.Token,
		user.Secret,
		user.PasswordSalt,
		user.PasswordHash,
	)
	err := row.Scan(
		&created.UserId,
		&created.Email,
		&created.Token,
		&created.Secret,
		&created.PasswordSalt,
		&created.PasswordHash,
		&created.CreatedOn,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateUser(userId int64, user models.User) (*models.User, error) {
	var updated models.User

	row := store.DB.QueryRow(
		updateUserQuery,
		user.Email,
		user.Token,
		user.Secret,
		user.PasswordSalt,
		user.PasswordHash,
		userId,
	)
	err := row.Scan(
		&updated.UserId,
		&updated.Email,
		&updated.Token,
		&updated.Secret,
		&updated.PasswordSalt,
		&updated.PasswordHash,
		&updated.CreatedOn,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteUser(userId int64) error {
	stmt, err := store.DB.Prepare(deleteUserQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userId)
	if err != nil {
		return err
	}

	return nil
}

const getUserListQuery = `
SELECT *
FROM users
`

const getUserQuery = `
SELECT *
FROM users
WHERE user_id = $1
`

const createUserQuery = `
INSERT INTO users (email, token, secret, password_salt, password_hash)
VALUES ($1, $2, $3, $4, $5)
RETURNING user_id, email, token, secret, password_salt, password_hash, created_on
`

const updateUserQuery = `
UPDATE users
SET email = $1, token = $2, secret = $3, password_salt = $4, password_hash = $5
WHERE user_id = $6
RETURNING user_id, email, token, secret, password_salt, password_hash, created_on
`

const deleteUserQuery = `
DELETE
FROM users
WHERE user_id = $1
`

const getUserCountQuery = `
SELECT count(*)
FROM users
`
