package phonenumber

import (
	"testing"
)

func TestNumber(t *testing.T) {
	runTests("Number", Number, func(tc testCase) string { return tc.expectedNumber }, t)
}

func BenchmarkNumber(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Number(test.input)
		}
	}
}

func TestAreaCode(t *testing.T) {
	runTests("AreaCode", AreaCode, func(tc testCase) string { return tc.expectedAreaCode }, t)
}

func BenchmarkAreaCode(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			AreaCode(test.input)
		}
	}
}

func TestFormat(t *testing.T) {
	runTests("Format", Format, func(tc testCase) string { return tc.expectedFormatted }, t)
}

func BenchmarkFormat(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Format(test.input)
		}
	}
}

func runTests(
	funcName string,
	f func(s string) (string, error),
	getExpected func(tc testCase) string,
	t *testing.T,
) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, actualErr := f(tc.input)
			switch {
			case tc.expectErr:
				if actualErr == nil {
					t.Fatalf("%s(%q) expected error, got: %q", funcName, tc.input, actual)
				}
			case actualErr != nil:
				t.Fatalf("%s(%q) returned error: %v, want: %q", funcName, tc.input, actualErr, getExpected(tc))
			case actual != getExpected(tc):
				t.Fatalf("%s(%q) = %q, want: %q", funcName, tc.input, actual, getExpected(tc))
			}
		})
	}
}
