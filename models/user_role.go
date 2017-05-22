package models

type UserRole struct {
	UserRoleID int64 `json:"user_role_id"`
	UserID     int64 `json:"user_id"`
	RoleID     int64 `json:"role_id"`
	Role       *Role `json:role`
}
