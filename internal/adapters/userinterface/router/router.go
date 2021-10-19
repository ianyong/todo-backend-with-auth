package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/handlers"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/routes"
	"github.com/ianyong/todo-backend/internal/config"
	"github.com/ianyong/todo-backend/internal/services"
)

// SetUp sets up the middleware stack and routes for a chi.Mux and returns it.
func SetUp(s *services.Services, cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	setUpMiddleware(r, cfg)
	setUpRoutes(r, s)
	return r
}

// setUpMiddleware sets up the middleware stack for a chi.Router.
func setUpMiddleware(r chi.Router, cfg *config.Config) {
	// Injects a request ID in the context of each request.
	r.Use(middleware.RequestID)
	// Sets a http.Request's RemoteAddr to that of either the X-Forwarded-For or X-Real-IP header.
	r.Use(middleware.RealIP)
	// Logs the start and end of each request.
	r.Use(middleware.Logger)
	// Recovers from panics and return a 500 Internal Service Error.
	r.Use(middleware.Recoverer)
	// Returns a 504 Gateway Timeout after 1 min.
	r.Use(middleware.Timeout(time.Minute))
	// Sets up Cross-Origin Resource Sharing.
	r.Use(corsMiddleware(cfg.Environment))
}

// setUpRoutes sets up the routes for a chi.Router. All API routes are namespaced with '/api/v1'.
func setUpRoutes(r chi.Router, s *services.Services) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/", routes.GetUserRoutes(s))
		r.Route("/todos", routes.GetTodoRoutes(s))
	})
	r.NotFound(api.WrapHandler(s, handlers.NotFound))
}
