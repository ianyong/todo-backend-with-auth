package externalerrors

import "net/http"

type AuthenticationError struct{}

func (e *AuthenticationError) Error() string {
	return "Authentication failed, please log in"
}

func (e *AuthenticationError) StatusCode() int {
	return http.StatusUnauthorized
}
