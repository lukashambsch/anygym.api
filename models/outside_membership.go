package models

type OutsideMembership struct {
	OutsideMembershipId int64 `json:"outside_membership_id"`
	MemberId            int64 `json:"member_id"`
	GymLocationId       int64 `json:"gym_location_id"`
	GymId               int64 `json:"gym_id"`
}
