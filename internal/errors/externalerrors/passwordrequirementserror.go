package externalerrors

import "net/http"

type PasswordRequirementsError struct {
	Message string
}

func (e *PasswordRequirementsError) Error() string {
	return e.Message
}

func (e *PasswordRequirementsError) StatusCode() int {
	return http.StatusBadRequest
}
