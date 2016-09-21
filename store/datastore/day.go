package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetDayList(where string) ([]models.Day, error) {
	var (
		days []models.Day
		day  models.Day
	)

	query := fmt.Sprintf("%s %s", getDayListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&day.DayID, &day.DayName)
		days = append(days, day)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return days, nil
}

func GetDayCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getDayCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetDay(dayID int64) (*models.Day, error) {
	var day models.Day

	row := store.DB.QueryRow(getDayQuery, dayID)
	err := row.Scan(&day.DayID, &day.DayName)
	if err != nil {
		return nil, err
	}

	return &day, nil
}

func CreateDay(day models.Day) (*models.Day, error) {
	var created models.Day

	row := store.DB.QueryRow(createDayQuery, day.DayName)
	err := row.Scan(&created.DayID, &created.DayName)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateDay(dayID int64, day models.Day) (*models.Day, error) {
	var updated models.Day

	row := store.DB.QueryRow(updateDayQuery, day.DayName, dayID)
	err := row.Scan(&updated.DayID, &updated.DayName)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteDay(dayID int64) error {
	stmt, err := store.DB.Prepare(deleteDayQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(dayID)
	if err != nil {
		return err
	}

	return nil
}

const getDayListQuery = `
SELECT *
FROM days
`

const getDayQuery = `
SELECT *
FROM days
WHERE day_id = $1
`

const createDayQuery = `
INSERT INTO days (day_name)
VALUES ($1)
RETURNING day_id, day_name
`

const updateDayQuery = `
UPDATE days
SET day_name = $1
WHERE day_id = $2
RETURNING day_id, day_name
`

const deleteDayQuery = `
DELETE
FROM days
WHERE day_id = $1
`

const getDayCountQuery = `
SELECT count(*)
FROM days
`
