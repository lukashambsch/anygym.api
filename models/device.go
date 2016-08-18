package models

type Device struct {
	DeviceId    int64  `json:"device_id"`
	UserId      *int64 `json:"user_id"`
	DeviceToken string `json:"device_token"`
}
