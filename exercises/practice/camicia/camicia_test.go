package camicia

import "testing"

func (a Outcome) Equals(b Outcome) bool {
	return a.finishes == b.finishes && a.cards == b.cards && a.tricks == b.tricks
}

func TestSimulateGame(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := SimulateGame(tc.playerA, tc.playerB); !actual.Equals(tc.expected) {
				t.Fatalf("SimulateGame(%#v, %#v) =\n%#v\nWant:\n%#v", tc.playerA, tc.playerB, actual, tc.expected)
			}
		})
	}
}

func BenchmarkSimulateGame(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			SimulateGame(tc.playerA, tc.playerB)
		}
	}
}
