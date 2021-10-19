package userviews

import (
	"github.com/ianyong/todo-backend/internal/core/domainmodels"
)

type View struct {
	Email string `json:"email"`
}

func ViewFrom(user *domainmodels.User) View {
	return View{
		Email: user.Email,
	}
}
