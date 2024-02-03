package domain

import "time"

type Event struct {
	Id          int32
	User_id     int32
	Name        string
	Description string
	Start_time  time.Time
	Created_at  time.Time
	Updated_at  time.Time
}
