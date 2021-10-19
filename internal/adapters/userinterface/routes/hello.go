package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/handlers/hellohandlers"
	"github.com/ianyong/todo-backend/internal/services"
)

func GetHelloRoutes(s *services.Services) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/user", api.WrapHandler(s, hellohandlers.User))
		r.Get("/admin", api.WrapHandler(s, hellohandlers.Admin))
	}
}
