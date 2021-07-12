package handlers

import (
	"net/http"
)

type ApiError interface {
	ResponseStatus() int
	ResponseMessage() string
	ResponseError() string
}

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func (e *apiError) ResponseStatus() int {
	return e.Status
}
func (e *apiError) ResponseMessage() string {
	return e.Message
}
func (e *apiError) ResponseError() string {
	return e.Error
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{Status: statusCode,
		Message: message}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{Status: http.StatusBadRequest,
		Message: message}
}
