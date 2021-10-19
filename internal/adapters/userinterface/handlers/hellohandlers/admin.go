package hellohandlers

import (
	"errors"
	"net/http"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/contextkeys"
	"github.com/ianyong/todo-backend/internal/errors/externalerrors"
	"github.com/ianyong/todo-backend/internal/services"
)

func Admin(r *http.Request, s *services.Services) (*api.Response, error) {
	userEmail, ok := r.Context().Value(contextkeys.UserEmail).(string)
	if !ok {
		return nil, errors.New("unable to retrieve user email")
	}

	user, err := s.UserService.GetUserByEmail(userEmail)
	if err != nil {
		return nil, err
	}

	if user.Role != "admin" {
		return nil, &externalerrors.AuthorizationError{}
	}

	return &api.Response{
		Messages: api.StatusMessages{
			api.SuccessMessage("Hello admin!"),
		},
		Code: http.StatusOK,
	}, nil
}
