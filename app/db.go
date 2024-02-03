package app

import (
	"database/sql"
	"time"

	"github.com/Andhika-GIT/Go-REST-Event-Management/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root/jura-event-management")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
