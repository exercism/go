package lsproduct

// Source: exercism/problem-specifications
// Commit: 85da7a5 largest-series-product: testdata to have standard error indicators (#1331)
// Problem Specifications Version: 1.2.0

var tests = []struct {
	digits  string
	span    int
	product int64
	ok      bool
	error   string
}{
	{"29", 2, 18, true, ""},
	{"0123456789", 2, 72, true, ""},
	{"576802143", 2, 48, true, ""},
	{"0123456789", 3, 504, true, ""},
	{"1027839564", 3, 270, true, ""},
	{"0123456789", 5, 15120, true, ""},
	{"73167176531330624919225119674426574742355349194934", 6, 23520, true, ""},
	{"0000", 2, 0, true, ""},
	{"99099", 3, 0, true, ""},
	{"123", 4, -1, false, "span must be smaller than string length"},
	{"", 0, 1, true, ""},
	{"123", 0, 1, true, ""},
	{"", 1, -1, false, "span must be smaller than string length"},
	{"1234a5", 2, -1, false, "digits input must only contain digits"},
	{"12345", -1, -1, false, "span must be greater than zero"},
}
