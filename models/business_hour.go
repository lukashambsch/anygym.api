package models

import "time"

type BusinessHour struct {
	BusinessHourId int64     `json:"business_hour_id"`
	LocationId     int64     `json:"location_id"`
	HolidayId      int64     `json:"holiday_id"`
	DayId          int64     `json:"day_id"`
	OpenTime       time.Time `json:"open_time"`
	CloseTime      time.Time `json:"close_time"`
}
