package panicrecover

import (
	"runtime/debug"
	"testing"
)

type PanicTestFunc func()

func TestAddPanic(t *testing.T) {

	t.Run("AddPanicTest", func(t *testing.T) {
		result, msg, _ := didPanic(AddPanic)
		if result != true {
			t.Errorf(msg.(string))
		}
	})
}

func TestRecoverPanic(t *testing.T) {
}

func didPanic(f PanicTestFunc) (didPanic bool, message interface{}, stack string) {
	didPanic = true

	defer func() {
		message = recover()
		if didPanic {
			stack = string(debug.Stack())
		}
	}()

	// call the target function
	f()
	didPanic = false

	return
}
