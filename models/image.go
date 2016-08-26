package models

type Image struct {
	ImageId       int64  `json:"image_id"`
	GymId         *int64 `json:"gym_id"`
	GymLocationId *int64 `json:"gym_location_id"`
	UserId        *int64 `json:"user_id"`
	ImagePath     string `json:"image_path"`
}
