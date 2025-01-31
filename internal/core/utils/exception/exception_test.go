package exception

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewException(t *testing.T) {
	e := New(http.StatusBadRequest, "Bad Request", nil)

	assert.Equal(t, http.StatusBadRequest, e.Code)
	assert.Equal(t, "Bad Request", e.Message)
	assert.Equal(t, "Bad Request", e.Error())
	assert.Nil(t, e.Cause)
	assert.Empty(t, e.Errors)
}

func TestNewExceptionWithError(t *testing.T) {
	e := New(http.StatusInternalServerError, "Internal Server Error", errors.New("some error"))

	assert.Equal(t, http.StatusInternalServerError, e.Code)
	assert.Equal(t, "Internal Server Error", e.Message)
	assert.Equal(t, "some error", e.Error())
	assert.Equal(t, "some error", e.Cause.Error())
	assert.Empty(t, e.Errors)
}

func TestNewExceptionWithOutCause(t *testing.T) {
	e := New(http.StatusForbidden, "Forbidden")

	assert.Equal(t, http.StatusForbidden, e.Code)
	assert.Equal(t, "Forbidden", e.Message)
	assert.Equal(t, "Forbidden", e.Error())
	assert.Nil(t, e.Cause)
	assert.Empty(t, e.Errors)
}

func TestIntoWithNil(t *testing.T) {
	e := Into(nil)

	assert.Nil(t, e)
}

func TestIntoWithException(t *testing.T) {
	e := New(http.StatusBadRequest, "Bad Request", nil)
	ee := Into(e)

	assert.Equal(t, e, ee)
}

func TestIntoWithError(t *testing.T) {
	e := errors.New("some error")
	ee := Into(e)

	assert.Equal(t, http.StatusInternalServerError, ee.Code)
	assert.Equal(t, "some error", ee.Message)
	assert.Equal(t, "some error", ee.Error())
	assert.Equal(t, "some error", ee.Cause.Error())
	assert.Empty(t, ee.Errors)
}

func TestAddErrors(t *testing.T) {
	e := New(http.StatusBadRequest, "Bad Request", nil)
	e.AddError("key", "error 1")
	e.AddError("key", "error 2")
	e.AddError("other", "error 3")

	assert.Equal(t, []string{"error 1", "error 2"}, e.Errors["key"])
	assert.Equal(t, []string{"error 3"}, e.Errors["other"])
}

func TestHasError(t *testing.T) {
	e := New(http.StatusBadRequest, "Bad Request", nil)
	e.AddError("key", "error 1")
	e.AddError("key", "error 2")
	e.AddError("other", "error 3")

	assert.True(t, e.HasError())
}

func TestHasNoError(t *testing.T) {
	e := New(http.StatusBadRequest, "Bad Request", nil)

	assert.False(t, e.HasError())
}
