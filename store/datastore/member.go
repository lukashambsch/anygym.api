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
			&member.MemberID,
			&member.UserID,
			&member.AddressID,
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

func GetMemberCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getMemberCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetMember(memberID int64) (*models.Member, error) {
	var member models.Member

	row := store.DB.QueryRow(getMemberQuery, memberID)
	err := row.Scan(
		&member.MemberID,
		&member.UserID,
		&member.AddressID,
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
		member.UserID,
		member.AddressID,
		member.FirstName,
		member.LastName,
	)
	err := row.Scan(
		&created.MemberID,
		&created.UserID,
		&created.AddressID,
		&created.FirstName,
		&created.LastName,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateMember(memberID int64, member models.Member) (*models.Member, error) {
	var updated models.Member

	row := store.DB.QueryRow(
		updateMemberQuery,
		member.UserID,
		member.AddressID,
		member.FirstName,
		member.LastName,
		memberID,
	)
	err := row.Scan(
		&updated.MemberID,
		&updated.UserID,
		&updated.AddressID,
		&updated.FirstName,
		&updated.LastName,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteMember(memberID int64) error {
	stmt, err := store.DB.Prepare(deleteMemberQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(memberID)
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
