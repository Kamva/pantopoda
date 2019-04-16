package pantopoda

import (
	"github.com/Kamva/nautilus"
	"github.com/Kamva/shark"
)

// RequestData is an interface for incoming request payload
type RequestData interface {
	nautilus.Taggable

	// Validate runs request data validation and returns validation error if
	Validate() ValidationError
}

// ErrorType is a string subtype used inside ValidationError type which
// determines the type of validation error. There are two type of validation
// errors; BadRequest, that returned when the request payload did not match
// with request data structure. And RuleViolation, that returned when any of
// specified rules in validation tag violated in incoming request payload.
type ErrorType string

// BadRequest  that returned when the request payload did not match with
// request data structure
const BadRequest ErrorType = "bad_request"

// RuleViolation that returned when any of specified rules in validation tags
// violated in incoming request payload.
const RuleViolation ErrorType = "validation_failed"

// ValidationError is an error type containing validation error bag and error
// type determine type of validation error.
type ValidationError struct {
	// ErrorBag contains any validation errors on request fields.
	ErrorBag shark.ErrorBag

	// ErrorType is a string determine type of the validation error.
	ErrorType ErrorType
}
