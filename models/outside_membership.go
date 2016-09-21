package models

type OutsideMembership struct {
	OutsideMembershipID int64  `json:"outside_membership_id"`
	MemberID            int64  `json:"member_id"`
	GymLocationID       *int64 `json:"gym_location_id"`
	GymID               *int64 `json:"gym_id"`
}
