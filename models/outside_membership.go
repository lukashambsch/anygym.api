package models

type OutsideMembership struct {
	OutsideMembershipId int64 `json:"outside_membership_id"`
	MemberId            int64 `json:"member_id"`
	LocationId          int64 `json:"location_id"`
	GymId               int64 `json:"gym_id"`
}

type OutsideMemberships []OutsideMembership
