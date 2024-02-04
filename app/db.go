package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Andhika-GIT/Go-REST-Event-Management/helper"
)

func NewDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/jura-event-management?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	helper.PanicIfError(err)

	return db
}
