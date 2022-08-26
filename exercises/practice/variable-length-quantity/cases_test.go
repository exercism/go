package variablelengthquantity

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

var encodeTestCases = []struct {
	description string
	input       []uint32
	expected    []byte
}{
	{
		description: "zero",
		input:       []uint32{0x0},
		expected:    []byte{0x0},
	},
	{
		description: "arbitrary single byte",
		input:       []uint32{0x40},
		expected:    []byte{0x40},
	},
	{
		description: "largest single byte",
		input:       []uint32{0x7f},
		expected:    []byte{0x7f},
	},
	{
		description: "smallest double byte",
		input:       []uint32{0x80},
		expected:    []byte{0x81, 0x0},
	},
	{
		description: "arbitrary double byte",
		input:       []uint32{0x2000},
		expected:    []byte{0xc0, 0x0},
	},
	{
		description: "largest double byte",
		input:       []uint32{0x3fff},
		expected:    []byte{0xff, 0x7f},
	},
	{
		description: "smallest triple byte",
		input:       []uint32{0x4000},
		expected:    []byte{0x81, 0x80, 0x0},
	},
	{
		description: "arbitrary triple byte",
		input:       []uint32{0x100000},
		expected:    []byte{0xc0, 0x80, 0x0},
	},
	{
		description: "largest triple byte",
		input:       []uint32{0x1fffff},
		expected:    []byte{0xff, 0xff, 0x7f},
	},
	{
		description: "smallest quadruple byte",
		input:       []uint32{0x200000},
		expected:    []byte{0x81, 0x80, 0x80, 0x0},
	},
	{
		description: "arbitrary quadruple byte",
		input:       []uint32{0x8000000},
		expected:    []byte{0xc0, 0x80, 0x80, 0x0},
	},
	{
		description: "largest quadruple byte",
		input:       []uint32{0xfffffff},
		expected:    []byte{0xff, 0xff, 0xff, 0x7f},
	},
	{
		description: "smallest quintuple byte",
		input:       []uint32{0x10000000},
		expected:    []byte{0x81, 0x80, 0x80, 0x80, 0x0},
	},
	{
		description: "arbitrary quintuple byte",
		input:       []uint32{0xff000000},
		expected:    []byte{0x8f, 0xf8, 0x80, 0x80, 0x0},
	},
	{
		description: "maximum 32-bit integer input",
		input:       []uint32{0xffffffff},
		expected:    []byte{0x8f, 0xff, 0xff, 0xff, 0x7f},
	},
	{
		description: "two single-byte values",
		input:       []uint32{0x40, 0x7f},
		expected:    []byte{0x40, 0x7f},
	},
	{
		description: "two multi-byte values",
		input:       []uint32{0x4000, 0x123456},
		expected:    []byte{0x81, 0x80, 0x0, 0xc8, 0xe8, 0x56},
	},
	{
		description: "many multi-byte values",
		input:       []uint32{0x2000, 0x123456, 0xfffffff, 0x0, 0x3fff, 0x4000},
		expected:    []byte{0xc0, 0x0, 0xc8, 0xe8, 0x56, 0xff, 0xff, 0xff, 0x7f, 0x0, 0xff, 0x7f, 0x81, 0x80, 0x0},
	},
}

var decodeTestCases = []struct {
	description   string
	input         []byte
	expected      []uint32
	errorExpected bool
}{{
	description:   "one byte",
	input:         []byte{0x7f},
	expected:      []uint32{0x7f},
	errorExpected: false,
},
	{
		description:   "two bytes",
		input:         []byte{0xc0, 0x0},
		expected:      []uint32{0x2000},
		errorExpected: false,
	},
	{
		description:   "three bytes",
		input:         []byte{0xff, 0xff, 0x7f},
		expected:      []uint32{0x1fffff},
		errorExpected: false,
	},
	{
		description:   "four bytes",
		input:         []byte{0x81, 0x80, 0x80, 0x0},
		expected:      []uint32{0x200000},
		errorExpected: false,
	},
	{
		description:   "maximum 32-bit integer",
		input:         []byte{0x8f, 0xff, 0xff, 0xff, 0x7f},
		expected:      []uint32{0xffffffff},
		errorExpected: false,
	},
	{
		description:   "incomplete sequence causes error",
		input:         []byte{0xff},
		expected:      []uint32(nil),
		errorExpected: true,
	},
	{
		description:   "incomplete sequence causes error, even if value is zero",
		input:         []byte{0x80},
		expected:      []uint32(nil),
		errorExpected: true,
	},
	{
		description:   "multiple values",
		input:         []byte{0xc0, 0x0, 0xc8, 0xe8, 0x56, 0xff, 0xff, 0xff, 0x7f, 0x0, 0xff, 0x7f, 0x81, 0x80, 0x0},
		expected:      []uint32{0x2000, 0x123456, 0xfffffff, 0x0, 0x3fff, 0x4000},
		errorExpected: false,
	},
}
