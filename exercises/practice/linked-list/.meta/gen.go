package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"slices"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"list": &[]testCase{},
	}
	if err := gen.Gen("linked-list", j, t); err != nil {
		log.Fatal(err)
	}
}

type Operation struct {
	Operation string `json:"operation"`
	Value     int    `json:"value"`
	Expected  int    `json:"expected"`
}

func (op Operation) String() string {
	hasInput := []string{"push", "unshift", "delete"}
	opFunc := strings.ToUpper(op.Operation[0:1]) + op.Operation[1:]
	if slices.Contains(hasInput, op.Operation) {
		return fmt.Sprintf("{operation: %q, value: %d}", opFunc, op.Value)
	}
	return fmt.Sprintf("{operation: %q, expected: %d}", opFunc, op.Expected)
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Operations []Operation `json:"operations"`
	} `json:"input"`
}

var tmpl = `// This file contains tests from the shared problem specifications repo.
{{.Header}}

type Operation struct{
	operation string
	value int
	expected int
}

var testCases = []struct{
	description string
	operations []Operation
}{ {{range .J.list}}
	{
		description: 	{{printf "%q" .Description}},
		operations:     []Operation{ {{range .Input.Operations}}
			{{.String}},{{end}}
		},
	},{{end}}
}`
