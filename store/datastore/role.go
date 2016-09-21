package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetRoleList(where string) ([]models.Role, error) {
	var (
		roles []models.Role
		role  models.Role
	)

	query := fmt.Sprintf("%s %s", getRoleListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&role.RoleID, &role.RoleName)
		roles = append(roles, role)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return roles, nil
}

func GetRoleCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getRoleCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetRole(roleID int64) (*models.Role, error) {
	var role models.Role

	row := store.DB.QueryRow(getRoleQuery, roleID)
	err := row.Scan(&role.RoleID, &role.RoleName)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func CreateRole(role models.Role) (*models.Role, error) {
	var created models.Role

	row := store.DB.QueryRow(createRoleQuery, role.RoleName)
	err := row.Scan(&created.RoleID, &created.RoleName)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateRole(roleID int64, role models.Role) (*models.Role, error) {
	var updated models.Role

	row := store.DB.QueryRow(updateRoleQuery, role.RoleName, roleID)
	err := row.Scan(&updated.RoleID, &updated.RoleName)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteRole(roleID int64) error {
	stmt, err := store.DB.Prepare(deleteRoleQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(roleID)
	if err != nil {
		return err
	}

	return nil
}

const getRoleListQuery = `
SELECT *
FROM roles
`

const getRoleQuery = `
SELECT *
FROM roles
WHERE role_id = $1
`

const createRoleQuery = `
INSERT INTO roles (role_name)
VALUES ($1)
RETURNING role_id, role_name
`

const updateRoleQuery = `
UPDATE roles
SET role_name = $1
WHERE role_id = $2
RETURNING role_id, role_name
`

const deleteRoleQuery = `
DELETE
FROM roles
WHERE role_id = $1
`

const getRoleCountQuery = `
SELECT count(*)
FROM roles
`
