package models

type Location struct {
	LocationId       int64   `json:"location_id"`
	AddressId        int64   `json:"address_id"`
	LocationName     string  `json:"location_name"`
	PhoneNumber      string  `json:"phone_number"`
	WebsiteUrl       string  `json:"website_url"`
	InNetwork        bool    `json:"in_network"`
	MonthlyMemberFee float64 `json:"monthly_member_fee"`
}
