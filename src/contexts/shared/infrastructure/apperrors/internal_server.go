package apperrors

import (
	"fmt"
	"net/http"
)

type ErrorInternalServer struct {
	message string
}

func NewErrorInternalServer(message string) Error {
	return &ErrorInternalServer{
		message: message,
	}
}

func (e *ErrorInternalServer) Message() string {
	return fmt.Sprintf("internal server error: %s", e.message)
}

func (e *ErrorInternalServer) Error() string {
	return e.Message()
}

func (e *ErrorInternalServer) Code() int {
	return http.StatusInternalServerError
}
