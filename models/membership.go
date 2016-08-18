package models

import "time"

type Membership struct {
	MembershipId int64     `json:"membership_id"`
	PlanId       *int64    `json:"plan_id"`
	MemberId     *int64    `json:"member_id"`
	StartDate    time.Time `json:"start_date"`
	RenewDate    time.Time `json:"renew_date"`
	EndDate      time.Time `json:"end_date"`
	Active       bool      `json:"active"`
}
