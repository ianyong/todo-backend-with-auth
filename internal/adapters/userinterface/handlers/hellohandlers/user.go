package hellohandlers

import (
	"net/http"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/services"
)

func User(r *http.Request, s *services.Services) (*api.Response, error) {
	return &api.Response{
		Messages: api.StatusMessages{
			api.SuccessMessage("Hello user!"),
		},
		Code: http.StatusOK,
	}, nil
}
