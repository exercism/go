package allyourbase

import "testing"

const targetTestVersion = 1

// Note: ConvertToBase should accept leading zeroes in its input,
// but never emit leading zeroes in its output.
// Exception: If the value of the output is zero, represent it with a single zero.
var testCases = []struct {
	description  string
	inputBase    uint64
	outputBase   uint64
	inputDigits  []uint64
	outputDigits []uint64
	error        error
}{
	// converted from x-common by hand
	// cases with negative inputs are excluded, as Go uses unsigned ints.
	// x-common version: 1.0.0
	{
		description:  "single bit one to decimal",
		inputBase:    2,
		inputDigits:  []uint64{1},
		outputBase:   10,
		outputDigits: []uint64{1},
	},
	{
		description:  "binary to single decimal",
		inputBase:    2,
		inputDigits:  []uint64{1, 0, 1},
		outputBase:   10,
		outputDigits: []uint64{5},
	},
	{
		description:  "single decimal to binary",
		inputBase:    10,
		inputDigits:  []uint64{5},
		outputBase:   2,
		outputDigits: []uint64{1, 0, 1},
	},
	{
		description:  "binary to multiple decimal",
		inputBase:    2,
		inputDigits:  []uint64{1, 0, 1, 0, 1, 0},
		outputBase:   10,
		outputDigits: []uint64{4, 2},
	},
	{
		description:  "decimal to binary",
		inputBase:    10,
		inputDigits:  []uint64{4, 2},
		outputBase:   2,
		outputDigits: []uint64{1, 0, 1, 0, 1, 0},
	},
	{
		description:  "trinary to hexadecimal",
		inputBase:    3,
		inputDigits:  []uint64{1, 1, 2, 0},
		outputBase:   16,
		outputDigits: []uint64{2, 10},
	},
	{
		description:  "hexadecimal to trinary",
		inputBase:    16,
		inputDigits:  []uint64{2, 10},
		outputBase:   3,
		outputDigits: []uint64{1, 1, 2, 0},
	},
	{
		description:  "15-bit integer",
		inputBase:    97,
		inputDigits:  []uint64{3, 46, 60},
		outputBase:   73,
		outputDigits: []uint64{6, 10, 45},
	},
	{
		description:  "empty list",
		inputBase:    2,
		inputDigits:  []uint64{},
		outputBase:   10,
		outputDigits: []uint64{0},
		error:        nil,
	},
	{
		description:  "single zero",
		inputBase:    10,
		inputDigits:  []uint64{0},
		outputBase:   2,
		outputDigits: []uint64{0},
	},
	{
		description:  "multiple zeros",
		inputBase:    10,
		inputDigits:  []uint64{0, 0, 0},
		outputBase:   2,
		outputDigits: []uint64{0},
	},
	{
		description:  "leading zeros",
		inputBase:    7,
		inputDigits:  []uint64{0, 6, 0},
		outputBase:   10,
		outputDigits: []uint64{4, 2},
	},
	{
		description:  "invalid positive digit",
		inputBase:    2,
		inputDigits:  []uint64{1, 2, 1, 0, 1, 0},
		outputBase:   10,
		outputDigits: nil,
		error:        ErrInvalidDigit,
	},
	{
		description:  "first base is one",
		inputBase:    1,
		inputDigits:  []uint64{},
		outputBase:   10,
		outputDigits: nil,
		error:        ErrInvalidBase,
	},
	{
		description:  "second base is one",
		inputBase:    2,
		inputDigits:  []uint64{1, 0, 1, 0, 1, 0},
		outputBase:   1,
		outputDigits: nil,
		error:        ErrInvalidBase,
	},
	{
		description:  "first base is zero",
		inputBase:    0,
		inputDigits:  []uint64{},
		outputBase:   10,
		outputDigits: nil,
		error:        ErrInvalidBase,
	},
	{
		description:  "second base is zero",
		inputBase:    10,
		inputDigits:  []uint64{7},
		outputBase:   0,
		outputDigits: nil,
		error:        ErrInvalidBase,
	},
}

func digitsEqual(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestConvertToBase(t *testing.T) {
	for _, c := range testCases {
		output, err := ConvertToBase(c.inputBase, c.inputDigits, c.outputBase)
		if err != c.error {
			t.Fatalf(`FAIL: %s
	Expected error: %v
	Got: %v`, c.description, c.error, err)
		}

		if !digitsEqual(c.outputDigits, output) {
			t.Fatalf(`FAIL: %s
    Input base: %d
    Input digits: %v
    Output base: %d
    Expected output digits: %v
    Got: %v`, c.description, c.inputBase, c.inputDigits, c.outputBase, c.outputDigits, output)
		} else {
			t.Logf("PASS: %s", c.description)
		}
	}
}
