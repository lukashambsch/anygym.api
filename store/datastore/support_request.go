package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetSupportRequestList(where string) ([]models.SupportRequest, error) {
	var (
		supportRequests []models.SupportRequest
		supportRequest  models.SupportRequest
	)

	query := fmt.Sprintf("%s %s", getSupportRequestListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&supportRequest.SupportRequestId,
			&supportRequest.UserId,
			&supportRequest.SupportSourceId,
			&supportRequest.Content,
			&supportRequest.Notes,
			&supportRequest.CreatedOn,
			&supportRequest.ResolvedOn,
		)
		supportRequests = append(supportRequests, supportRequest)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return supportRequests, nil
}

func GetSupportRequestCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getSupportRequestCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetSupportRequest(supportRequestId int64) (*models.SupportRequest, error) {
	var supportRequest models.SupportRequest

	row := store.DB.QueryRow(getSupportRequestQuery, supportRequestId)
	err := row.Scan(
		&supportRequest.SupportRequestId,
		&supportRequest.UserId,
		&supportRequest.SupportSourceId,
		&supportRequest.Content,
		&supportRequest.Notes,
		&supportRequest.CreatedOn,
		&supportRequest.ResolvedOn,
	)

	if err != nil {
		return nil, err
	}

	return &supportRequest, nil
}

func CreateSupportRequest(supportRequest models.SupportRequest) (*models.SupportRequest, error) {
	var created models.SupportRequest

	row := store.DB.QueryRow(
		createSupportRequestQuery,
		supportRequest.UserId,
		supportRequest.SupportSourceId,
		supportRequest.Content,
		supportRequest.Notes,
		supportRequest.CreatedOn,
		supportRequest.ResolvedOn,
	)
	err := row.Scan(
		&created.SupportRequestId,
		&created.UserId,
		&created.SupportSourceId,
		&created.Content,
		&created.Notes,
		&created.CreatedOn,
		&created.ResolvedOn,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateSupportRequest(supportRequestId int64, supportRequest models.SupportRequest) (*models.SupportRequest, error) {
	var updated models.SupportRequest

	row := store.DB.QueryRow(
		updateSupportRequestQuery,
		supportRequest.UserId,
		supportRequest.SupportSourceId,
		supportRequest.Content,
		supportRequest.Notes,
		supportRequest.CreatedOn,
		supportRequest.ResolvedOn,
		supportRequestId,
	)
	err := row.Scan(
		&updated.SupportRequestId,
		&updated.UserId,
		&updated.SupportSourceId,
		&updated.Content,
		&updated.Notes,
		&updated.CreatedOn,
		&updated.ResolvedOn,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteSupportRequest(supportRequestId int64) error {
	stmt, err := store.DB.Prepare(deleteSupportRequestQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(supportRequestId)
	if err != nil {
		return err
	}

	return nil
}

const getSupportRequestListQuery = `
SELECT *
FROM support_requests
`

const getSupportRequestQuery = `
SELECT *
FROM support_requests
WHERE support_request_id = $1
`

const createSupportRequestQuery = `
INSERT INTO support_requests (user_id, support_source_id, content, notes, created_on, resolved_on)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING support_request_id, user_id, support_source_id, content, notes, created_on, resolved_on
`

const updateSupportRequestQuery = `
UPDATE support_requests
SET user_id = $1, support_source_id = $2, content = $3, notes = $4, created_on = $5, resolved_on = $6
WHERE support_request_id = $7
RETURNING support_request_id, user_id, support_source_id, content, notes, created_on, resolved_on
`

const deleteSupportRequestQuery = `
DELETE
FROM support_requests
WHERE support_request_id = $1
`

const getSupportRequestCountQuery = `
SELECT count(*)
FROM support_requests
`
