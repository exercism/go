package panicrecover

import (
	"runtime/debug"
	"testing"
)

type PanicTestFunc func([]string, int)

func TestAddPanic(t *testing.T) {

	t.Run("AddPanicTest", func(t *testing.T) {
		didFuncPanic, _, _ := didPanic(AddPanic, t)
		if didFuncPanic == false {
			t.Errorf("Expected: AddPanic method should Panic, Result: AddPanic method did not Panic")
		}
	})
}

func TestRecoverPanic(t *testing.T) {

	t.Run("Test1", func(t *testing.T) {
		didFuncPanic, _, _ := didPanic(AddPanic, t)
		if didFuncPanic == true {
			t.Errorf("Expected: RecoverPanic() method should have recovered panic created by AddPanic() method. Result: Panic Not Recovered")
		}
	})
}

func didPanic(f PanicTestFunc, t *testing.T) (didPanic bool, message interface{}, stack string) {

	tests := []struct {
		testName string
		name     []string
		index    int
	}{
		{
			testName: "test1",
			name:     []string{"person1", "person2", "person3"},
			index:    3,
		},
		{
			testName: "test2",
			name:     []string{"person1", "person2", "person3"},
			index:    4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			didPanic = true

			defer func() {
				message = recover()
				if didPanic {
					stack = string(debug.Stack())
				}
			}()

			// call the target function
			f(tt.name, tt.index)
			didPanic = false
		})

	}
	return
}
