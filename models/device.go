package models

type Device struct {
	DeviceID    int64  `json:"device_id"`
	UserID      int64  `json:"user_id"`
	DeviceToken string `json:"device_token"`
}
