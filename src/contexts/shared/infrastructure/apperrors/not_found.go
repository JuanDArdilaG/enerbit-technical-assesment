package apperrors

import (
	"fmt"
	"net/http"
)

type ErrorNotFound struct {
	domain string
	id     string
}

func NewErrorNotFound(domain string, id string) Error {
	return &ErrorNotFound{
		domain: domain,
		id:     id,
	}
}

func (e *ErrorNotFound) Message() string {
	return fmt.Sprintf("%s with id %s not found", e.domain, e.id)
}

func (e *ErrorNotFound) Error() string {
	return e.Message()
}

func (e *ErrorNotFound) Code() int {
	return http.StatusNotFound
}
