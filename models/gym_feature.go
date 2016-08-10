package models

type GymFeature struct {
	GymFeatureId int64 `json:"gym_feature"`
	GymId        int64 `json:"gym_id"`
	FeatureId    int64 `json:"feature_id"`
}
