package railfence

import "testing"

func testCases(
	name string,
	op func(string, int) string,
	cases []testCase, t *testing.T,
) {
	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := op(tc.message, tc.rails); actual != tc.expected {
				t.Fatalf("%s(%q,%d)\n got:%q\nwant:%q", name, tc.message, tc.rails, actual, tc.expected)
			}
		})
	}
}

func runBenchmark(op func(string, int) string, cases []testCase, b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range cases {
			op(test.message, test.rails)
		}
	}
}

func TestEncode(t *testing.T) { testCases("Encode", Encode, encodeTests, t) }
func TestDecode(t *testing.T) { testCases("Decode", Decode, decodeTests, t) }

func BenchmarkEncode(b *testing.B) { runBenchmark(Encode, encodeTests, b) }
func BenchmarkDecode(b *testing.B) { runBenchmark(Decode, decodeTests, b) }
