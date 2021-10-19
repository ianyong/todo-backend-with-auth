package utils

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/ianyong/todo-backend/internal/errors/externalerrors"
)

const HashingCost = 10

var DummyPasswordHash, _ = bcrypt.GenerateFromPassword([]byte(""), HashingCost)

func Hash(password string) (string, error) {
	err := validate(password)
	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), HashingCost)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func validate(password string) error {
	if len(password) < 8 {
		return &externalerrors.PasswordRequirementsError{
			Message: "Password must contain at least 8 characters",
		}
	}
	return nil
}
