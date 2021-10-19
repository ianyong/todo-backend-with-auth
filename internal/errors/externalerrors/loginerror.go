package externalerrors

import "net/http"

type LoginError struct{}

func (e *LoginError) Error() string {
	return "Invalid email/password combination"
}

func (e *LoginError) StatusCode() int {
	return http.StatusUnauthorized
}
