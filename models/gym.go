package models

type Gym struct {
	GymID            int64    `json:"gym_id"`
	UserID           *int64   `json:"user_id"`
	GymName          string   `json:"gym_name"`
	MonthlyMemberFee *float64 `json:"monthly_member_fee"`
}
