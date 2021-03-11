package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_panic(t *testing.T) {
	assert.True(t, assert.PanicsWithError(t, PANICERROR.Error(), DefaultPanic))
}

func Test_panicHandler(t *testing.T) {
	assert.True(t, assert.NotPanics(t, func() {
		defer PanicHandler()
		DefaultPanic()
	}))
}

func Test_panicHandler_withWarpError(t *testing.T) {
	assert.True(t, assert.NotPanics(t, func() {
		defer PanicHandler()
		ErrInvalidUser := NewDefWarpError("invalid user", nil)

		err1 := NewDefWarpError("wrap err1", ErrInvalidUser)
		err2 := NewDefWarpError("wrap err2", err1)
		Panic(err2)
	}))
}
