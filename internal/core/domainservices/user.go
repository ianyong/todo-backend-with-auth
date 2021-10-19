package domainservices

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/ianyong/todo-backend/internal/auth"
	"github.com/ianyong/todo-backend/internal/core/domainmodels"
	"github.com/ianyong/todo-backend/internal/errors/externalerrors"
	"github.com/ianyong/todo-backend/internal/ports/repositories"
	"github.com/ianyong/todo-backend/internal/utils"
)

type UserService struct {
	userRepo   repositories.User
	jwtManager auth.JWTManager
}

func NewUserService(userRepo repositories.User, jwtManager *auth.JWTManager) *UserService {
	return &UserService{
		userRepo:   userRepo,
		jwtManager: *jwtManager,
	}
}

func (s *UserService) Login(email string, password string) (*domainmodels.User, string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		// Prevent enumeration attacks via timing.
		_ = bcrypt.CompareHashAndPassword(utils.DummyPasswordHash, []byte(password))
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, "", &externalerrors.AuthenticationError{}
	}
	if err != nil {
		return nil, "", err
	}

	token, err := s.jwtManager.Generate(user.Email)
	if err != nil {
		return nil, "", fmt.Errorf("unable to generate JWT: %v", err)
	}

	return user, token, nil
}

func (s *UserService) Register(email string, password string) (*domainmodels.User, string, error) {
	hashedPassword, err := utils.Hash(password)
	if err != nil {
		return nil, "", fmt.Errorf("error hashing password: %w", err)
	}

	user, err := s.userRepo.Add(&domainmodels.User{
		Email:          email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return nil, "", err
	}

	token, err := s.jwtManager.Generate(user.Email)
	if err != nil {
		return nil, "", fmt.Errorf("unable to generate JWT: %v", err)
	}

	return user, token, nil
}
