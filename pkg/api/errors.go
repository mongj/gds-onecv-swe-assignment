package api

import "fmt"

var (
	MissingInputErrorStr    = "missing input in request body"
	JSONParseErrorStr       = "failed to parse JSON request body"
	StudentNotFoundErrorStr = "no student found with email"
)

// ErrMissingInput is takes as arguments two strings indicating the expected input
// and the actual input, and builds an error with the customized message
func ErrMissingInput(expected string, actual string) error {
	return fmt.Errorf("%s: expected %s, got %s", MissingInputErrorStr, expected, actual)
}

// Custom error wrappers to provide more
// context to the error message
func WrapError(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func WrapErrorJSONRequestBody(err error) error {
	return WrapError(err, JSONParseErrorStr)
}
