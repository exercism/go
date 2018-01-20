package grains

import (
	"testing"
)

func TestSquare(t *testing.T) {
	for _, test := range squareTests {
		actualVal, actualErr := Square(test.input)

		// check actualVal only if no error expected
		if !test.expectError && actualVal != test.expectedVal {
			t.Fatalf("FAIL: %s\nSquare(%d) expected %d, Actual %d", test.description, test.input, test.expectedVal, actualVal)
		}

		// if we expect an error and there isn't one
		if test.expectError && actualErr == nil {
			t.Fatalf("FAIL: %s\nSquare(%d) expected an error, but error is nil", test.description, test.input)
		}
		// if we don't expect an error and there is one
		if !test.expectError && actualErr != nil {
			var _ error = actualErr
			t.Fatalf("FAIL: %s\nSquare(%d) expected no error, but error is: %s", test.description, test.input, actualErr)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func TestTotal(t *testing.T) {
	var expected uint64 = 18446744073709551615
	if actual := Total(); actual != expected {
		t.Errorf("Total() expected %d, Actual %d", expected, actual)
	}
}

func BenchmarkSquare(b *testing.B) {

	for i := 0; i < b.N; i++ {

		for _, test := range squareTests {
			Square(test.input)
		}

	}
}

func BenchmarkTotal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Total()
	}
}
