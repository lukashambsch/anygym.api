package models

import "time"

type SupportRequest struct {
	SupportRequestId int64     `json:"support_request_id"`
	UserId           int64     `json:"user_id"`
	SupportSourceId  int64     `json:"support_source_id"`
	Content          string    `json:"content"`
	CreatedOn        time.Time `json:"created_on"`
	ResolvedOn       time.Time `json:"resolved_on"`
}

type SupportRequests []SupportRequest
