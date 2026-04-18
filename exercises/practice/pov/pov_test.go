package pov

import (
	"slices"
	"sort"
	"testing"
)

var helperTestCases = []struct {
	description string
	tree        *Tree
	root        string
	children    []*Tree
}{
	{
		description: "singleton",
		tree:        New("x"),
		root:        "x",
		children:    nil,
	},
	{
		description: "parent and one sibling",
		tree:        New("parent", New("x"), New("sibling")),
		root:        "parent",
		children:    []*Tree{New("x"), New("sibling")},
	},
	{
		description: "parent and kids",
		tree:        New("parent", New("x", New("kid-0"), New("kid-1"))),
		root:        "parent",
		children:    []*Tree{New("x", New("kid-0"), New("kid-1"))},
	},
}

func TestNewNotNil(t *testing.T) {
	for _, tt := range helperTestCases {
		t.Run(tt.description, func(t *testing.T) {
			if tt.tree == nil {
				t.Fatalf("tree should not be nil")
			}
		})
	}
}

func TestValue(t *testing.T) {
	for _, tt := range helperTestCases {
		t.Run(tt.description, func(t *testing.T) {
			tree := tt.tree
			got := tree.Value()
			want := tt.root
			if want != got {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestChildren(t *testing.T) {
	for _, tt := range helperTestCases {
		t.Run(tt.description, func(t *testing.T) {
			tree := tt.tree
			got := tree.Children()
			want := tt.children
			if !treeSliceEqual(want, got) {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestFromPov(t *testing.T) {
	for _, tt := range povTestCases {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.tree.FromPov(tt.from)
			want := tt.expected
			if !treeEqual(want, got) {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestPathTo(t *testing.T) {
	for _, tt := range pathTestCases {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.tree.PathTo(tt.from, tt.to)
			want := tt.expected
			if !slices.Equal(want, got) {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

var benchmarkResultPov *Tree

func BenchmarkFromPov(b *testing.B) {
	var result *Tree
	b.ResetTimer()
	for range b.N {
		tree := New("grandparent", New("parent",
			New("x", New("kid-0"), New("kid-1")), New("sibling-0"),
			New("sibling-1")), New("uncle", New("cousin-0"), New("cousin-1")))
		from := "x"
		result = tree.FromPov(from)
	}
	benchmarkResultPov = result
}

var benchmarkResultPathTo []string

func BenchmarkPathTo(b *testing.B) {
	var result []string
	b.ResetTimer()
	for range b.N {
		tree := New("grandparent", New("parent",
			New("x", New("kid-0"), New("kid-1")), New("sibling-0"),
			New("sibling-1")), New("uncle", New("cousin-0"), New("cousin-1")))
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
