package models

import "time"

type Visit struct {
	VisitId       int64      `json:"visit_id"`
	MemberId      int64      `json:"member_id"`
	GymLocationId int64      `json:"gym_location_id"`
	StatusId      int64      `json:"status_id"`
	CreatedOn     time.Time  `json:"created_on"`
	ModifiedOn    *time.Time `json:"modified_on"`
}
