package relativedistance

import "testing"

func TestDegreeOfSeparation(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual, ok := DegreeOfSeparation(tc.tree, tc.personA, tc.personB); ok != tc.expectedOk {
				t.Fatalf("DegreeOfSeparation(%#v), ok = %t, want: ok = %t", tc.tree, ok, tc.expectedOk)
			} else if actual != tc.expectedDegree {
				t.Fatalf("DegreeOfSeparation(%#v) = %d, want: %d", tc.tree, actual, tc.expectedDegree)
			}
		})
	}
}

func BenchmarkDegreeOfSeparation(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			DegreeOfSeparation(tc.tree, tc.personA, tc.personB)
		}
	}
}
