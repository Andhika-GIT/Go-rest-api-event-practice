package main

import (
	"net/http"

	"github.com/Andhika-GIT/Go-REST-Event-Management/helper"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := InitializedServer()

	err := http.ListenAndServe("localhost:8000", r)
	helper.PanicIfError(err)
}
