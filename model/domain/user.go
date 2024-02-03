package domain

import "time"

type User struct {
	Id         int32
	Name       string
	Email      string
	Password   string
	Created_at time.Time
	Updated_at time.Time
}
