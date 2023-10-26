package apperrors

type Error interface {
	error
	Message() string
	Code() int
}

type JSONError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
