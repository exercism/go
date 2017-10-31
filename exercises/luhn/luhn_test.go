package luhn

import "testing"

func TestValid(t *testing.T) {
	for _, test := range testCases {
		if ok := Valid(test.input); ok != test.ok {
			t.Fatalf("Valid(%s): %s\n\t Expected: %t\n\t Got: %t", test.input, test.description, test.ok, ok)
		}
	}
}

func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Valid("2323 2005 7766 3554")
	}
}
