package repositories

import "github.com/ianyong/todo-backend/internal/core/domainmodels"

type User interface {
	GetByEmail(email string) (*domainmodels.User, error)
	Add(user *domainmodels.User) (*domainmodels.User, error)
}
