package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const secretKeyMinLength = 64

type JWTManager struct {
	secretKey     string
	tokenDuration *time.Duration
}

func NewJWTManager(secretKey string, tokenDuration *time.Duration) (*JWTManager, error) {
	if len(secretKey) < secretKeyMinLength {
		return nil, fmt.Errorf("secret key must be at least %d characters", secretKeyMinLength)
	}
	return &JWTManager{
		secretKey:     secretKey,
		tokenDuration: tokenDuration,
	}, nil
}

func (m *JWTManager) Generate(userID int64) (string, error) {
	// An expiry claim of 0 indicates that the token never expires.
	expiresAt := int64(0)
	if m.tokenDuration != nil {
		expiresAt = time.Now().Add(*m.tokenDuration).Unix()
	}

	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JWTManager) Verify(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected token signing method")
		}
		return []byte(m.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
