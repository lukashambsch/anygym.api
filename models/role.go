package models

type Role struct {
	RoleId   int64  `json:"role_id"`
	RoleName string `json:"role_name"`
}

type Roles []Role
