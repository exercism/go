package greeting

import "testing"

// Define a function HelloWorld(string) string.

// In any exercise solution using these tests you must also define a
// testVersion constant with a value that matches the targetTestVersion
// here.
const targetTestVersion = 3

// TestTestVersion is used to ensure the exercise solution is considered
// compatable with the rest of the tests. It should be the initial testing
// function in any test suite. If there is an incompatibility the tests can
// be considered unreliable therefore this function will abort the testing.
func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name, expected string
	}{
		{"", "Hello, World!"},
		{"Gopher", "Hello, Gopher!"},
		{"ゴーファー", "Hello, ゴーファー!"},
	}
	for _, test := range tests {
		if observed := HelloWorld(test.name); observed != test.expected {
			t.Fatalf("HelloWorld(%s) = %v, want %v", test.name, observed, test.expected)
		}
	}
}
