package models

type Address struct {
	AddressID     int64    `json:"address_id"`
	Country       string   `json:"country"`
	StateRegion   string   `json:"state_region"`
	City          string   `json:"city"`
	PostalArea    string   `json:"postal_area"`
	StreetAddress string   `json:"street_address"`
	Latitude      *float64 `json:"latitude"`
	Longitude     *float64 `json:"longitude"`
}
