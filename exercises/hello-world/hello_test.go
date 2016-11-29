package greeting

import "testing"

// Define a function HelloWorld(string) string.
//
// Also define a testVersion with a value that matches
// the targetTestVersion here.

const targetTestVersion = 3

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

	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}
