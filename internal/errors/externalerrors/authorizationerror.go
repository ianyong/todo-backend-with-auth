package externalerrors

import "net/http"

type AuthorizationError struct{}

func (e *AuthorizationError) Error() string {
	return "You do not have permission to view the requested resource"
}

func (e *AuthorizationError) StatusCode() int {
	return http.StatusForbidden
}
