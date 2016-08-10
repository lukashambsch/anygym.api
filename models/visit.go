package models

import "time"

type Visit struct {
	VisitId    int64     `json:"visit_id"`
	MemberId   int64     `json:"member_id"`
	LocationId int64     `json:"location_id"`
	StatusId   int64     `json:"status_id"`
	CreatedOn  time.Time `json:"created_on"`
}

type Visits []Visit
