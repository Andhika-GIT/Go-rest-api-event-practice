package test

import (
	"testing"

	"github.com/Andhika-GIT/Go-REST-Event-Management/app"
	"github.com/stretchr/testify/assert"
)

var db = app.NewDB()

func TestConnection(t *testing.T) {

	assert.NotNil(t, db)
}
