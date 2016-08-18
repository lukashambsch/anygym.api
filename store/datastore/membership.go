package datastore

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetMembershipList() ([]models.Membership, error) {
	var (
		memberships []models.Membership
		membership  models.Membership
	)

	rows, err := store.DB.Query(getMembershipListQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&membership.MembershipId,
			&membership.PlanId,
			&membership.MemberId,
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

func GetMembershipCount() (*int, error) {
	var count int

	row := store.DB.QueryRow(getMembershipCountQuery)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetMembership(membershipId int64) (*models.Membership, error) {
	var membership models.Membership

	row := store.DB.QueryRow(getMembershipQuery, membershipId)
	err := row.Scan(
		&membership.MembershipId,
		&membership.PlanId,
		&membership.MemberId,
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
		membership.PlanId,
		membership.MemberId,
		membership.StartDate,
		membership.RenewDate,
		membership.EndDate,
		membership.Active,
	)
	err := row.Scan(
		&created.MembershipId,
		&created.PlanId,
		&created.MemberId,
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

func UpdateMembership(membershipId int64, membership models.Membership) (*models.Membership, error) {
	var updated models.Membership

	row := store.DB.QueryRow(
		updateMembershipQuery,
		membership.PlanId,
		membership.MemberId,
		membership.StartDate,
		membership.RenewDate,
		membership.EndDate,
		membership.Active,
		membershipId,
	)
	err := row.Scan(
		&updated.MembershipId,
		&updated.PlanId,
		&updated.MemberId,
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

func DeleteMembership(membershipId int64) error {
	stmt, err := store.DB.Prepare(deleteMembershipQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(membershipId)
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
