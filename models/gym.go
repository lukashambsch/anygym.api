package models

type Gym struct {
	GymId            int64   `json:"gym_id"`
	UserId           int64   `json:"user_id"`
	GymName          int64   `json:"gym_name"`
	MonthlyMemberFee float64 `json:"monthly_member_fee"`
}

type Gyms []Gym
