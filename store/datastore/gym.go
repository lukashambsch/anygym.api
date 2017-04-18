package datastore

import (
	"fmt"

	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store"
)

func GetGymList(where string) ([]models.Gym, error) {
	var (
		gyms []models.Gym
		gym  models.Gym
	)

	query := fmt.Sprintf("%s %s", getGymListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&gym.GymID, &gym.UserID, &gym.GymName, &gym.MonthlyMemberFee)
		gyms = append(gyms, gym)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return gyms, nil
}

func GetGymCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getGymCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetGym(gymID int64) (*models.Gym, error) {
	var gym models.Gym

	row := store.DB.QueryRow(getGymQuery, gymID)
	err := row.Scan(&gym.GymID, &gym.UserID, &gym.GymName, &gym.MonthlyMemberFee)

	if err != nil {
		return nil, err
	}

	return &gym, nil
}

func CreateGym(gym models.Gym) (*models.Gym, error) {
	var created models.Gym

	row := store.DB.QueryRow(createGymQuery, gym.UserID, gym.GymName, gym.MonthlyMemberFee)
	err := row.Scan(&created.GymID, &created.UserID, &created.GymName, &created.MonthlyMemberFee)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateGym(gymID int64, gym models.Gym) (*models.Gym, error) {
	var updated models.Gym

	row := store.DB.QueryRow(updateGymQuery, gym.UserID, gym.GymName, gym.MonthlyMemberFee, gymID)
	err := row.Scan(&updated.GymID, &updated.UserID, &updated.GymName, &updated.MonthlyMemberFee)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteGym(gymID int64) error {
	stmt, err := store.DB.Prepare(deleteGymQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(gymID)
	if err != nil {
		return err
	}

	return nil
}

const getGymListQuery = `
SELECT *
FROM gyms
`

const getGymQuery = `
SELECT *
FROM gyms
WHERE gym_id = $1
`

const createGymQuery = `
INSERT INTO gyms (user_id, gym_name, monthly_member_fee)
VALUES ($1, $2, $3)
RETURNING gym_id, user_id, gym_name, monthly_member_fee
`

const updateGymQuery = `
UPDATE gyms
SET user_id = $1, gym_name = $2, monthly_member_fee = $3
WHERE gym_id = $4
RETURNING gym_id, user_id, gym_name, monthly_member_fee
`

const deleteGymQuery = `
DELETE
FROM gyms
WHERE gym_id = $1
`

const getGymCountQuery = `
SELECT count(*)
FROM gyms
`
