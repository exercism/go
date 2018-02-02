package variablelengthquantity

// Source: exercism/problem-specifications
// Commit: 48dc5e6 variable-length-quantity: Apply new "input" policy
// Problem Specifications Version: 1.1.0

// Encode a series of integers, producing a series of bytes.
var encodeTestCases = []struct {
	description string
	input       []uint32
	output      []byte
}{
	{
		"zero",
		[]uint32{0x0},
		[]byte{0x0},
	},
	{
		"arbitrary single byte",
		[]uint32{0x40},
		[]byte{0x40},
	},
	{
		"largest single byte",
		[]uint32{0x7f},
		[]byte{0x7f},
	},
	{
		"smallest double byte",
		[]uint32{0x80},
		[]byte{0x81, 0x0},
	},
	{
		"arbitrary double byte",
		[]uint32{0x2000},
		[]byte{0xc0, 0x0},
	},
	{
		"largest double byte",
		[]uint32{0x3fff},
		[]byte{0xff, 0x7f},
	},
	{
		"smallest triple byte",
		[]uint32{0x4000},
		[]byte{0x81, 0x80, 0x0},
	},
	{
		"arbitrary triple byte",
		[]uint32{0x100000},
		[]byte{0xc0, 0x80, 0x0},
	},
	{
		"largest triple byte",
		[]uint32{0x1fffff},
		[]byte{0xff, 0xff, 0x7f},
	},
	{
		"smallest quadruple byte",
		[]uint32{0x200000},
		[]byte{0x81, 0x80, 0x80, 0x0},
	},
	{
		"arbitrary quadruple byte",
		[]uint32{0x8000000},
		[]byte{0xc0, 0x80, 0x80, 0x0},
	},
	{
		"largest quadruple byte",
		[]uint32{0xfffffff},
		[]byte{0xff, 0xff, 0xff, 0x7f},
	},
	{
		"smallest quintuple byte",
		[]uint32{0x10000000},
		[]byte{0x81, 0x80, 0x80, 0x80, 0x0},
	},
	{
		"arbitrary quintuple byte",
		[]uint32{0xff000000},
		[]byte{0x8f, 0xf8, 0x80, 0x80, 0x0},
	},
	{
		"maximum 32-bit integer input",
		[]uint32{0xffffffff},
		[]byte{0x8f, 0xff, 0xff, 0xff, 0x7f},
	},
	{
		"two single-byte values",
		[]uint32{0x40, 0x7f},
		[]byte{0x40, 0x7f},
	},
	{
		"two multi-byte values",
		[]uint32{0x4000, 0x123456},
		[]byte{0x81, 0x80, 0x0, 0xc8, 0xe8, 0x56},
	},
	{
		"many multi-byte values",
		[]uint32{0x2000, 0x123456, 0xfffffff, 0x0, 0x3fff, 0x4000},
		[]byte{0xc0, 0x0, 0xc8, 0xe8, 0x56, 0xff, 0xff, 0xff, 0x7f, 0x0, 0xff, 0x7f, 0x81, 0x80, 0x0},
	},
}

// Decode a series of bytes, producing a series of integers.
var decodeTestCases = []struct {
	description string
	input       []byte
	output      []uint32 // nil slice indicates error expected.
}{

	{
		"one byte",
		[]byte{0x7f},
		[]uint32{0x7f},
	},
	{
		"two bytes",
		[]byte{0xc0, 0x0},
		[]uint32{0x2000},
	},
	{
		"three bytes",
		[]byte{0xff, 0xff, 0x7f},
		[]uint32{0x1fffff},
	},
	{
		"four bytes",
		[]byte{0x81, 0x80, 0x80, 0x0},
		[]uint32{0x200000},
	},
	{
		"maximum 32-bit integer",
		[]byte{0x8f, 0xff, 0xff, 0xff, 0x7f},
		[]uint32{0xffffffff},
	},
	{
		"incomplete sequence causes error",
		[]byte{0xff},
		[]uint32(nil),
	},
	{
		"incomplete sequence causes error, even if value is zero",
		[]byte{0x80},
		[]uint32(nil),
	},
	{
		"multiple values",
		[]byte{0xc0, 0x0, 0xc8, 0xe8, 0x56, 0xff, 0xff, 0xff, 0x7f, 0x0, 0xff, 0x7f, 0x81, 0x80, 0x0},
		[]uint32{0x2000, 0x123456, 0xfffffff, 0x0, 0x3fff, 0x4000},
	},
}
