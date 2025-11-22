package main

import (
	"../../../../gen"
	"log"
	"strconv"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"empty":        &[]testCase{},
		"contains":     &[]testCase{},
		"subset":       &[]testCase{},
		"disjoint":     &[]testCase{},
		"equal":        &[]testCase{},
		"add":          &[]testCase{},
		"intersection": &[]testCase{},
		"difference":   &[]testCase{},
		"union":        &[]testCase{},
	}
	if err := gen.Gen("custom-set", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Set     []int `json:"set"`     // "empty"/"contains"/"add" cases
		Set1    []int `json:"set1"`    // "subset"/"disjoint"/"equal"/"difference"/"intersection"/"union" cases
		Set2    []int `json:"set2"`    // "subset"/"disjoint"/"equal"/"difference"/"intersection"/"union" cases
		Element int   `json:"element"` // "contains"/"add" cases
	}
	Expected any `json:"expected"` // bool or []int
}

// There's some extra complexity because canonical-data.json uses integers, but we are using strings.
// It just seemed like strings would make a better and more practical example.

// str converts an integer 1..26 to a letter 'a'..'z'.
func str(n int) string {
	if n >= 1 && n <= 26 {
		return string(rune(n) + 'a' - 1)
	}
	return strconv.Itoa(n)
}

// getStringSet converts a slice of integers to a slice of strings (integers are converted to characters using str)
func getStringSet(in []int) []string {
	out := make([]string, 0)
	for _, v := range in {
		out = append(out, str(v))
	}
	return out
}

func (t testCase) GetExpectedBool() bool {
	b, ok := t.Expected.(bool)
	if !ok {
		log.Fatal("[ERROR] expected value for `expected` to be of type bool")
	}
	return b
}

func (t testCase) GetExpectedList() []string {
	v, ok := t.Expected.([]any)
	if !ok {
		log.Fatal("[ERROR] invalid type for field `expected`")
	}
	var numbers []int
	for _, v2 := range v {
		number, ok := v2.(float64)
		if !ok {
			log.Fatal("[ERROR] invalid type for field `expected`")
		}
		numbers = append(numbers, int(number))
	}
	return getStringSet(numbers)
}

func (t testCase) GetSet() []string {
	return getStringSet(t.Input.Set)
}

func (t testCase) GetSet1() []string {
	return getStringSet(t.Input.Set1)
}

func (t testCase) GetSet2() []string {
	return getStringSet(t.Input.Set2)
}

func (t testCase) GetElement() string {
	return str(t.Input.Element)
}

var tmpl = `package stringset

{{.Header}}

type (
	// unary function, bool result (IsEmpty)
	unaryBoolCase struct {
		description string
		set         []string
		want    	bool
	}
	// set-element function, bool result (Has)
	eleBoolCase struct {
		description string
		set         []string
		element         string
		want        bool
	}
	// binary function, bool result (Subset, Disjoint, Equal)
	binBoolCase struct {
		description string
		set1        []string
		set2        []string
		want        bool
	}
	// set-element operator (Add)
	eleOpCase struct {
		description string
		set         []string
		element         string
		want        []string
	}
	// set-set operator (Intersection, Difference, Union)
	binOpCase struct {
		description string
		set1        []string
		set2        []string
		want        []string
	}
)

// Returns true if the set contains no elements
var emptyCases = []unaryBoolCase{
	{{range .J.empty}} {
		description: 	{{printf "%q" .Description}},
		set:  		 	{{printf "%#v" .GetSet}},
		want: 			{{printf "%v"  .GetExpectedBool}},
	},
	{{end}}
}

// Sets can report if they contain an element
var containsCases = []eleBoolCase{
	{{range .J.contains}} {
		description: 	{{printf "%q"  .Description}},
		set:  		 	{{printf "%#v" .GetSet}},
		element:        {{printf "%q"  .GetElement}},
		want: 			{{printf "%v"  .GetExpectedBool}},
	},
	{{end}}
}

// A set is a subset if all of its elements are contained in the other set
var subsetCases = []binBoolCase{
	{{range .J.subset}} {
		description: 	{{printf "%q"   .Description}},
		set1:  		 	{{printf "%#v"  .GetSet1}},
		set2:        	{{printf "%#v"  .GetSet2}},
		want: 			{{printf "%v"   .GetExpectedBool}},
	},
	{{end}}
}

// Sets are disjoint if they share no elements
var disjointCases = []binBoolCase{
	{{range .J.disjoint}} {
		description: 	{{printf "%q"   .Description}},
		set1:  		 	{{printf "%#v"  .GetSet1}},
		set2:        	{{printf "%#v"  .GetSet2}},
		want: 			{{printf "%v"   .GetExpectedBool}},
	},
	{{end}}
}

// Sets with the same elements are equal
var equalCases = []binBoolCase{
	{{range .J.equal}} {
		description: 	{{printf "%q"   .Description}},
		set1:  		 	{{printf "%#v"  .GetSet1}},
		set2:        	{{printf "%#v"  .GetSet2}},
		want: 			{{printf "%v"   .GetExpectedBool}},
	},
	{{end}}
}

// Unique elements can be added to a set
var addCases = []eleOpCase{
	{{range .J.add}} {
		description: 	{{printf "%q"  .Description}},
		set:  		 	{{printf "%#v" .GetSet}},
		element:        {{printf "%q"  .GetElement}},
		want: 			{{printf "%#v"  .GetExpectedList}},
	},
{{end}}
}

// Intersection returns a set of all shared elements
var intersectionCases = []binOpCase{
	{{range .J.intersection}} {
		description: 	{{printf "%q"  .Description}},
		set1:  		 	{{printf "%#v" .GetSet1}},
		set2:        	{{printf "%#v"  .GetSet2}},
		want: 			{{printf "%#v"  .GetExpectedList}},
	},
	{{end}}
}

// Difference (or Complement) of a set is a set of all elements that are only in the first set
var differenceCases = []binOpCase{
	{{range .J.difference}} {
		description: 	{{printf "%q"  .Description}},
		set1:  		 	{{printf "%#v" .GetSet1}},
		set2:        	{{printf "%#v"  .GetSet2}},
		want: 			{{printf "%#v"  .GetExpectedList}},
	},
	{{end}}
}

// Union returns a set of all elements in either set
var unionCases = []binOpCase{
	{{range .J.union}} {
		description: 	{{printf "%q"  .Description}},
		set1:  		 	{{printf "%#v" .GetSet1}},
		set2:        	{{printf "%#v"  .GetSet2}},
		want: 			{{printf "%#v"  .GetExpectedList}},
	},
	{{end}}
}
`
