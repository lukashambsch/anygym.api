package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetBusinessHourList(where string) ([]models.BusinessHour, error) {
	var (
		businessHours []models.BusinessHour
		businessHour  models.BusinessHour
	)

	query := fmt.Sprintf("%s %s", getBusinessHourListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&businessHour.BusinessHourId,
			&businessHour.GymLocationId,
			&businessHour.HolidayId,
			&businessHour.DayId,
			&businessHour.OpenTime,
			&businessHour.CloseTime,
		)
		businessHours = append(businessHours, businessHour)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return businessHours, nil
}

func GetBusinessHourCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getBusinessHourCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetBusinessHour(businessHourId int64) (*models.BusinessHour, error) {
	var businessHour models.BusinessHour

	row := store.DB.QueryRow(getBusinessHourQuery, businessHourId)
	err := row.Scan(
		&businessHour.BusinessHourId,
		&businessHour.GymLocationId,
		&businessHour.HolidayId,
		&businessHour.DayId,
		&businessHour.OpenTime,
		&businessHour.CloseTime,
	)

	if err != nil {
		return nil, err
	}

	return &businessHour, nil
}

func CreateBusinessHour(businessHour models.BusinessHour) (*models.BusinessHour, error) {
	var created models.BusinessHour

	row := store.DB.QueryRow(
		createBusinessHourQuery,
		businessHour.GymLocationId,
		businessHour.HolidayId,
		businessHour.DayId,
		businessHour.OpenTime.Format("15:00"),
		businessHour.CloseTime.Format("15:00"),
	)
	err := row.Scan(
		&created.BusinessHourId,
		&created.GymLocationId,
		&created.HolidayId,
		&created.DayId,
		&created.OpenTime,
		&created.CloseTime,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateBusinessHour(businessHourId int64, businessHour models.BusinessHour) (*models.BusinessHour, error) {
	var updated models.BusinessHour

	row := store.DB.QueryRow(
		updateBusinessHourQuery,
		businessHour.GymLocationId,
		businessHour.HolidayId,
		businessHour.DayId,
		businessHour.OpenTime.Format("15:00"),
		businessHour.CloseTime.Format("15:00"),
		businessHourId,
	)
	err := row.Scan(
		&updated.BusinessHourId,
		&updated.GymLocationId,
		&updated.HolidayId,
		&updated.DayId,
		&updated.OpenTime,
		&updated.CloseTime,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteBusinessHour(businessHourId int64) error {
	stmt, err := store.DB.Prepare(deleteBusinessHourQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(businessHourId)
	if err != nil {
		return err
	}

	return nil
}

const getBusinessHourListQuery = `
SELECT *
FROM business_hours
`

const getBusinessHourQuery = `
SELECT *
FROM business_hours
WHERE business_hour_id = $1
`

const createBusinessHourQuery = `
INSERT INTO business_hours (gym_location_id, holiday_id, day_id, open_time, close_time)
VALUES ($1, $2, $3, $4, $5)
RETURNING business_hour_id, gym_location_id, holiday_id, day_id, open_time, close_time
`

const updateBusinessHourQuery = `
UPDATE business_hours
SET gym_location_id = $1, holiday_id = $2, day_id = $3, open_time = $4, close_time = $5
WHERE business_hour_id = $6
RETURNING business_hour_id, gym_location_id, holiday_id, day_id, open_time, close_time
`

const deleteBusinessHourQuery = `
DELETE
FROM business_hours
WHERE business_hour_id = $1
`

const getBusinessHourCountQuery = `
SELECT count(*)
FROM business_hours
`
