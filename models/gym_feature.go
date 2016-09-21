package models

type GymFeature struct {
	GymFeatureID int64 `json:"gym_feature"`
	GymID        int64 `json:"gym_id"`
	FeatureID    int64 `json:"feature_id"`
}
