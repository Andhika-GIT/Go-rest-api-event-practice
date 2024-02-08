package domain

import "time"

type Attendee struct {
	Id         int32     `gorm:"primary_key;column:id"`
	User_id    int32     `gorm:"column:user_id"`
	Event_id   int32     `gorm:"column:event_id"`
	Created_at time.Time `gorm:"column:created_at"`
	Updated_at time.Time `gorm:"column:updated_at"`
}

func (a *Attendee) TableName() string {
	return "attendees"
}
