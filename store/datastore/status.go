package datastore

import (
	"fmt"

	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store"
)

func GetStatusList(where string) ([]models.Status, error) {
	var (
		statuses []models.Status
		status   models.Status
	)

	query := fmt.Sprintf("%s %s", getStatusListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&status.StatusID, &status.StatusName)
		statuses = append(statuses, status)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return statuses, nil
}

func GetStatusCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getStatusCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetStatus(statusID int64) (*models.Status, error) {
	var status models.Status

	row := store.DB.QueryRow(getStatusQuery, statusID)
	err := row.Scan(&status.StatusID, &status.StatusName)

	if err != nil {
		return nil, err
	}

	return &status, nil
}

func CreateStatus(status models.Status) (*models.Status, error) {
	var created models.Status

	row := store.DB.QueryRow(createStatusQuery, status.StatusName)
	err := row.Scan(&created.StatusID, &created.StatusName)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateStatus(statusID int64, status models.Status) (*models.Status, error) {
	var updated models.Status

	row := store.DB.QueryRow(updateStatusQuery, status.StatusName, statusID)
	err := row.Scan(&updated.StatusID, &updated.StatusName)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteStatus(statusID int64) error {
	stmt, err := store.DB.Prepare(deleteStatusQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(statusID)
	if err != nil {
		return err
	}

	return nil
}

const getStatusListQuery = `
SELECT *
FROM statuses
`

const getStatusQuery = `
SELECT *
FROM statuses
WHERE status_id = $1
`

const createStatusQuery = `
INSERT INTO statuses (status_name)
VALUES ($1)
RETURNING status_id, status_name
`

const updateStatusQuery = `
UPDATE statuses
SET status_name = $1
WHERE status_id = $2
RETURNING status_id, status_name
`

const deleteStatusQuery = `
DELETE
FROM statuses
WHERE status_id = $1
`

const getStatusCountQuery = `
SELECT count(*)
FROM statuses
`
