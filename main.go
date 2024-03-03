package main

import (
	"net/http"

	"github.com/Andhika-GIT/Go-REST-Event-Management/helper"
)

func main() {

	r := InitializedServer()

	err := http.ListenAndServe("localhost:8000", r)
	helper.PanicIfError(err)
}
