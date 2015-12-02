package helloworld

import "testing"

// Define a function HelloWorld() string.
//
// Also define an exported TestVersion with a value that matches
// the internal testVersion here.

const testVersion = 1

func TestHelloWorld(t *testing.T) {
	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
	}
	expected := "Hello, World!"
	observed := HelloWorld()
	if observed != expected {
		t.Fatalf("HelloWorld() = %v, want %v", observed, expected)
	}
}
