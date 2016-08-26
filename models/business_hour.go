package models

import "time"

type BusinessHour struct {
	BusinessHourId int64     `json:"business_hour_id"`
	GymLocationId  int64     `json:"gym_location_id"`
	HolidayId      int64     `json:"holiday_id"`
	DayId          int64     `json:"day_id"`
	OpenTime       time.Time `json:"open_time"`
	CloseTime      time.Time `json:"close_time"`
}
