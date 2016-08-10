package models

type Feature struct {
	FeatureId          int64  `json:"feature_id"`
	FeatureName        string `json:"feature_name"`
	FeatureDescription string `json:"feature_description"`
}

type Features []Feature
