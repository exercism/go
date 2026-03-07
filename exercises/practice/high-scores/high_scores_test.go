package highscores

import (
	"slices"
	"testing"
)

func TestSquareRoot(t *testing.T) {
	for _, tc := range intTestCases {
		t.Run(tc.description, func(t *testing.T) {
			h := NewHighScores(tc.input)
			if actual := tc.call(h); actual != tc.expected {
				t.Fatalf("h = NewHighScores(%#v); %s = %d, want: %d", tc.input, tc.callMsg, actual, tc.expected)
			}
		})
	}
	for _, tc := range sliceTestCases {
		t.Run(tc.description, func(t *testing.T) {
			h := NewHighScores(tc.input)
			if actual := tc.call(h); !slices.Equal(actual, tc.expected) {
				t.Fatalf("h = NewHighScores(%#v); %s = %#v, want: %#v", tc.input, tc.callMsg, actual, tc.expected)
			}
		})
	}
}

func BenchmarkHighScores(b *testing.B) {
	for range b.N {
		for _, tc := range intTestCases {
			tc.call(NewHighScores(tc.input))
		}
		for _, tc := range sliceTestCases {
			tc.call(NewHighScores(tc.input))
		}
	}
}
