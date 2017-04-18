package datastore

import (
	"fmt"

	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store"
)

func GetHolidayList(where string) ([]models.Holiday, error) {
	var (
		holidays []models.Holiday
		holiday  models.Holiday
	)

	query := fmt.Sprintf("%s %s", getHolidayListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&holiday.HolidayID, &holiday.HolidayName)
		holidays = append(holidays, holiday)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return holidays, nil
}

func GetHolidayCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getHolidayCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetHoliday(holidayID int64) (*models.Holiday, error) {
	var holiday models.Holiday

	row := store.DB.QueryRow(getHolidayQuery, holidayID)
	err := row.Scan(&holiday.HolidayID, &holiday.HolidayName)
	if err != nil {
		return nil, err
	}

	return &holiday, nil
}

func CreateHoliday(holiday models.Holiday) (*models.Holiday, error) {
	var created models.Holiday

	row := store.DB.QueryRow(createHolidayQuery, holiday.HolidayName)
	err := row.Scan(&created.HolidayID, &created.HolidayName)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateHoliday(holidayID int64, holiday models.Holiday) (*models.Holiday, error) {
	var updated models.Holiday

	row := store.DB.QueryRow(updateHolidayQuery, holiday.HolidayName, holidayID)
	err := row.Scan(&updated.HolidayID, &updated.HolidayName)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteHoliday(holidayID int64) error {
	stmt, err := store.DB.Prepare(deleteHolidayQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(holidayID)
	if err != nil {
		return err
	}

	return nil
}

const getHolidayListQuery = `
SELECT *
FROM holidays
`

const getHolidayQuery = `
SELECT *
FROM holidays
WHERE holiday_id = $1
`

const createHolidayQuery = `
INSERT INTO holidays (holiday_name)
VALUES ($1)
RETURNING holiday_id, holiday_name
`

const updateHolidayQuery = `
UPDATE holidays
SET holiday_name = $1
WHERE holiday_id = $2
RETURNING holiday_id, holiday_name
`

const deleteHolidayQuery = `
DELETE
FROM holidays
WHERE holiday_id = $1
`

const getHolidayCountQuery = `
SELECT count(*)
FROM holidays
`
