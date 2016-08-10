package models

type UserRole struct {
	UserRoleId int64 `json:"user_role_id"`
	UserId     int64 `json:"user_id"`
	RoleId     int64 `json:"role_id"`
}
