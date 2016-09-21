package models

import "time"

type Visit struct {
	VisitID       int64      `json:"visit_id"`
	MemberID      int64      `json:"member_id"`
	GymLocationID int64      `json:"gym_location_id"`
	StatusID      int64      `json:"status_id"`
	CreatedOn     time.Time  `json:"created_on"`
	ModifiedOn    *time.Time `json:"modified_on"`
}
