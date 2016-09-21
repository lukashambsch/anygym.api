package models

import "time"

type SupportRequest struct {
	SupportRequestID int64      `json:"support_request_id"`
	UserID           *int64     `json:"user_id"`
	SupportSourceID  *int64     `json:"support_source_id"`
	Content          string     `json:"content"`
	Notes            string     `json:"notes"`
	CreatedOn        time.Time  `json:"created_on"`
	ResolvedOn       *time.Time `json:"resolved_on"`
}
