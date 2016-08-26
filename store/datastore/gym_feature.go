package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetGymFeatureList(where string) ([]models.GymFeature, error) {
	var (
		gymFeatures []models.GymFeature
		gymFeature  models.GymFeature
	)

	query := fmt.Sprintf("%s %s", getGymFeatureListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&gymFeature.GymFeatureId, &gymFeature.GymId, &gymFeature.FeatureId)
		gymFeatures = append(gymFeatures, gymFeature)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return gymFeatures, nil
}

func GetGymFeatureCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getGymFeatureCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetGymFeature(gymFeatureId int64) (*models.GymFeature, error) {
	var gymFeature models.GymFeature

	row := store.DB.QueryRow(getGymFeatureQuery, gymFeatureId)
	err := row.Scan(&gymFeature.GymFeatureId, &gymFeature.GymId, &gymFeature.FeatureId)
	if err != nil {
		return nil, err
	}

	return &gymFeature, nil
}

func CreateGymFeature(gymFeature models.GymFeature) (*models.GymFeature, error) {
	var created models.GymFeature

	row := store.DB.QueryRow(createGymFeatureQuery, gymFeature.GymId, gymFeature.FeatureId)
	err := row.Scan(&created.GymFeatureId, &created.GymId, &created.FeatureId)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateGymFeature(gymFeatureId int64, gymFeature models.GymFeature) (*models.GymFeature, error) {
	var updated models.GymFeature

	row := store.DB.QueryRow(updateGymFeatureQuery, gymFeature.GymId, gymFeature.FeatureId, gymFeatureId)
	err := row.Scan(&updated.GymFeatureId, &updated.GymId, &updated.FeatureId)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteGymFeature(gymFeatureId int64) error {
	stmt, err := store.DB.Prepare(deleteGymFeatureQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(gymFeatureId)
	if err != nil {
		return err
	}

	return nil
}

const getGymFeatureListQuery = `
SELECT *
FROM gym_features
`

const getGymFeatureQuery = `
SELECT *
FROM gym_features
WHERE gym_feature_id = $1
`

const createGymFeatureQuery = `
INSERT INTO gym_features (gym_id, feature_id)
VALUES ($1, $2)
RETURNING gym_feature_id, gym_id, feature_id
`

const updateGymFeatureQuery = `
UPDATE gym_features
SET gym_id = $1, feature_id = $2
WHERE gym_feature_id = $3
RETURNING gym_feature_id, gym_id, feature_id
`

const deleteGymFeatureQuery = `
DELETE
FROM gym_features
WHERE gym_feature_id = $1
`

const getGymFeatureCountQuery = `
SELECT count(*)
FROM gym_features
`
