package armstrong

import "testing"

func TestArmstrong(t *testing.T) {
	for _, tc := range testCases {
		if actual := IsNumber(tc.input); actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %v\nActual: %v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkIsNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, tc := range testCases {
			IsNumber(tc.input)
		}

	}
}
