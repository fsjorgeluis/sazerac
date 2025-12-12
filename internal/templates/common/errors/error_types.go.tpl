package errors

import (
	"fmt"
	"net/http"
)

var (
	ErrNotFound = &DomainError{
		Code:       "NOT_FOUND",
		Message:    "Resource not found",
		HTTPStatus: http.StatusNotFound,
	}

	ErrBadRequest = &DomainError{
		Code:       "BAD_REQUEST",
		Message:    "Invalid request",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInternalServer = &DomainError{
		Code:       "INTERNAL_SERVER_ERROR",
		Message:    "Internal server error",
		HTTPStatus: http.StatusInternalServerError,
	}

	ErrUnauthorized = &DomainError{
		Code:       "UNAUTHORIZED",
		Message:    "Unauthorized",
		HTTPStatus: http.StatusUnauthorized,
	}

	ErrForbidden = &DomainError{
		Code:       "FORBIDDEN",
		Message:    "Forbidden",
		HTTPStatus: http.StatusForbidden,
	}

	ErrConflict = &DomainError{
		Code:       "CONFLICT",
		Message:    "Resource already exists",
		HTTPStatus: http.StatusConflict,
	}

	ErrBadGateway = &DomainError{
		Code:       "BAD_GATEWAY",
		Message:    "Bad gateway",
		HTTPStatus: http.StatusBadGateway,
	}
)

// NewNotFoundError creates a not found error for a specific resource
func NewNotFoundError(resource string) *DomainError {
	return &DomainError{
		Code:       "NOT_FOUND",
		Message:    fmt.Sprintf("%s not found", resource),
		HTTPStatus: http.StatusNotFound,
	}
}

// NewValidationError creates a validation error
func NewValidationError(message string) *DomainError {
	return &DomainError{
		Code:       "VALIDATION_ERROR",
		Message:    message,
		HTTPStatus: http.StatusBadRequest,
	}
}
