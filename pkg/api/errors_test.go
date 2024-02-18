package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrMissingInput(t *testing.T) {
	err := ErrMissingInput("foo", "bar")
	require.EqualError(t, err, MissingInputErrorStr+": expected foo, got bar")
}

func TestWrapError(t *testing.T) {
	err := WrapError(fmt.Errorf("ERROR"), "test error")
	require.EqualError(t, err, "test error: ERROR")
}

func TestWrapErrorJSONRequestBody(t *testing.T) {
	err := WrapErrorJSONRequestBody(fmt.Errorf("JSON ERROR"))
	require.EqualError(t, err, JSONParseErrorStr+": JSON ERROR")
}
