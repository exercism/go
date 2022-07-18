package encode

import "testing"

func TestRunLengthEncode(t *testing.T) {
	for _, tc := range encodeTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := RunLengthEncode(tc.input); actual != tc.expected {
				t.Errorf("RunLengthEncode(%q) = %q, want:%q", tc.input, actual, tc.expected)
			}
		})
	}
}
func TestRunLengthDecode(t *testing.T) {
	for _, tc := range decodeTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := RunLengthDecode(tc.input); actual != tc.expected {
				t.Errorf("RunLengthDecode(%q) = %q, want:%q", tc.input, actual, tc.expected)
			}
		})
	}
}
func TestRunLengthEncodeDecode(t *testing.T) {
	for _, tc := range encodeDecodeTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := RunLengthDecode(RunLengthEncode(tc.input)); actual != tc.expected {
				t.Errorf("RunLengthDecode(RunLengthEncode(%q)) = %q, want:%q", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkRunLengthEncode(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range encodeTests {
			RunLengthEncode(test.input)
		}
	}
}

func BenchmarkRunLengthDecode(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range decodeTests {
			RunLengthDecode(test.input)
		}
	}
}
