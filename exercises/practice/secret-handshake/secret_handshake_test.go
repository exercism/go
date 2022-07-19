package secret

import (
	"reflect"
	"testing"
)

func TestHandshake(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			actual := Handshake(tc.input)
			// use len() to allow either nil or empty list, because they are not equal by DeepEqual
			if len(actual) == 0 && len(tc.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("Handshake(%d) = %q, want: %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkHandshake(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for j := uint(0); j < 32; j++ {
			Handshake(j)
		}
	}
}
