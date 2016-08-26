package datastore

import (
	"fmt"
	"time"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetVisitList(where string) ([]models.Visit, error) {
	var (
		visits []models.Visit
		visit  models.Visit
	)

	query := fmt.Sprintf("%s %s", getVisitListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&visit.VisitId,
			&visit.MemberId,
			&visit.GymLocationId,
			&visit.StatusId,
			&visit.CreatedOn,
			&visit.ModifiedOn,
		)
		visits = append(visits, visit)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return visits, nil
}

func GetVisitCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getVisitCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetVisit(visitId int64) (*models.Visit, error) {
	var visit models.Visit

	row := store.DB.QueryRow(getVisitQuery, visitId)
	err := row.Scan(
		&visit.VisitId,
		&visit.MemberId,
		&visit.GymLocationId,
		&visit.StatusId,
		&visit.CreatedOn,
		&visit.ModifiedOn,
	)

	if err != nil {
		return nil, err
	}

	return &visit, nil
}

func CreateVisit(visit models.Visit) (*models.Visit, error) {
	var created models.Visit

	row := store.DB.QueryRow(
		createVisitQuery,
		visit.MemberId,
		visit.GymLocationId,
		visit.StatusId,
	)
	err := row.Scan(
		&created.VisitId,
		&created.MemberId,
		&created.GymLocationId,
		&created.StatusId,
		&created.CreatedOn,
		&created.ModifiedOn,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateVisit(visitId int64, visit models.Visit) (*models.Visit, error) {
	var updated models.Visit

	row := store.DB.QueryRow(
		updateVisitQuery,
		visit.MemberId,
		visit.GymLocationId,
		visit.StatusId,
		time.Now().Format("01/01/2001 15:00:00"),
		visitId,
	)
	err := row.Scan(
		&updated.VisitId,
		&updated.MemberId,
		&updated.GymLocationId,
		&updated.StatusId,
		&updated.CreatedOn,
		&updated.ModifiedOn,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteVisit(visitId int64) error {
	stmt, err := store.DB.Prepare(deleteVisitQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(visitId)
	if err != nil {
		return err
	}

	return nil
}

const getVisitListQuery = `
SELECT *
FROM visits
`

const getVisitQuery = `
SELECT *
FROM visits
WHERE visit_id = $1
`

const createVisitQuery = `
INSERT INTO visits (member_id, gym_location_id, status_id)
VALUES ($1, $2, $3)
RETURNING visit_id, member_id, gym_location_id, status_id, created_on, modified_on
`

const updateVisitQuery = `
UPDATE visits
SET member_id = $1, gym_location_id = $2, status_id = $3, modified_on = $4
WHERE visit_id = $5
RETURNING visit_id, member_id, gym_location_id, status_id, created_on, modified_on
`

const deleteVisitQuery = `
DELETE
FROM visits
WHERE visit_id = $1
`

const getVisitCountQuery = `
SELECT count(*)
FROM visits
`
