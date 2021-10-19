package services

import (
	"github.com/jmoiron/sqlx"

	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/dbrepositories"
	"github.com/ianyong/todo-backend/internal/core/domainservices"
)

type Services struct {
	UserService *domainservices.UserService
	TodoService *domainservices.TodoService
}

func SetUp(db *sqlx.DB) *Services {
	userRepo := dbrepositories.NewUserDatabaseRepository(db)
	userService := domainservices.NewUserService(userRepo)

	todoRepo := dbrepositories.NewTodoDatabaseRepository(db)
	todoService := domainservices.NewTodoService(todoRepo)

	return &Services{
		UserService: userService,
		TodoService: todoService,
	}
}
