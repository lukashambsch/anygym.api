package datastore

import (
	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetGymList() ([]models.Gym, error) {
	var (
		gyms []models.Gym
		gym  models.Gym
	)

	rows, err := store.DB.Query(getGymListQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&gym.GymId, &gym.UserId, &gym.GymName, &gym.MonthlyMemberFee)
		gyms = append(gyms, gym)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return gyms, nil
}

func GetGymCount() (*int, error) {
	var count int

	row := store.DB.QueryRow(getGymCountQuery)
	err := row.Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetGym(gymId int64) (*models.Gym, error) {
	var gym models.Gym

	row := store.DB.QueryRow(getGymQuery, gymId)
	err := row.Scan(&gym.GymId, &gym.UserId, &gym.GymName, &gym.MonthlyMemberFee)

	if err != nil {
		return nil, err
	}

	return &gym, nil
}

func CreateGym(gym models.Gym) (*models.Gym, error) {
	var created models.Gym

	row := store.DB.QueryRow(createGymQuery, gym.UserId, gym.GymName, gym.MonthlyMemberFee)
	err := row.Scan(&created.GymId, &created.UserId, &created.GymName, &created.MonthlyMemberFee)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateGym(gymId int64, gym models.Gym) (*models.Gym, error) {
	var updated models.Gym
	t, err := store.DB.Begin()

	if err != nil {
		return nil, err
	}

	row := t.QueryRow(updateGymQuery, gym.UserId, gym.GymName, gym.MonthlyMemberFee, gymId)
	err = row.Scan(&updated.GymId, &updated.UserId, &updated.GymName, &updated.MonthlyMemberFee)

	if err != nil {
		t.Rollback()
		return nil, err
	}

	t.Commit()

	return &updated, nil
}

func DeleteGym(gymId int64) error {
	stmt, err := store.DB.Prepare(deleteGymQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(gymId)
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
