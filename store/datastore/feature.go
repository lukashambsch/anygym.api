package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetFeatureList(where string) ([]models.Feature, error) {
	var (
		features []models.Feature
		feature  models.Feature
	)

	query := fmt.Sprintf("%s %s", getFeatureListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&feature.FeatureId, &feature.FeatureName, &feature.FeatureDescription)
		features = append(features, feature)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return features, nil
}

func GetFeatureCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getFeatureCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetFeature(featureId int64) (*models.Feature, error) {
	var feature models.Feature

	row := store.DB.QueryRow(getFeatureQuery, featureId)
	err := row.Scan(&feature.FeatureId, &feature.FeatureName, &feature.FeatureDescription)
	if err != nil {
		return nil, err
	}

	return &feature, nil
}

func CreateFeature(feature models.Feature) (*models.Feature, error) {
	var created models.Feature

	row := store.DB.QueryRow(createFeatureQuery, feature.FeatureName, feature.FeatureDescription)
	err := row.Scan(&created.FeatureId, &created.FeatureName, &created.FeatureDescription)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdateFeature(featureId int64, feature models.Feature) (*models.Feature, error) {
	var updated models.Feature

	row := store.DB.QueryRow(updateFeatureQuery, feature.FeatureName, feature.FeatureDescription, featureId)
	err := row.Scan(&updated.FeatureId, &updated.FeatureName, &updated.FeatureDescription)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteFeature(featureId int64) error {
	stmt, err := store.DB.Prepare(deleteFeatureQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(featureId)
	if err != nil {
		return err
	}

	return nil
}

const getFeatureListQuery = `
SELECT *
FROM features
`

const getFeatureQuery = `
SELECT *
FROM features
WHERE feature_id = $1
`

const createFeatureQuery = `
INSERT INTO features (feature_name, feature_description)
VALUES ($1, $2)
RETURNING feature_id, feature_name, feature_description
`

const updateFeatureQuery = `
UPDATE features
SET feature_name = $1, feature_description = $2
WHERE feature_id = $3
RETURNING feature_id, feature_name, feature_description
`

const deleteFeatureQuery = `
DELETE
FROM features
WHERE feature_id = $1
`

const getFeatureCountQuery = `
SELECT count(*)
FROM features
`
