package externalerrors

import "net/http"

type AuthenticationError struct{}

func (e *AuthenticationError) Error() string {
	return "Invalid email/password combination"
}

func (e *AuthenticationError) StatusCode() int {
	return http.StatusUnauthorized
}
