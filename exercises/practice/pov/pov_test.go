package pov

import (
	"sort"
	"testing"
)

func TestNewNotNil(t *testing.T) {
	for _, treeName := range newValueChildrenTestTrees {
		t.Run(treeName+" not nil", func(t *testing.T) {
			tree := mkTestTree(treeName)
			if tree == nil {
				t.Fatalf("tree should not be nil: %v", treeName)
			}
		})
	}
}

func TestValue(t *testing.T) {
	for _, treeName := range newValueChildrenTestTrees {
		t.Run(treeName+" value", func(t *testing.T) {
			tree := mkTestTree(treeName)
			got := tree.Value()
			want := testTrees[treeName].root
			if want != got {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestChildren(t *testing.T) {
	for _, treeName := range newValueChildrenTestTrees {
		t.Run(treeName+" Children", func(t *testing.T) {
			tree := mkTestTree(treeName)
			got := tree.Children()
			want := testTrees[treeName].children
			if !treeSliceEqual(want, got) {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestFromPov(t *testing.T) {
	for _, tc := range fromPovTestCases {
		t.Run(tc.description, func(t *testing.T) {
			tree := mkTestTree(tc.treeName)
			got := tree.FromPov(tc.from)
			want := tc.expected
			if !treeEqual(want, got) {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestPathTo(t *testing.T) {
	for _, tc := range pathToTestCases {
		t.Run(tc.description, func(t *testing.T) {
			tree := mkTestTree(tc.treeName)
			got := tree.PathTo(tc.from, tc.to)
			want := tc.expected
			if !stringSliceEqual(want, got) {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

var benchmarkResultPov *Tree

func BenchmarkFromPov(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	var result *Tree
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree := mkTestTree("complex tree with cousins")
		from := "x"
		result = tree.FromPov(from)
	}
	benchmarkResultPov = result
}

var benchmarkResultPathTo []string

func BenchmarkPathTo(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	var result []string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree := mkTestTree("complex tree with cousins")
		from := "x"
		to := "cousin-1"
		result = tree.PathTo(from, to)
	}
	benchmarkResultPathTo = result
}

func treeEqual(tr1, tr2 *Tree) bool {
	switch {
	case tr1 == nil && tr2 == nil:
		return true
	case tr1 == nil && tr2 != nil:
		return false
	case tr1 != nil && tr2 == nil:
		return false
	default:
		return tr1.Value() == tr2.Value() && treeSliceEqual(tr1.Children(), tr2.Children())
	}
}

func treeSliceEqual(trs1, trs2 []*Tree) bool {
	// allows permutation of children
	if len(trs1) != len(trs2) {
		return false
	}
	if len(trs1) == 0 && len(trs2) == 0 {
		return true
	}
	sortByValue := func(xs []*Tree) func(int, int) bool {
		return func(i, j int) bool {
			return xs[i].Value() < xs[j].Value()
		}
	}
	sort.Slice(trs1, sortByValue(trs1))
	sort.Slice(trs2, sortByValue(trs2))
	for i := range trs1 {
		if !treeEqual(trs1[i], trs2[i]) {
			return false
		}
	}
	return true
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
