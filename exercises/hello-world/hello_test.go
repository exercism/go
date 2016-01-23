package hello

import "testing"

// Define a function HelloWorld(string) string.
//
// Also define an exported TestVersion with a value that matches
// the internal testVersion here.

const testVersion = 1

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name, expected string
	}{
		{"", "Hello, World!"},
		{"Gopher", "Hello, Gopher!"},
	}
	for _, test := range tests {
		observed := HelloWorld(test.name)
		if observed != test.expected {
			t.Fatalf("HelloWorld(%s) = %v, want %v", test.name, observed, test.expected)
		}
	}

	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
	}
}
