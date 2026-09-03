//go:build LANCE && cgo

package states

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldTryNextLanceVectorType(t *testing.T) {
	runtimeError := errors.New("read field column failed")
	require.False(t, shouldTryNextLanceVectorType(runtimeError))

	schemaError := markLanceSchemaCandidateError(errors.New("field type mismatch"))
	require.True(t, shouldTryNextLanceVectorType(schemaError))
}
