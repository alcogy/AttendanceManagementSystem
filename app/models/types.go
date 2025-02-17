package models

import "time"

type Item struct {
	Date         string
	In           string
	Out          string
	BreakHours   string
	WorkingHours string
}

type Attendance struct {
	ID        string
	Date      *time.Time
	InTime    *time.Time
	OutTime   *time.Time
	BreakHour *time.Time
	UserID    string
}

type WorkOutBody struct {
	ID string `json:"id"`
}
