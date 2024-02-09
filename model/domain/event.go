package domain

import "time"

type Event struct {
	Id          int32     `gorm:"primary_key;column:id"`
	User_id     int32     `gorm:"column:user_id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Start_time  time.Time `gorm:"column:start_time"`
	End_time    time.Time `gorm:"column:end_time"`
	Created_at  time.Time `gorm:"column:created_at;autoCreateTime"`
	Updated_at  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (e *Event) TableName() string {
	return "events"
}
