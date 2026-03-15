package gocounting

import (
	"cmp"
	"fmt"
	"slices"
	"testing"
)

func cmpPair(a, b [2]int) int {
	if a[0] != b[0] {
		return cmp.Compare(a[0], b[0])
	}
	return cmp.Compare(a[1], b[1])
}

func territoryEqual(a, b [][2]int) bool {
	if (a == nil || len(a) == 0) && (b == nil || len(b) == 0) {
		return true
	}
	a = slices.Clone(a)
	b = slices.Clone(b)
	slices.SortFunc(a, cmpPair)
	slices.SortFunc(b, cmpPair)
	return slices.Equal(a, b)
}

func TestTerritory(t *testing.T) {
	for _, tc := range oneTerritoryTestCases {
		call := fmt.Sprintf("NewGame(%#v).Territory(%d, %d)", tc.board, tc.posX, tc.posY)
		t.Run(tc.description, func(t *testing.T) {
			if owner, territory, err := NewGame(tc.board).Territory(tc.posX, tc.posY); err != nil {
				if tc.expectedErr == "" {
					t.Fatalf("%s unexpected returned error %v", call, err)
				} else if err.Error() != tc.expectedErr {
					t.Fatalf("%s returned error %v, want error %s", call, err, tc.expectedErr)
				}
			} else if owner != tc.owner || !territoryEqual(territory, tc.territory) {
				t.Fatalf("%s\ngot:  (%q, %#v, %v),\nwant: (%q, %#v, <nil>)", call, owner, territory, err, tc.owner, tc.territory)
			}
		})
	}
}

func TestTerritories(t *testing.T) {
	for _, tc := range allTerritoriesTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := NewGame(tc.input).Territories(); !(slices.Equal(actual.Black, tc.expected.Black) && slices.Equal(actual.White, tc.expected.White) && slices.Equal(actual.None, tc.expected.None)) {
				t.Fatalf("NewGame(%#v).Territories()\ngot:  %#v,\nwant: %#v", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkTerritory(b *testing.B) {
	for range b.N {
		for _, tc := range oneTerritoryTestCases {
			NewGame(tc.board).Territory(tc.posX, tc.posY)
		}
	}
}

func BenchmarkTerritories(b *testing.B) {
	for range b.N {
		for _, tc := range allTerritoriesTestCases {
			NewGame(tc.input).Territories()
		}
	}
}
