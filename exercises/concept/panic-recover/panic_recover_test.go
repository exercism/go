package panicrecover

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"testing"
)

type PanicTestFunc func([]string, int) string

func TestRecoverPanic(t *testing.T) {

	t.Run("AddPanicTest", func(t *testing.T) {
		didFuncPanic, _, _ := didPanic(AccessNamesPanic, t)
		if didFuncPanic == true {
			t.Errorf("Expected: AddPanic method should Panic, Result: AddPanic method did not Panic")
		}
	})
}

func TestRecoverPanicWithError(t *testing.T) {

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
			val := AccessNamesPanic(tt.name, tt.index)
			typeOfValue := reflect.TypeOf(val)
			fmt.Println(typeOfValue)
		})
	}
}

func TestResolveError(t *testing.T) {
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
			val := ResolveError(tt.name, tt.index)
			if val != tt.name[2] {
				t.Errorf("Expected: RecoverPanic() method should have recovered panic created by AddPanic() method. Result: Panic Not Recovered")
			}
		})
	}
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
			_ = f(tt.name, tt.index)
			didPanic = false
		})

	}
	return
}
