package models

type Plan struct {
	PlanId   int64   `json:"plan_id"`
	PlanName string  `json:"plan_name"`
	Price    float64 `json:"price"`
}

type Plans []Plan
