package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/json"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/params/userparams"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/views/userviews"
	"github.com/ianyong/todo-backend/internal/services"
)

func Register(r *http.Request, s *services.Services) (*api.Response, error) {
	var registerParams userparams.RegisterParams
	err := json.DecodeParams(r.Body, &registerParams)
	if err != nil {
		return nil, fmt.Errorf("unable to decode request body into params: %w", err)
	}

	err = registerParams.Validate()
	if err != nil {
		return nil, fmt.Errorf("params failed validation: %w", err)
	}

	user, err := s.UserService.Register(registerParams.Email.ValueOrZero(), registerParams.Password.ValueOrZero())
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	userView := userviews.ViewFrom(user)

	data, err := json.EncodeView(userView)
	if err != nil {
		return nil, err
	}

	return &api.Response{
		Payload: data,
		Code:    http.StatusCreated,
	}, nil
}
