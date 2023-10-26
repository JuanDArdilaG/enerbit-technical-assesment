package apperrors

import (
	"fmt"
	"net/http"
)

type BadRequest struct {
	domain string
	id     string
}

func NewBadRequest(domain string, id string) Error {
	return &BadRequest{
		domain: domain,
		id:     id,
	}
}

func (e *BadRequest) Message() string {
	return fmt.Sprintf("invalid value %s for %s", e.id, e.domain)
}

func (e *BadRequest) Error() string {
	return e.Message()
}

func (e *BadRequest) Code() int {
	return http.StatusBadRequest
}
