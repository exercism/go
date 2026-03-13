package piecingittogether

import "testing"

func Equal(a, b PuzzleDetails) bool {
	equal := a.Pieces == b.Pieces && a.Border == b.Border && a.Inside == b.Inside
	equal = equal && a.Rows == b.Rows && a.Columns == b.Columns
	equal = equal && a.AspectRatio == b.AspectRatio && a.Format == b.Format
	return equal
}

func TestJigsawData(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual, err := JigsawData(PuzzleDetails(tc.input)); err != nil && !tc.wantError {
				t.Fatalf("JigsawData(%#v) returned error: %v, want: %#v", tc.input, err, tc.expected)
			} else if tc.wantError && err == nil {
				t.Fatalf("JigsawData(%#v) = %#v, want error %v", tc.input, actual, tc.err)
			} else if tc.wantError && err.Error() != tc.err {
				t.Fatalf("JigsawData(%#v) = error %v, want error %v", tc.input, err, tc.err)
			} else if tc.wantError && err.Error() == tc.err {
			} else if !Equal(actual, tc.expected) {
				t.Fatalf("JigsawData(%#v) = %#v, want: %#v", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkJigsawData(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			JigsawData(tc.input)
		}
	}
}
