package phonenumber

import (
	"testing"
)

const targetTestVersion = 2

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestNumber(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := Number(test.input)
		if !test.expectErr {
			if actualErr != nil {
				// if we don't expect an error and there is one
				var _ error = actualErr
				t.Errorf("Number(%s): expected no error, but error is: %s", test.input, actualErr)
			}
			if actual != test.number {
				t.Errorf("Number(%s): expected [%s], actual: [%s]", test.input, test.number, actual)
			}
		} else if actualErr == nil {
			// if we expect an error and there isn't one
			t.Errorf("Number(%s): expected an error, but error is nil", test.input)
		}
	}
}

func BenchmarkNumber(b *testing.B) {
	b.StopTimer()
	for _, test := range numberTests {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			Number(test.input)
		}
		b.StopTimer()
	}
}

func TestAreaCode(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := AreaCode(test.input)
		if !test.expectErr {
			if actual != test.areaCode {
				t.Errorf("AreaCode(%s): expected [%s], actual: [%s]", test.input, test.areaCode, actual)
			}
			if actualErr != nil {
				// if we don't expect an error and there is one
				var _ error = actualErr
				t.Errorf("AreaCode(%s): expected no error, but error is: %s", test.input, actualErr)
			}
		} else if actualErr == nil {
			// if we expect an error and there isn't one
			t.Errorf("AreaCode(%s): expected an error, but error is nil", test.input)
		}
	}
}

func BenchmarkAreaCode(b *testing.B) {
	b.StopTimer()
	for _, test := range numberTests {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			AreaCode(test.input)
		}
		b.StopTimer()
	}
}

func TestFormat(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := Format(test.input)
		if !test.expectErr {
			if actualErr != nil {
				// if we don't expect an error and there is one
				var _ error = actualErr
				t.Errorf("Format(%s): expected no error, but error is: %s", test.input, actualErr)
			}
			if actual != test.formatted {
				t.Errorf("Format(%s): expected [%s], actual: [%s]", test.input, test.formatted, actual)
			}
		} else if actualErr == nil {
			// if we expect an error and there isn't one
			t.Errorf("Format(%s): expected an error, but error is nil", test.input)
		}
	}
}

func BenchmarkFormat(b *testing.B) {
	b.StopTimer()
	for _, test := range numberTests {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			Format(test.input)
		}
		b.StopTimer()
	}
}
