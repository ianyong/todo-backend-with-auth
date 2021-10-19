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

func Login(r *http.Request, s *services.Services) (*api.Response, error) {
	var loginParams userparams.LoginParams
	err := json.DecodeParams(r.Body, &loginParams)
	if err != nil {
		return nil, fmt.Errorf("unable to decode request body into params: %w", err)
	}

	err = loginParams.Validate()
	if err != nil {
		return nil, fmt.Errorf("params failed validation: %w", err)
	}

	user, accessToken, err := s.UserService.Login(loginParams.Email.ValueOrZero(), loginParams.Password.ValueOrZero())
	if err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}

	userView := userviews.ViewFrom(user)

	data, err := json.EncodeView(userView)
	if err != nil {
		return nil, err
	}

	return &api.Response{
		Payload:     data,
		Code:        http.StatusOK,
		AccessToken: accessToken,
	}, nil
}
