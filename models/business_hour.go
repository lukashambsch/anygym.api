package models

import "time"

type BusinessHour struct {
	BusinessHourID int64     `json:"business_hour_id"`
	GymLocationID  int64     `json:"gym_location_id"`
	HolidayID      *int64    `json:"holiday_id"`
	DayID          *int64    `json:"day_id"`
	OpenTime       time.Time `json:"open_time"`
	CloseTime      time.Time `json:"close_time"`
}
