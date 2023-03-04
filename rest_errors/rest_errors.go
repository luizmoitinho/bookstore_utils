package rest_errors

import (
	"errors"
	"net/http"
)

const (
	DATABASE_ERROR = "database error"
)

type RestError struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewError(message string) error {
	return errors.New(message)
}

func NewUnauthorizedError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) *RestError {
	result := &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}

	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}
