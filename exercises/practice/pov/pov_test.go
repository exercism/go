package pov

import (
	"sort"
	"testing"
)

func TestNewNotNil(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
	}{
		{
			name: "singleton",
			tree: New("x"),
		},
		{
			name: "parent and one sibling",
			tree: New("parent", New("x"), New("sibling")),
		},
		{
			name: "parent and kids",
			tree: New("parent", New("x", New("kid-0"), New("kid-1"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.tree == nil {
				t.Fatalf("tree should not be nil")
			}
		})
	}
}

func TestValue(t *testing.T) {
	tests := []struct {
		name     string
		root     string
		children []*Tree
	}{
		{
			name:     "singleton",
			root:     "x",
			children: nil,
		},
		{
			name:     "parent and one sibling",
			root:     "parent",
			children: []*Tree{New("x"), New("sibling")},
		},
		{
			name:     "parent and kids",
			root:     "parent",
			children: []*Tree{New("x", New("kid-0"), New("kid-1"))},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := New(tt.root, tt.children...)
			got := tree.Value()
			want := tt.root
			if want != got {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestChildren(t *testing.T) {
	tests := []struct {
		name     string
		root     string
		children []*Tree
	}{
		{
			name:     "singleton",
			root:     "x",
			children: nil,
		},
		{
			name:     "parent and one sibling",
			root:     "parent",
			children: []*Tree{New("x"), New("sibling")},
		},
		{
			name:     "parent and kids",
			root:     "parent",
			children: []*Tree{New("x", New("kid-0"), New("kid-1"))},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := New(tt.root, tt.children...)
			got := tree.Children()
			want := tt.children
			if !treeSliceEqual(want, got) {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestFromPov(t *testing.T) {
	tests := []struct {
		description string
		tree        *Tree
		from        string
		expected    *Tree
	}{
		{
			description: "Results in the same tree if the input tree is a singleton",
			tree:        New("x"),
			from:        "x",
			expected:    New("x"),
		},
		{
			description: "Can reroot a tree with a parent and one sibling",
			tree:        New("parent", New("x"), New("sibling")),
			from:        "x",
			expected:    New("x", New("parent", New("sibling"))),
		},
		{
			description: "Can reroot a tree with a parent and many siblings",
			tree:        New("parent", New("a"), New("x"), New("b"), New("c")),
			from:        "x",
			expected:    New("x", New("parent", New("a"), New("b"), New("c"))),
		},
		{
			description: "Can reroot a tree with new root deeply nested in tree",
			tree:        New("level-0", New("level-1", New("level-2", New("level-3", New("x"))))),
			from:        "x",
			expected:    New("x", New("level-3", New("level-2", New("level-1", New("level-0"))))),
		},
		{
			description: "Moves children of the new root to same level as former parent",
			tree:        New("parent", New("x", New("kid-0"), New("kid-1"))),
			from:        "x",
			expected:    New("x", New("kid-0"), New("kid-1"), New("parent")),
		},
		{
			description: "Can reroot a complex tree with cousins",
			tree: New("grandparent", New("parent",
				New("x", New("kid-0"), New("kid-1")), New("sibling-0"),
				New("sibling-1")), New("uncle", New("cousin-0"), New("cousin-1"))),
			from: "x",
			expected: New("x", New("kid-0"), New("kid-1"),
				New("parent", New("sibling-0"), New("sibling-1"),
					New("grandparent", New("uncle", New("cousin-0"), New("cousin-1"))))),
		},
		{
			description: "Errors if target does not exist in a singleton tree",
			tree:        New("x"),
			from:        "nonexistent",
			expected:    nil,
		},
		{
			description: "Errors if target does not exist in a large tree",
			tree: New("parent",
				New("x", New("kid-0"), New("kid-1")), New("sibling-0"), New("sibling-1")),
			from:     "nonexistent",
			expected: nil,
		},
	}
	for _, tt := range tests {
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
	tests := []struct {
		description string
		tree        *Tree
		from        string
		to          string
		expected    []string
	}{
		{
			description: "Can find path to parent",
			tree:        New("parent", New("x"), New("sibling")),
			from:        "x",
			to:          "parent",
			expected:    []string{"x", "parent"},
		},
		{
			description: "Can find path to sibling",
			tree:        New("parent", New("a"), New("x"), New("b"), New("c")),
			from:        "x",
			to:          "b",
			expected:    []string{"x", "parent", "b"},
		},
		{
			description: "Can find path to cousin",
			tree: New("grandparent", New("parent",
				New("x", New("kid-0"), New("kid-1")), New("sibling-0"),
				New("sibling-1")), New("uncle", New("cousin-0"), New("cousin-1"))),
			from:     "x",
			to:       "cousin-1",
			expected: []string{"x", "parent", "grandparent", "uncle", "cousin-1"},
		},
		{
			description: "Can find path not involving root",
			tree:        New("grandparent", New("parent", New("x"), New("sibling-0"), New("sibling-1"))),
			from:        "x",
			to:          "sibling-1",
			expected:    []string{"x", "parent", "sibling-1"},
		},
		{
			description: "Can find path from nodes other than x",
			tree:        New("parent", New("a"), New("x"), New("b"), New("c")),
			from:        "a",
			to:          "c",
			expected:    []string{"a", "parent", "c"},
		},
		{
			description: "Errors if destination does not exist",
			tree:        New("parent", New("x", New("kid-0"), New("kid-1")), New("sibling-0"), New("sibling-1")),
			from:        "x",
			to:          "nonexistent",
			expected:    nil,
		},
		{
			description: "Errors if source does not exist",
			tree:        New("parent", New("x", New("kid-0"), New("kid-1")), New("sibling-0"), New("sibling-1")),
			from:        "nonexistent",
			to:          "x",
			expected:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.tree.PathTo(tt.from, tt.to)
			want := tt.expected
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
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	var result []string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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
