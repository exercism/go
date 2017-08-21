package twofer

import "testing"

// Define a function ShareWith(string) string.

func TestShareWith(t *testing.T) {
	tests := []struct {
		name, expected string
	}{
		{"", "One for you, one for me."},
		{"Alice", "One for Alice, one for me."},
		{"Bob", "One for Bob, one for me."},
	}
	for _, test := range tests {
		if observed := ShareWith(test.name); observed != test.expected {
			t.Fatalf("ShareWith(%s) = %v, want %v", test.name, observed, test.expected)
		}
	}
}
