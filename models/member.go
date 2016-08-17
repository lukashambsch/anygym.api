package models

type Member struct {
	MemberId  int64  `json:"member_id"`
	UserId    *int64 `json:"user_id"`
	AddressId *int64 `json:"address_id"`
	PlanId    *int64 `json:"plan_id"`
	Active    bool   `json:"active"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
