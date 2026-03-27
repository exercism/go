package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"slices"
	"strings"
	"text/template"
)

var OpNames = map[string]string{
	"push":    "Push",
	"pop":     "Pop",
	"peek":    "Peek",
	"count":   "Size",
	"toList":  "Array",
	"reverse": "Reverse",
}

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"list": &[]testCase{},
	}
	if err := gen.Gen("simple-linked-list", j, t); err != nil {
		log.Fatal(err)
	}
}

type Operation struct {
	Operation string `json:"operation"`
	Expected  any    `json:"expected"`
	Value     int    `json:"value"`
}

func (o Operation) String() string {
	name, ok := OpNames[o.Operation]
	if !ok {
		panic("Invalid op " + o.Operation)
	}
	parts := []string{fmt.Sprintf("operation: %q", name)}
	if gen.IsError(o.Expected) {
		parts = append(parts, fmt.Sprintf("expectedErr: %q", gen.ErrorMessage(o.Expected)))
	} else if o.Operation == "push" {
		parts = append(parts, fmt.Sprintf("value: %d", o.Value))
	} else if o.Operation == "count" || o.Operation == "peek" || o.Operation == "pop" {
		parts = append(parts, fmt.Sprintf("expectedInt: %d", int(o.Expected.(float64))))
	} else if o.Operation == "toList" {
		e := o.Expected.([]any)
		nums := make([]int, len(e))
		for i, n := range e {
			nums[i] = int(n.(float64))
		}
		parts = append(parts, fmt.Sprintf("expectedInts: %#v", nums))
	}
	return fmt.Sprintf("Operation{%s}", strings.Join(parts, ", "))
}

type testCase struct {
	Description string   `json:"description"`
	Comments    []string `json:"comments"`
	Input       struct {
		InitialValues []int       `json:"initialValues"`
		Operations    []Operation `json:"operations"`
	} `json:"input"`
}

func (tc testCase) Include() bool {
	return !slices.Contains(tc.Comments, "toList LIFO")
}

var tmpl = `{{.Header}}

type Operation struct {
	operation    string
	value        int
	expectedInt  int
	expectedInts []int
	expectedErr  string
}

type testCase struct {
	description   string
	initialValues []int
	operations    []Operation
}

var testCases = []testCase { {{range .J.list}}{{if .Include}}
	{
		description:   {{printf "%q" .Description}},
		initialValues: {{printf "%#v" .Input.InitialValues}},
		operations:    []Operation{ {{range .Input.Operations}}
			{{.}},{{end}}
		},
	},{{end}}{{end}}
}`
