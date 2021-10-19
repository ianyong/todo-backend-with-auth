package domainservices

import (
	"github.com/ianyong/todo-backend/internal/core/domainmodels"
	"github.com/ianyong/todo-backend/internal/ports/repositories"
)

type UserService struct {
	userRepo repositories.User
}

func NewUserService(userRepo repositories.User) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Login(email string, password string) (*domainmodels.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	// TODO: Implement login logic.
	return user, nil
}

func (s *UserService) Register(email string, password string) (*domainmodels.User, error) {
	// TODO: Implement register logic.
	user, err := s.userRepo.Add(&domainmodels.User{
		Email:          email,
		HashedPassword: password,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
