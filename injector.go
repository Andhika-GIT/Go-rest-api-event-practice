//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Andhika-GIT/Go-REST-Event-Management/app"
	"github.com/Andhika-GIT/Go-REST-Event-Management/controller"
	"github.com/Andhika-GIT/Go-REST-Event-Management/repository"
	"github.com/Andhika-GIT/Go-REST-Event-Management/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func InitializedServer(option ...validator.Option) *chi.Mux {
	wire.Build(app.NewDB, app.NewRouter, validator.New, repository.NewUserRepository, service.NewUserService, controller.NewUserController)

	return nil
}
