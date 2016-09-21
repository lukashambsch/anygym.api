package models

import "time"

type Membership struct {
	MembershipID int64      `json:"membership_id"`
	PlanID       *int64     `json:"plan_id"`
	MemberID     *int64     `json:"member_id"`
	StartDate    time.Time  `json:"start_date"`
	RenewDate    *time.Time `json:"renew_date"`
	EndDate      *time.Time `json:"end_date"`
	Active       bool       `json:"active"`
}
