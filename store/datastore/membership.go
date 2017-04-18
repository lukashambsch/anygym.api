package datastore

import (
	"fmt"

	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store"
)

func GetMembershipList(where string) ([]models.Membership, error) {
	var (
		memberships []models.Membership
		membership  models.Membership
	)

	query := fmt.Sprintf("%s %s", getMembershipListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&membership.MembershipID,
			&membership.PlanID,
			&membership.MemberID,
			&membership.StartDate,
			&membership.RenewDate,
			&membership.EndDate,
			&membership.Active,
		)
		memberships = append(memberships, membership)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return memberships, nil
}

func GetMembershipCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getMembershipCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetMembership(membershipID int64) (*models.Membership, error) {
	var membership models.Membership

	row := store.DB.QueryRow(getMembershipQuery, membershipID)
	err := row.Scan(
		&membership.MembershipID,
		&membership.PlanID,
		&membership.MemberID,
		&membership.StartDate,
		&membership.RenewDate,
		&membership.EndDate,
		&membership.Active,
	)

	if err != nil {
		return nil, err
	}

	return &membership, nil
}

func CreateMembership(membership models.Membership) (*models.Membership, error) {
	var created models.Membership

	row := store.DB.QueryRow(
		createMembershipQuery,
		membership.PlanID,
		membership.MemberID,
		membership.StartDate,
		membership.RenewDate,
		membership.EndDate,
		membership.Active,
	)
	err := row.Scan(
		&created.MembershipID,
		&created.PlanID,
		&created.MemberID,
		&created.StartDate,
		&created.RenewDate,
		&created.EndDate,
		&created.Active,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateMembership(membershipID int64, membership models.Membership) (*models.Membership, error) {
	var updated models.Membership

	row := store.DB.QueryRow(
		updateMembershipQuery,
		membership.PlanID,
		membership.MemberID,
		membership.StartDate,
		membership.RenewDate,
		membership.EndDate,
		membership.Active,
		membershipID,
	)
	err := row.Scan(
		&updated.MembershipID,
		&updated.PlanID,
		&updated.MemberID,
		&updated.StartDate,
		&updated.RenewDate,
		&updated.EndDate,
		&updated.Active,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteMembership(membershipID int64) error {
	stmt, err := store.DB.Prepare(deleteMembershipQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(membershipID)
	if err != nil {
		return err
	}

	return nil
}

const getMembershipListQuery = `
SELECT *
FROM memberships
`

const getMembershipQuery = `
SELECT *
FROM memberships
WHERE membership_id = $1
`

const createMembershipQuery = `
INSERT INTO memberships (plan_id, member_id, start_date, renew_date, end_date, active)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING membership_id, plan_id, member_id, start_date, renew_date, end_date, active
`

const updateMembershipQuery = `
UPDATE memberships
SET plan_id = $1, member_id = $2, start_date = $3, renew_date = $4, end_date = $5, active = $6
WHERE membership_id = $7
RETURNING membership_id, plan_id, member_id, start_date, renew_date, end_date, active
`

const deleteMembershipQuery = `
DELETE
FROM memberships
WHERE membership_id = $1
`

const getMembershipCountQuery = `
SELECT count(*)
FROM memberships
`
