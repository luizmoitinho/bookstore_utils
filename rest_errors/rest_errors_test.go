package rest_errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	err := NewError("this a message error")

	assert.NotNil(t, err)
	assert.EqualValues(t, "this a message error", err.Error())
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("this is an unauthorized message")

	assert.NotNil(t, err)
	assert.EqualValues(t, 401, err.Status)
	assert.EqualValues(t, "this is an unauthorized message", err.Message)
	assert.EqualValues(t, "unauthorized", err.Error)
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is a bad request message")

	assert.NotNil(t, err)
	assert.EqualValues(t, 400, err.Status)
	assert.EqualValues(t, "this is a bad request message", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is a not found message")

	assert.NotNil(t, err)
	assert.EqualValues(t, 404, err.Status)
	assert.EqualValues(t, "this is a not found message", err.Message)
	assert.EqualValues(t, "not_found", err.Error)
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is a internal server message", errors.New("database error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, 500, err.Status)
	assert.EqualValues(t, "this is a internal server message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0])
}
