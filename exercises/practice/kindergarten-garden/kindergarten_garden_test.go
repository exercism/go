package kindergarten

import (
	"reflect"
	"sort"
	"testing"
)

type lookup struct {
	child  string
	plants []string
	ok     bool
}

type gardenTest struct {
	description string
	diagram     string
	children    []string
	expectError bool
	lookups     []lookup
}

var tests = []gardenTest{
	{
		description: "garden with single student",
		diagram:     "\nRC\nGG",
		children:    []string{"Alice"},
		expectError: false,
		lookups:     []lookup{{child: "Alice", plants: []string{"radishes", "clover", "grass", "grass"}, ok: true}},
	},
	{
		description: "different garden with single student",
		diagram:     "\nVC\nRC",
		children:    []string{"Alice"},
		expectError: false,
		lookups:     []lookup{{child: "Alice", plants: []string{"violets", "clover", "radishes", "clover"}, ok: true}},
	},
	{
		description: "garden with two students",
		diagram:     "\nVVCG\nVVRC",
		children:    []string{"Alice", "Bob"},
		expectError: false,
		lookups:     []lookup{{child: "Bob", plants: []string{"clover", "grass", "radishes", "clover"}, ok: true}},
	},
	{
		description: "garden with three students",
		diagram:     "\nVVCCGG\nVVCCGG",
		children:    []string{"Alice", "Bob", "Charlie"},
		expectError: false,
		lookups: []lookup{
			{child: "Bob", plants: []string{"clover", "clover", "clover", "clover"}, ok: true},
			{child: "Charlie", plants: []string{"grass", "grass", "grass", "grass"}, ok: true},
		},
	},
	test5,
	test6,
	{
		description: "lookup invalid name",
		diagram:     "\nRC\nGG",
		children:    []string{"Alice"},
		expectError: false,
		lookups:     []lookup{{child: "Bob", plants: []string{"radishes", "clover", "grass", "grass"}, ok: false}},
	},
	// failure tests
	{
		description: "wrong diagram format",
		diagram:     "RC\nGG",
		children:    []string{"Alice"},
		expectError: true,
		lookups:     nil,
	},
	{
		description: "mismatched rows",
		diagram:     "\nRCCC\nGG",
		children:    []string{""},
		expectError: true,
		lookups:     nil,
	},
	{
		description: "odd number of cups",
		diagram:     "\nRCC\nGGC",
		children:    []string{"Alice"},
		expectError: true,
		lookups:     nil,
	},
	{
		description: "duplicate name",
		diagram:     "\nRCCC\nGGCC",
		children:    []string{"Alice", "Alice"},
		expectError: true,
		lookups:     nil,
	},
	{
		description: "invalid cup codes",
		diagram:     "\nrc\ngg",
		children:    []string{"Alice"},
		expectError: true,
		lookups:     nil,
	},
}

var test5 = gardenTest{
	description: "full garden",
	diagram:     "\nVRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
	children:    []string{"Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"},
	expectError: false,
	lookups: []lookup{
		{child: "Alice", plants: []string{"violets", "radishes", "violets", "radishes"}, ok: true},
		{child: "Bob", plants: []string{"clover", "grass", "clover", "clover"}, ok: true},
		{child: "Charlie", plants: []string{"violets", "violets", "clover", "grass"}, ok: true},
		{child: "David", plants: []string{"radishes", "violets", "clover", "radishes"}, ok: true},
		{child: "Eve", plants: []string{"clover", "grass", "radishes", "grass"}, ok: true},
		{child: "Fred", plants: []string{"grass", "clover", "violets", "clover"}, ok: true},
		{child: "Ginny", plants: []string{"clover", "grass", "grass", "clover"}, ok: true},
		{child: "Harriet", plants: []string{"violets", "radishes", "radishes", "violets"}, ok: true},
		{child: "Ileana", plants: []string{"grass", "clover", "violets", "clover"}, ok: true},
		{child: "Joseph", plants: []string{"violets", "clover", "violets", "grass"}, ok: true},
		{child: "Kincaid", plants: []string{"grass", "clover", "clover", "grass"}, ok: true},
		{child: "Larry", plants: []string{"grass", "violets", "clover", "violets"}, ok: true},
	}}

var (
	test6names = []string{"Samantha", "Patricia", "Xander", "Roger"}
	test6      = gardenTest{
		description: "names out of order",
		diagram:     "\nVCRRGVRG\nRVGCCGCV",
		children:    test6names,
		expectError: false,
		lookups: []lookup{
			{child: "Patricia", plants: []string{"violets", "clover", "radishes", "violets"}, ok: true},
			{child: "Roger", plants: []string{"radishes", "radishes", "grass", "clover"}, ok: true},
			{child: "Samantha", plants: []string{"grass", "violets", "clover", "grass"}, ok: true},
			{child: "Xander", plants: []string{"radishes", "grass", "clover", "violets"}, ok: true},
		},
	}
)

func TestGarden(t *testing.T) {
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actual, err := NewGarden(test.diagram, test.children)
			switch {
			case test.expectError:
				if err == nil {
					t.Fatal("NewGarden expected error but got nil")
				}
			case err != nil:
				t.Fatalf("NewGarden returned unexpected error: %v ", err)
			}
			for _, l := range test.lookups {
				switch plants, ok := actual.Plants(l.child); {
				case ok != l.ok:
					t.Fatalf("Lookup %s returned ok = %t, want %t", l.child, ok, l.ok)
				case ok && !reflect.DeepEqual(plants, l.plants):
					t.Fatalf("Lookup %s = %q, want: %q", l.child, plants, l.plants)
				}
			}
		})
	}
}

// The lazy way to meet the alphabetizing requirement is with sort.Strings
// on the argument slice.  That's an in-place sort though and it's bad practice
// to have a side effect.
func TestNamesNotModified(t *testing.T) {
	cp := append([]string{}, test6names...)
	_, err := NewGarden(test6.diagram, cp)
	if err != nil || sort.StringsAreSorted(cp) {
		t.Fatalf("error in test setup: TestNamesNotModified requires valid garden and unsorted children")
	}
	if !reflect.DeepEqual(cp, test6names) {
		t.Fatalf("NewGarden modified children argment. Arguments should not be modified.")
	}
}

// A test taken from the Ruby tests.  It checks that Garden objects
// are self-contained and do not rely on package variables.
func TestTwoGardens(t *testing.T) {
	diagram := `
VCRRGVRG
RVGCCGCV`
	g1, err1 := NewGarden(diagram, []string{"Alice", "Bob", "Charlie", "Dan"})
	g2, err2 := NewGarden(diagram, []string{"Bob", "Charlie", "Dan", "Erin"})
	if err1 != nil || err2 != nil {
		t.Fatalf("error in test setup: Two garden test needs valid gardens")
	}
	tf := func(g *Garden, n int, child string, expPlants []string) {
		switch plants, ok := g.Plants(child); {
		case !ok:
			t.Fatalf("error in test setup: Garden %d lookup %s returned ok = false, want true. Check if the child exists in the garden", n, child)
		case !reflect.DeepEqual(plants, expPlants):
			t.Fatalf("Garden %d lookup %s = %q, want %q.",
				n, child, plants, expPlants)
		}
	}
	tf(g1, 1, "Bob", []string{"radishes", "radishes", "grass", "clover"})
	tf(g2, 2, "Bob", []string{"violets", "clover", "radishes", "violets"})
	tf(g1, 1, "Charlie", []string{"grass", "violets", "clover", "grass"})
	tf(g2, 2, "Charlie", []string{"radishes", "radishes", "grass", "clover"})
}

func BenchmarkNewGarden(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			NewGarden(test.diagram, test.children)
		}
	}
}

func BenchmarkGarden_Plants(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	g, err := NewGarden(test5.diagram, test5.children)
	if err != nil {
		b.Fatalf("error in benchmark setup: BenchmarkGarden_Plants requires valid garden")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, l := range test5.lookups {
			g.Plants(l.child)
		}
	}
}
