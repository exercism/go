package phonenumber

import (
	"testing"
)

const targetTestVersion = 1

type testCase struct {
	input     string
	expected  string
	expectErr bool
}

var numberTests = []testCase{
	{"(123) 456-7890", "1234567890", false},
	{"123.456.7890", "1234567890", false},
	{"1234567890", "1234567890", false},
	{"12345678901234567", "", true},
	{"21234567890", "", true},
	{"123456789", "", true},
}

func TestNumber(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
	for _, test := range numberTests {
		actual, actualErr := Number(test.input)
		if !test.expectErr {
			if actualErr != nil {
				// if we don't expect an error and there is one
				var _ error = actualErr
				t.Errorf("Number(%s): expected no error, but error is: %s", test.input, actualErr)
			}
			if actual != test.expected {
				t.Errorf("Number(%s): expected [%s], actual: [%s]", test.input, test.expected, actual)
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

var areaCodeTests = []testCase{
	{"1234567890", "123", false},
	{"213.456.7890", "213", false},
	{"213.456.7890.2345", "", true},
	{"213.456", "", true},
}

func TestAreaCode(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
	for _, test := range areaCodeTests {
		actual, actualErr := AreaCode(test.input)
		if !test.expectErr {
			if actual != test.expected {
				t.Errorf("AreaCode(%s): expected [%s], actual: [%s]", test.input, test.expected, actual)
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
	for _, test := range areaCodeTests {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			AreaCode(test.input)
		}
		b.StopTimer()
	}
}

var formatTests = []testCase{
	{"1234567890", "(123) 456-7890", false},
	{"11234567890", "(123) 456-7890", false},
	{"112345", "", true},
	{"11234590870986", "", true},
}

func TestFormat(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
	for _, test := range formatTests {
		actual, actualErr := Format(test.input)
		if !test.expectErr {
			if actualErr != nil {
				// if we don't expect an error and there is one
				var _ error = actualErr
				t.Errorf("Format(%s): expected no error, but error is: %s", test.input, actualErr)
			}
			if actual != test.expected {
				t.Errorf("Format(%s): expected [%s], actual: [%s]", test.input, test.expected, actual)
			}
		} else if actualErr == nil {
			// if we expect an error and there isn't one
			t.Errorf("Format(%s): expected an error, but error is nil", test.input)
		}
	}
}

func BenchmarkFormat(b *testing.B) {
	b.StopTimer()
	for _, test := range areaCodeTests {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			Format(test.input)
		}
		b.StopTimer()
	}
}
