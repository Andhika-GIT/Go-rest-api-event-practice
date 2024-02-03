package domain

import "time"

type Attendee struct {
	Id         int32
	User_id    int32
	Event_id   int32
	Created_at time.Time
	Updated_at time.Time
}
