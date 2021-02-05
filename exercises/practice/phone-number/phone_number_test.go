package phonenumber

import (
	"testing"
)

func TestNumber(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := Number(test.input)
		if !test.expectErr {
			if actualErr != nil {
				// if we don't expect an error and there is one
				var _ error = actualErr
				t.Errorf("FAIL: %s\nNumber(%q): expected no error, but error is: %s", test.description, test.input, actualErr)
			}
			if actual != test.number {
				t.Errorf("FAIL: %s\nNumber(%q): expected [%s], actual: [%s]", test.description, test.input, test.number, actual)
			}
		} else if actualErr == nil {
			// if we expect an error and there isn't one
			t.Errorf("FAIL: %s\nNumber(%q): expected an error, but error is nil", test.description, test.input)
		}
	}
}

func BenchmarkNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range numberTests {
			Number(test.input)
		}
	}
}

func TestAreaCode(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := AreaCode(test.input)
		if !test.expectErr {
			if actual != test.areaCode {
				t.Errorf("FAIL: %s\nAreaCode(%q): expected [%s], actual: [%s]", test.description, test.input, test.areaCode, actual)
			}
			if actualErr != nil {
				// if we don't expect an error and there is one
				var _ error = actualErr
				t.Errorf("FAIL: %s\nAreaCode(%q): expected no error, but error is: %s", test.description, test.input, actualErr)
			}
		} else if actualErr == nil {
			// if we expect an error and there isn't one
			t.Errorf("FAIL: %s\nAreaCode(%q): expected an error, but error is nil", test.description, test.input)
		}
	}
}

func BenchmarkAreaCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range numberTests {
			AreaCode(test.input)
		}
	}
}

func TestFormat(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := Format(test.input)
		if !test.expectErr {
			if actualErr != nil {
				// if we don't expect an error and there is one
				var _ error = actualErr
				t.Errorf("FAIL: %s\nFormat(%q): expected no error, but error is: %s", test.description, test.input, actualErr)
			}
			if actual != test.formatted {
				t.Errorf("FAIL: %s\nFormat(%q): expected [%s], actual: [%s]", test.description, test.input, test.formatted, actual)
			}
		} else if actualErr == nil {
			// if we expect an error and there isn't one
			t.Errorf("FAIL: %s\nFormat(%q): expected an error, but error is nil", test.description, test.input)
		}
	}
}

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range numberTests {
			Format(test.input)
		}
	}
}
