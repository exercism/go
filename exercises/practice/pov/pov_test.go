package pov

import (
	"reflect"
	"sort"
	"testing"
)

type arc struct{ fr, to string }

type testCase struct {
	description string
	leaves      []string
	arcPairs    []arc
	root        string
	arcStrings  []string
	reRooted    []string
}

var testCases = []testCase{
	{
		description: "singleton",
		leaves:      []string{"x"},
		arcPairs:    nil,
		root:        "x",
		arcStrings:  nil,
		reRooted:    nil,
	},
	{
		description: "simple tree",
		leaves:      []string{"sibling", "x"},
		arcPairs: []arc{
			{"parent", "sibling"},
			{"parent", "x"},
		},
		root: "parent",
		arcStrings: []string{
			"parent -> sibling",
			"parent -> x",
		},
		reRooted: []string{
			"parent -> sibling",
			"x -> parent",
		},
	},
	{
		description: "large flat",
		leaves:      []string{"sib-a", "sib-b", "x", "sib-c", "sib-d"},
		arcPairs: []arc{
			{"parent", "sib-a"},
			{"parent", "sib-b"},
			{"parent", "x"},
			{"parent", "sib-c"},
			{"parent", "sib-d"},
		},
		root: "parent",
		arcStrings: []string{
			"parent -> sib-a",
			"parent -> sib-b",
			"parent -> sib-c",
			"parent -> sib-d",
			"parent -> x",
		},
		reRooted: []string{
			"parent -> sib-a",
			"parent -> sib-b",
			"parent -> sib-c",
			"parent -> sib-d",
			"x -> parent",
		},
	},
	{
		description: "deeply nested",
		leaves:      []string{"x"},
		arcPairs: []arc{
			{"level-4", "x"},
			{"level-3", "level-4"},
			{"level-2", "level-3"},
			{"level-1", "level-2"},
			{"level-0", "level-1"},
		},
		root: "level-0",
		arcStrings: []string{
			"level-0 -> level-1",
			"level-1 -> level-2",
			"level-2 -> level-3",
			"level-3 -> level-4",
			"level-4 -> x",
		},
		reRooted: []string{
			"level-1 -> level-0",
			"level-2 -> level-1",
			"level-3 -> level-2",
			"level-4 -> level-3",
			"x -> level-4",
		},
	},
	{
		description: "cousins",
		leaves:      []string{"sib-1", "x", "sib-2", "cousin-1", "cousin-2"},
		arcPairs: []arc{
			{"parent", "sib-1"},
			{"parent", "x"},
			{"parent", "sib-2"},
			{"aunt", "cousin-1"},
			{"aunt", "cousin-2"},
			{"grand-parent", "parent"},
			{"grand-parent", "aunt"},
		},
		root: "grand-parent",
		arcStrings: []string{
			"aunt -> cousin-1",
			"aunt -> cousin-2",
			"grand-parent -> aunt",
			"grand-parent -> parent",
			"parent -> sib-1",
			"parent -> sib-2",
			"parent -> x",
		},
		reRooted: []string{
			"aunt -> cousin-1",
			"aunt -> cousin-2",
			"grand-parent -> aunt",
			"parent -> grand-parent",
			"parent -> sib-1",
			"parent -> sib-2",
			"x -> parent",
		},
	},
	{
		description: "target with children",
		leaves: []string{"child-1", "child-2", "nephew", "niece",
			"2nd-cousin-1", "2nd-cousin-2", "2nd-cousin-3", "2nd-cousin-4"},
		arcPairs: []arc{
			{"x", "child-1"},
			{"x", "child-2"},
			{"sibling", "nephew"},
			{"sibling", "niece"},
			{"cousin-1", "2nd-cousin-1"},
			{"cousin-1", "2nd-cousin-2"},
			{"cousin-2", "2nd-cousin-3"},
			{"cousin-2", "2nd-cousin-4"},
			{"parent", "x"},
			{"parent", "sibling"},
			{"aunt", "cousin-1"},
			{"aunt", "cousin-2"},
			{"grand-parent", "parent"},
			{"grand-parent", "aunt"},
		},
		root: "grand-parent",
		arcStrings: []string{
			"aunt -> cousin-1",
			"aunt -> cousin-2",
			"cousin-1 -> 2nd-cousin-1",
			"cousin-1 -> 2nd-cousin-2",
			"cousin-2 -> 2nd-cousin-3",
			"cousin-2 -> 2nd-cousin-4",
			"grand-parent -> aunt",
			"grand-parent -> parent",
			"parent -> sibling",
			"parent -> x",
			"sibling -> nephew",
			"sibling -> niece",
			"x -> child-1",
			"x -> child-2",
		},
		reRooted: []string{
			"aunt -> cousin-1",
			"aunt -> cousin-2",
			"cousin-1 -> 2nd-cousin-1",
			"cousin-1 -> 2nd-cousin-2",
			"cousin-2 -> 2nd-cousin-3",
			"cousin-2 -> 2nd-cousin-4",
			"grand-parent -> aunt",
			"parent -> grand-parent",
			"parent -> sibling",
			"sibling -> nephew",
			"sibling -> niece",
			"x -> child-1",
			"x -> child-2",
			"x -> parent",
		},
	},
}

func (tc testCase) graph() *Graph {
	g := New()
	for _, l := range tc.leaves {
		g.AddNode(l)
	}
	for _, a := range tc.arcPairs {
		g.AddArc(a.fr, a.to)
	}
	return g
}

func (tc testCase) testResult(got, want []string, msg string, t *testing.T) {
	if len(got)+len(want) == 0 {
		return
	}
	gs := append([]string{}, got...)
	sort.Strings(gs)
	if reflect.DeepEqual(gs, want) {
		return
	}
	// test has failed
	t.Log(tc.description, "test case")
	t.Log(msg)
	t.Logf("got %d arcs:", len(got))
	for _, s := range got {
		t.Log(" ", s)
	}
	t.Logf("that result sorted:")
	for _, s := range gs {
		t.Log(" ", s)
	}
	t.Logf("want %d arcs:", len(want))
	for _, s := range want {
		t.Log(" ", s)
	}
	t.FailNow()
}

func TestConstruction(t *testing.T) {
	for _, tc := range testCases {
		got := tc.graph().ArcList()
		want := tc.arcStrings
		tc.testResult(got, want, "incorrect graph construction", t)
	}
}

func TestChangeRoot(t *testing.T) {
	for _, tc := range testCases {
		got := tc.graph().ChangeRoot(tc.root, "x").ArcList()
		want := tc.reRooted
		tc.testResult(got, want, "incorrect root change", t)
	}
}

func BenchmarkConstructOnlyNoChange(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			tc.graph()
		}
	}
}

func BenchmarkConstructAndChangeRoot(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			tc.graph().ChangeRoot(tc.root, "x")
		}
	}
}
