package errors

import (
	"fmt"
)

// DomainError represents a domain-level error
type DomainError struct {
	Code       string
	Message    string
	HTTPStatus int
	Err        error
}

func (e *DomainError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *DomainError) Unwrap() error {
	return e.Err
}
