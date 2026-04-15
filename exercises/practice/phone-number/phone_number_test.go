package phonenumber

import (
	"testing"
)

func TestPhoneNumber(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			checkFunction(t, Number, "Number", tc.expectedNumber, tc)
			checkFunction(t, AreaCode, "AreaCode", tc.expectedAreaCode, tc)
			checkFunction(t, Format, "Format", tc.expectedFormatted, tc)
		})
	}
}

func checkFunction(
	t *testing.T,
	f func(string) (string, error),
	funcName string,
	want string,
	tc testCase,
) {
	got, gotErr := f(tc.input)
	switch {
	case tc.expectErr:
		if gotErr == nil {
			t.Errorf("%s(%q) expected error, got: %q", funcName, tc.input, got)
		}
	case gotErr != nil:
		t.Errorf("%s(%q) returned error: %v, want: %q", funcName, tc.input, gotErr, want)
	case got != want:
		t.Errorf("%s(%q) = %q, want: %q", funcName, tc.input, got, want)
	}
}

func BenchmarkNumber(b *testing.B) {
	for range b.N {
		for _, test := range testCases {
			Number(test.input)
		}
	}
}

func BenchmarkAreaCode(b *testing.B) {
	for range b.N {
		for _, test := range testCases {
			AreaCode(test.input)
		}
	}
}

func BenchmarkFormat(b *testing.B) {
	for range b.N {
		for _, test := range testCases {
			Format(test.input)
		}
	}
}
