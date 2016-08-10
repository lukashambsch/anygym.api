package models

type Status struct {
	StatusId   int64  `json:"status_id"`
	StatusName string `json:"status_name"`
}

type Statuses []Status
