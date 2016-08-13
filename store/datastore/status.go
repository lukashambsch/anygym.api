package datastore

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetStatusList() ([]models.Status, error) {
	var (
		statuses []models.Status
		status   models.Status
	)

	rows, err := store.DB.Query(getStatusListQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&status.StatusId, &status.StatusName)
		statuses = append(statuses, status)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return statuses, nil
}

func GetStatusCount() (*int, error) {
	var count int

	row := store.DB.QueryRow(getStatusCountQuery)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetStatus(statusId int64) (*models.Status, error) {
	var status models.Status

	row := store.DB.QueryRow(getStatusQuery, statusId)
	err := row.Scan(&status.StatusId, &status.StatusName)

	if err != nil {
		return nil, err
	}

	return &status, nil
}

func CreateStatus(status models.Status) (*models.Status, error) {
	var created models.Status

	row := store.DB.QueryRow(createStatusQuery, status.StatusName)
	err := row.Scan(&created.StatusId, &created.StatusName)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateStatus(statusId int64, status models.Status) (*models.Status, error) {
	var updated models.Status
	t, err := store.DB.Begin()

	if err != nil {
		return nil, err
	}

	row := t.QueryRow(updateStatusQuery, status.StatusName, statusId)
	err = row.Scan(&updated.StatusId, &updated.StatusName)

	if err != nil {
		t.Rollback()
		return nil, err
	}

	t.Commit()

	return &updated, nil
}

func DeleteStatus(statusId int64) error {
	stmt, err := store.DB.Prepare(deleteStatusQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(statusId)
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
