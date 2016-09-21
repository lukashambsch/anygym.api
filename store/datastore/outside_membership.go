package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetOutsideMembershipList(where string) ([]models.OutsideMembership, error) {
	var (
		outsideMemberships []models.OutsideMembership
		outsideMembership  models.OutsideMembership
	)

	query := fmt.Sprintf("%s %s", getOutsideMembershipListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&outsideMembership.OutsideMembershipID,
			&outsideMembership.MemberID,
			&outsideMembership.GymLocationID,
			&outsideMembership.GymID,
		)
		outsideMemberships = append(outsideMemberships, outsideMembership)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return outsideMemberships, nil
}

func GetOutsideMembershipCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getOutsideMembershipCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetOutsideMembership(outsideMembershipID int64) (*models.OutsideMembership, error) {
	var outsideMembership models.OutsideMembership

	row := store.DB.QueryRow(getOutsideMembershipQuery, outsideMembershipID)
	err := row.Scan(
		&outsideMembership.OutsideMembershipID,
		&outsideMembership.MemberID,
		&outsideMembership.GymLocationID,
		&outsideMembership.GymID,
	)

	if err != nil {
		return nil, err
	}

	return &outsideMembership, nil
}

func CreateOutsideMembership(outsideMembership models.OutsideMembership) (*models.OutsideMembership, error) {
	var created models.OutsideMembership

	row := store.DB.QueryRow(
		createOutsideMembershipQuery,
		outsideMembership.MemberID,
		outsideMembership.GymLocationID,
		outsideMembership.GymID,
	)
	err := row.Scan(
		&created.OutsideMembershipID,
		&created.MemberID,
		&created.GymLocationID,
		&created.GymID,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateOutsideMembership(outsideMembershipID int64, outsideMembership models.OutsideMembership) (*models.OutsideMembership, error) {
	var updated models.OutsideMembership

	row := store.DB.QueryRow(
		updateOutsideMembershipQuery,
		outsideMembership.MemberID,
		outsideMembership.GymLocationID,
		outsideMembership.GymID,
		outsideMembershipID,
	)
	err := row.Scan(
		&updated.OutsideMembershipID,
		&updated.MemberID,
		&updated.GymLocationID,
		&updated.GymID,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteOutsideMembership(outsideMembershipID int64) error {
	stmt, err := store.DB.Prepare(deleteOutsideMembershipQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(outsideMembershipID)
	if err != nil {
		return err
	}

	return nil
}

const getOutsideMembershipListQuery = `
SELECT *
FROM outside_memberships
`

const getOutsideMembershipQuery = `
SELECT *
FROM outside_memberships
WHERE outside_membership_id = $1
`

const createOutsideMembershipQuery = `
INSERT INTO outside_memberships (member_id, gym_location_id, gym_id)
VALUES ($1, $2, $3)
RETURNING outside_membership_id, member_id, gym_location_id, gym_id
`

const updateOutsideMembershipQuery = `
UPDATE outside_memberships
SET member_id = $1, gym_location_id = $2, gym_id = $3
WHERE outside_membership_id = $4
RETURNING outside_membership_id, member_id, gym_location_id, gym_id
`

const deleteOutsideMembershipQuery = `
DELETE
FROM outside_memberships
WHERE outside_membership_id = $1
`

const getOutsideMembershipCountQuery = `
SELECT count(*)
FROM outside_memberships
`
