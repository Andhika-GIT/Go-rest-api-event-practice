package app

import (
	"github.com/Andhika-GIT/Go-REST-Event-Management/controller"
	"github.com/go-chi/chi/v5"
)

func NewRouter(userController controller.UserController) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/api/users", userController.Create)

	return r
}
