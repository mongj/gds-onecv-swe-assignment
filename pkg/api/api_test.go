package api

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildError(t *testing.T) {
	errMessage := "test error"
	err := errors.New(errMessage)

	builtErr := BuildError(err)

	require.Equal(t, errMessage, builtErr.Message)
	require.Equal(t, GenericResponse{Message: errMessage}, builtErr)
}
