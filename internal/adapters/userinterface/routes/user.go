package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/handlers/userhandlers"
	"github.com/ianyong/todo-backend/internal/services"
)

func GetUserRoutes(s *services.Services) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/login", api.WrapHandler(s, userhandlers.Login))
		r.Post("/register", api.WrapHandler(s, userhandlers.Register))
	}
}
