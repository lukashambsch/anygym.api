package datastore

import (
	"fmt"

	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store"
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
		err = rows.Scan(&userRole.UserRoleID, &userRole.UserID, &userRole.RoleID)

		if err != nil {
			return nil, err
		}

		userRole.Role, err = GetRole(userRole.RoleID)

		if err != nil {
			return nil, err
		}

		userRoles = append(userRoles, userRole)
	}
	defer rows.Close()

	return userRoles, nil
}

func GetUserRoleCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getUserRoleCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetUserRole(userRoleID int64) (*models.UserRole, error) {
	var userRole models.UserRole

	row := store.DB.QueryRow(getUserRoleQuery, userRoleID)
	err := row.Scan(&userRole.UserRoleID, &userRole.UserID, &userRole.RoleID)
	if err != nil {
		return nil, err
	}

	userRole.Role, err = GetRole(userRole.RoleID)
	if err != nil {
		return nil, err
	}

	return &userRole, nil
}

func CreateUserRole(userRole models.UserRole) (*models.UserRole, error) {
	var created models.UserRole

	row := store.DB.QueryRow(createUserRoleQuery, userRole.UserID, userRole.RoleID)
	err := row.Scan(&created.UserRoleID, &created.UserID, &created.RoleID)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateUserRole(userRoleID int64, userRole models.UserRole) (*models.UserRole, error) {
	var updated models.UserRole

	row := store.DB.QueryRow(updateUserRoleQuery, userRole.UserID, userRole.RoleID, userRoleID)
	err := row.Scan(&updated.UserRoleID, &updated.UserID, &updated.RoleID)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteUserRole(userRoleID int64) error {
	stmt, err := store.DB.Prepare(deleteUserRoleQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userRoleID)
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
