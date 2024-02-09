package domain

import "time"

type User struct {
	Id         int32     `gorm:"primary_key;column:id"`
	Name       string    `gorm:"column:name"`
	Email      string    `gorm:"column:email"`
	Password   string    `gorm:"column:password "`
	Created_at time.Time `gorm:"column:created_at;autoCreateTime"`
	Updated_at time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *User) TableName() string {
	return "users"
}
