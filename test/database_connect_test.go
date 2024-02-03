package test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@/jura-event-management")
	if err != nil {
		panic(err)
	}

	defer db.Close()

}
