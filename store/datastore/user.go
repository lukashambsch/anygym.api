package datastore

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

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
			&user.UserID,
			&user.Email,
			&user.Token,
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

func GetUserCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getUserCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetUser(userID int64) (*models.User, error) {
	var user models.User

	row := store.DB.QueryRow(getUserQuery, userID)
	err := row.Scan(
		&user.UserID,
		&user.Email,
		&user.Token,
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
	var err error

	user.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	row := store.DB.QueryRow(
		createUserQuery,
		user.Email,
		user.Token,
		user.PasswordHash,
	)
	err = row.Scan(
		&created.UserID,
		&created.Email,
		&created.Token,
		&created.PasswordHash,
		&created.CreatedOn,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateUser(userID int64, user models.User) (*models.User, error) {
	var updated models.User

	row := store.DB.QueryRow(
		updateUserQuery,
		user.Email,
		user.Token,
		user.PasswordHash,
		userID,
	)
	err := row.Scan(
		&updated.UserID,
		&updated.Email,
		&updated.Token,
		&updated.PasswordHash,
		&updated.CreatedOn,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteUser(userID int64) error {
	stmt, err := store.DB.Prepare(deleteUserQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID)
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
INSERT INTO users (email, token, password_hash)
VALUES ($1, $2, $3)
RETURNING user_id, email, token, password_hash, created_on
`

const updateUserQuery = `
UPDATE users
SET email = $1, token = $2, password_hash = $3
WHERE user_id = $4
RETURNING user_id, email, token, password_hash, created_on
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
