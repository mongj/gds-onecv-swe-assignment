package api

import "fmt"

// Custom errors
func ErrMissingInput(expected string, actual string) error {
	return fmt.Errorf("missing input in request body: expected %s, got %s", expected, actual)
}

// Custom error wrappers to provide more
// context to the error message
func WrapError(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func WrapErrorJSONRequestBody(err error) error {
	return WrapError(err, "failed to parse JSON request body")
}
