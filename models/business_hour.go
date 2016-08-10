package models

import "time"

type BusinessHour struct {
	BusinessHourId int64     `json:"business_hour_id"`
	GymId          int64     `json:"gym_id"`
	Day            int       `json:"day"`
	OpenTime       time.Time `json:"open_time"`
	CloseTime      time.Time `json:close_time"`
}
