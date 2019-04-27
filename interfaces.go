package pantopoda

import (
	"encoding/json"
	"fmt"
	"strings"

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

// RequestHeaders represents the key-value pairs in an HTTP header.
type RequestHeaders map[string]string

// RequestBody represent the json body in an HTTP request body.
type RequestBody map[string]interface{}

// QueryParams represent url query params.
type QueryParams map[string][]string

// ToString converts QueryParams map to its string representation.
func (q QueryParams) ToString() string {
	outSlice := make([]string, 0)

	for key, value := range q {
		if len(value) > 1 {
			for _, v := range value {
				outSlice = append(outSlice, fmt.Sprintf("%s[]=%s", key, v))
			}
		} else {
			outSlice = append(outSlice, fmt.Sprintf("%s=%s", key, value))
		}
	}

	return strings.Join(outSlice, "&")
}

func (q QueryParams) Empty() bool {
	return len(q) == 0
}

// Request represent all data such as payload, query params, and header of a
// JSON HTTP request call.
type Request struct {
	// Payload represent json body of HTTP call.
	Payload RequestBody

	// Query represent query params of HTTP call endpoint.
	Query QueryParams

	// Headers represent headers of HTTP call.
	Headers RequestHeaders
}

// Response represents HTTP call response body.
type Response []byte

// Unmarshal parses the JSON-encoded response and stores the result in the value
// pointed to by v.
func (r Response) Unmarshal(v interface{}) error {
	return json.Unmarshal(r, v)
}

// ToString convert the response body to its string value.
func (r Response) ToString() string {
	return string(r)
}
