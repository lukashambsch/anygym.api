package models

type Image struct {
	ImageID       int64  `json:"image_id"`
	GymID         *int64 `json:"gym_id"`
	GymLocationID *int64 `json:"gym_location_id"`
	UserID        *int64 `json:"user_id"`
	ImagePath     string `json:"image_path"`
}
