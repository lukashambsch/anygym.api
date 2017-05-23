package models

type GymLocation struct {
	GymLocationID    int64    `json:"gym_location_id"`
	GymID            int64    `json:"gym_id"`
	AddressID        int64    `json:"address_id"`
	LocationName     string   `json:"location_name"`
	PhoneNumber      string   `json:"phone_number"`
	WebsiteUrl       string   `json:"website_url"`
	InNetwork        bool     `json:"in_network"`
	MonthlyMemberFee *float64 `json:"monthly_member_fee"`
	Address          Address  `json:"address"`
}
