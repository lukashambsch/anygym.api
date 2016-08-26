package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetUserRoleList(where string) ([]models.UserRole, error) {
	var (
		userRoles []models.UserRole
		userRole  models.UserRole
	)

	query := fmt.Sprintf("%s %s", getUserRoleListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&userRole.UserRoleId, &userRole.UserId, &userRole.RoleId)
		userRoles = append(userRoles, userRole)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return userRoles, nil
}

func GetUserRoleCount() (*int, error) {
	var count int

	row := store.DB.QueryRow(getUserRoleCountQuery)
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetUserRole(userRoleId int64) (*models.UserRole, error) {
	var userRole models.UserRole

	row := store.DB.QueryRow(getUserRoleQuery, userRoleId)
	err := row.Scan(&userRole.UserRoleId, &userRole.UserId, &userRole.RoleId)
	if err != nil {
		return nil, err
	}

	return &userRole, nil
}

func CreateUserRole(userRole models.UserRole) (*models.UserRole, error) {
	var created models.UserRole

	row := store.DB.QueryRow(createUserRoleQuery, userRole.UserId, userRole.RoleId)
	err := row.Scan(&created.UserRoleId, &created.UserId, &created.RoleId)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateUserRole(userRoleId int64, userRole models.UserRole) (*models.UserRole, error) {
	var updated models.UserRole

	row := store.DB.QueryRow(updateUserRoleQuery, userRole.UserId, userRole.RoleId, userRoleId)
	err := row.Scan(&updated.UserRoleId, &updated.UserId, &updated.RoleId)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteUserRole(userRoleId int64) error {
	stmt, err := store.DB.Prepare(deleteUserRoleQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userRoleId)
	if err != nil {
		return err
	}

	return nil
}

const getUserRoleListQuery = `
SELECT *
FROM user_roles
`

const getUserRoleQuery = `
SELECT *
FROM user_roles
WHERE user_role_id = $1
`

const createUserRoleQuery = `
INSERT INTO user_roles (user_id, role_id)
VALUES ($1, $2)
RETURNING user_role_id, user_id, role_id
`

const updateUserRoleQuery = `
UPDATE user_roles
SET user_id = $1, role_id = $2
WHERE user_role_id = $3
RETURNING user_role_id, user_id, role_id
`

const deleteUserRoleQuery = `
DELETE
FROM user_roles
WHERE user_role_id = $1
`

const getUserRoleCountQuery = `
SELECT count(*)
FROM user_roles
`
