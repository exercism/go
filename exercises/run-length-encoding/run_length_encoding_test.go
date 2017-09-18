package run_length_encoding

import "testing"

func TestRunLengthEncode(t *testing.T) {
	for _, test := range encodeTests {
		if actual := RunLengthEncode(test.input); actual != test.expected {
			t.Errorf("RunLengthEncode(%s) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}
func TestRunLengthDecode(t *testing.T) {
	for _, test := range decodeTests {
		if actual := RunLengthDecode(test.input); actual != test.expected {
			t.Errorf("RunLengthDecode(%s) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}
func TestRunLengthEncodeDecode(t *testing.T) {
	for _, test := range encodeDecodeTests {
		if actual := RunLengthDecode(RunLengthEncode(test.input)); actual != test.expected {
			t.Errorf("RunLengthDecode(RunLengthEncode(%s)) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}