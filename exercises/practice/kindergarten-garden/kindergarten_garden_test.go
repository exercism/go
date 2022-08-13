package kindergarten

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

type lookup struct {
	child  string
	plants []string
	ok     bool
}

type gardenTest struct {
	number    int
	diagram   string
	children  []string
	errorHint string
	lookups   []lookup
}

var tests = []gardenTest{
	{
		number:    1,
		diagram:   "\nRC\nGG",
		children:  []string{"Alice"},
		errorHint: "",
		lookups:   []lookup{{"Alice", []string{"radishes", "clover", "grass", "grass"}, true}},
	},
	{
		number:    2,
		diagram:   "\nVC\nRC",
		children:  []string{"Alice"},
		errorHint: "",
		lookups:   []lookup{{"Alice", []string{"violets", "clover", "radishes", "clover"}, true}},
	},
	{
		number:    3,
		diagram:   "\nVVCG\nVVRC",
		children:  []string{"Alice", "Bob"},
		errorHint: "",
		lookups:   []lookup{{"Bob", []string{"clover", "grass", "radishes", "clover"}, true}},
	},
	{
		number:    4,
		diagram:   "\nVVCCGG\nVVCCGG",
		children:  []string{"Alice", "Bob", "Charlie"},
		errorHint: "",
		lookups:   []lookup{{"Bob", []string{"clover", "clover", "clover", "clover"}, true}, {"Charlie", []string{"grass", "grass", "grass", "grass"}, true}},
	},
	test5, // full garden test
	test6, // out of order names test
	{
		number:    7,
		diagram:   "\nRC\nGG",
		children:  []string{"Alice"},
		errorHint: "",
		lookups:   []lookup{{"Bob", []string{"radishes", "clover", "grass", "grass"}, false}}, // lookup invalid name
	},

	// failure tests
	{
		number:    8,
		diagram:   "RC\nGG",
		children:  []string{"Alice"},
		errorHint: "wrong diagram format",
		lookups:   nil,
	},
	{
		number:    9,
		diagram:   "\nRCCC\nGG",
		children:  []string{""},
		errorHint: "mismatched rows",
		lookups:   nil,
	},
	{
		number:    10,
		diagram:   "\nRCC\nGGC",
		children:  []string{"Alice"},
		errorHint: "odd number of cups",
		lookups:   nil,
	},
	{
		number:    11,
		diagram:   "\nRCCC\nGGCC",
		children:  []string{"Alice", "Alice"},
		errorHint: "duplicate name",
		lookups:   nil,
	},
	{
		number:    12,
		diagram:   "\nrc\ngg",
		children:  []string{"Alice"},
		errorHint: "invalid cup codes",
		lookups:   nil,
	},
}

// full garden test
var test5 = gardenTest{
	number:    5,
	diagram:   "\nVRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
	children:  []string{"Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"},
	errorHint: "",
	lookups: []lookup{
		{"Alice", []string{"violets", "radishes", "violets", "radishes"}, true},
		{"Bob", []string{"clover", "grass", "clover", "clover"}, true},
		{"Charlie", []string{"violets", "violets", "clover", "grass"}, true},
		{"David", []string{"radishes", "violets", "clover", "radishes"}, true},
		{"Eve", []string{"clover", "grass", "radishes", "grass"}, true},
		{"Fred", []string{"grass", "clover", "violets", "clover"}, true},
		{"Ginny", []string{"clover", "grass", "grass", "clover"}, true},
		{"Harriet", []string{"violets", "radishes", "radishes", "violets"}, true},
		{"Ileana", []string{"grass", "clover", "violets", "clover"}, true},
		{"Joseph", []string{"violets", "clover", "violets", "grass"}, true},
		{"Kincaid", []string{"grass", "clover", "clover", "grass"}, true},
		{"Larry", []string{"grass", "violets", "clover", "violets"}, true},
	}}

// out of order names test
var (
	test6names = []string{"Samantha", "Patricia", "Xander", "Roger"}
	test6      = gardenTest{
		number:    6,
		diagram:   "\nVCRRGVRG\nRVGCCGCV",
		children:  test6names,
		errorHint: "",
		lookups: []lookup{
			{"Patricia", []string{"violets", "clover", "radishes", "violets"}, true},
			{"Roger", []string{"radishes", "radishes", "grass", "clover"}, true},
			{"Samantha", []string{"grass", "violets", "clover", "grass"}, true},
			{"Xander", []string{"radishes", "grass", "clover", "violets"}, true},
		},
	}
)

func TestGarden(t *testing.T) {
	for _, test := range tests {
		t.Run(strconv.Itoa(test.number), func(t *testing.T) {
			actual, err := NewGarden(test.diagram, test.children)
			switch {
			case test.errorHint != "":
				if err == nil {
					t.Fatalf("Testcase %d for NewGarden expected error but got nil\nhint:%s", test.number, test.errorHint)
				}
			case err != nil:
				t.Fatalf("Testcase %d for NewGarden returned unexpected error: %v ", test.number, err)
			}
			for _, l := range test.lookups {
				switch plants, ok := actual.Plants(l.child); {
				case ok != l.ok:
					t.Fatalf("Testcase %d lookup %s returned ok = %t, want %t", test.number, l.child, ok, l.ok)
				case ok && !reflect.DeepEqual(plants, l.plants):
					t.Fatalf("Testcase %d lookup %s = %q, want: %q", test.number, l.child, plants, l.plants)
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
