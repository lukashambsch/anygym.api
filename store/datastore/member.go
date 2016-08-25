package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetMemberList(where string) ([]models.Member, error) {
	var (
		members []models.Member
		member  models.Member
	)

	query := fmt.Sprintf("%s %s", getMemberListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&member.MemberId,
			&member.UserId,
			&member.AddressId,
			&member.FirstName,
			&member.LastName,
		)
		members = append(members, member)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return members, nil
}

func GetMemberCount() (*int, error) {
	var count int

	row := store.DB.QueryRow(getMemberCountQuery)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetMember(memberId int64) (*models.Member, error) {
	var member models.Member

	row := store.DB.QueryRow(getMemberQuery, memberId)
	err := row.Scan(
		&member.MemberId,
		&member.UserId,
		&member.AddressId,
		&member.FirstName,
		&member.LastName,
	)

	if err != nil {
		return nil, err
	}

	return &member, nil
}

func CreateMember(member models.Member) (*models.Member, error) {
	var created models.Member

	row := store.DB.QueryRow(
		createMemberQuery,
		member.UserId,
		member.AddressId,
		member.FirstName,
		member.LastName,
	)
	err := row.Scan(
		&created.MemberId,
		&created.UserId,
		&created.AddressId,
		&created.FirstName,
		&created.LastName,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateMember(memberId int64, member models.Member) (*models.Member, error) {
	var updated models.Member

	row := store.DB.QueryRow(
		updateMemberQuery,
		member.UserId,
		member.AddressId,
		member.FirstName,
		member.LastName,
		memberId,
	)
	err := row.Scan(
		&updated.MemberId,
		&updated.UserId,
		&updated.AddressId,
		&updated.FirstName,
		&updated.LastName,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteMember(memberId int64) error {
	stmt, err := store.DB.Prepare(deleteMemberQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(memberId)
	if err != nil {
		return err
	}

	return nil
}

const getMemberListQuery = `
SELECT *
FROM members
`

const getMemberQuery = `
SELECT *
FROM members
WHERE member_id = $1
`

const createMemberQuery = `
INSERT INTO members (user_id, address_id, first_name, last_name)
VALUES ($1, $2, $3, $4)
RETURNING member_id, user_id, address_id, first_name, last_name
`

const updateMemberQuery = `
UPDATE members
SET user_id = $1, address_id = $2, first_name = $3, last_name = $4
WHERE member_id = $5
RETURNING member_id, user_id, address_id, first_name, last_name
`

const deleteMemberQuery = `
DELETE
FROM members
WHERE member_id = $1
`

const getMemberCountQuery = `
SELECT count(*)
FROM members
`
