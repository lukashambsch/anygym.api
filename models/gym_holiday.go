package models

type GymHoliday struct {
	GymHolidayId   int64 `json:"gym_holiday_id"`
	GymId          int64 `json:"gym_id"`
	HolidayId      int64 `json:"holiday_id"`
	BusinessHourId int64 `json:"business_hour_id"`
}
