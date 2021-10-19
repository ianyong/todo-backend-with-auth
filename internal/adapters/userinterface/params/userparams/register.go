package userparams

import (
	"gopkg.in/guregu/null.v4"

	"github.com/ianyong/todo-backend/internal/errors/externalerrors"
)

type RegisterParams struct {
	Email    null.String `json:"email"`
	Password null.String `json:"password"`
}

func (params *RegisterParams) Validate() error {
	if params.Email.IsZero() {
		return &externalerrors.MissingParamError{Param: "email"}
	}
	if params.Password.IsZero() {
		return &externalerrors.MissingParamError{Param: "password"}
	}
	return nil
}
