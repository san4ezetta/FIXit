package apierr

import (
	"errors"
	"fmt"
	"net/http"
)

type APIError struct {
	HTTPStatus int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Details    any    `json:"details,omitempty"`
	err        error
}

func (e *APIError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %v", e.Code, e.err)
	}
	return e.Code
}

func (e *APIError) Unwrap() error { return e.err }

func New(status int, code, msg string) *APIError {
	return &APIError{HTTPStatus: status, Code: code, Message: msg}
}

func Wrap(base *APIError, cause error, details ...any) *APIError {
	cp := *base
	cp.err = cause
	if len(details) > 0 {
		cp.Details = details[0]
	}
	return &cp
}

func WithDetails(base *APIError, details any) *APIError {
	cp := *base
	cp.Details = details
	return &cp
}

func From(err error) *APIError {
	var ae *APIError
	if errors.As(err, &ae) {
		return ae
	}
	return Wrap(ErrInternal, err)
}

var (
	ErrBadRequest      = New(http.StatusBadRequest, "bad_request", "Invalid request.")
	ErrInvalidJSON     = New(http.StatusBadRequest, "bad_request", "Invalid JSON.")
	ErrUnauthorized    = New(http.StatusUnauthorized, "unauthorized", "Unauthorized.")
	ErrUserNotFound    = New(http.StatusUnauthorized, "user_not_found", "User not found.")
	ErrInvalidPassword = New(http.StatusUnauthorized, "invalid_password", "Invalid password.")
	ErrTokenCreation   = New(http.StatusInternalServerError, "token_creation_failed", "Failed to create token.")
	ErrInternal        = New(http.StatusInternalServerError, "internal_error", "Something went wrong.")
)
