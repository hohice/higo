package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefError(t *testing.T) {
	err1 := &DefineError{"this is a define error type"}
	err2 := fmt.Errorf("wrap err2: %w\n", err1)
	err3 := fmt.Errorf("wrap err3: %w\n", err2)
	var err4 *DefineError
	assert.True(t, errors.As(err3, &err4))
}

func Test_DefWarpError(t *testing.T) {
	ErrInvalidUser := NewDefWarpError("invalid user", nil)

	err1 := NewDefWarpError("wrap err1", ErrInvalidUser)
	err2 := NewDefWarpError("wrap err2", err1)

	ErrInvalidUser1 := NewDefWarpError("invalid user", nil)

	assert.True(t, errors.Is(err2, ErrInvalidUser))
	assert.False(t, errors.Is(err2, ErrInvalidUser1))
	assert.True(t, errors.As(err2, &ErrInvalidUser1))

	assert.Equal(t, 3, len(err2.Stack()))
}
