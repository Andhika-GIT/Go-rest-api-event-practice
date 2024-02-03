package main

import (
	"net/http"

	"github.com/Andhika-GIT/Go-REST-Event-Management/app"
	"github.com/Andhika-GIT/Go-REST-Event-Management/helper"
)

func main() {
	// db := app.NewDB()
	r := app.NewRouter()

	err := http.ListenAndServe("localhost:8000", r)
	helper.PanicIfError(err)
}
