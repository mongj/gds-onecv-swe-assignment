package testutil

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// test package contains some utility functions and types for testing

type TestTable map[string]struct {
	URL                string
	JSONBody           []byte
	ExpectedStatus     int
	ExpectedBody       string
	ExpectedResponse   interface{}
	ExpectIsError      bool
	ExpectErrorMessage string
}

// ExecuteRequest executes a request and returns the response recorder
func ExecuteRequest(req *http.Request, h http.HandlerFunc) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h)
	handler.ServeHTTP(rr, req)

	return rr
}

// checkResponseCode checks the response code of the response
func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// CheckErrorMessage checks the error message of the response
// It only checks if the error message that we defined is in the response,
// not the whole error message which can be changed by the implementer
// of external libraries etc.
func CheckErrorMessage(t *testing.T, expected, actual string) {
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected error message to contain: '%s'. Got '%s'\n", expected, actual)
	}
}
