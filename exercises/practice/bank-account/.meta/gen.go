package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"bankAccount": &[]testCase{},
	}
	if err := gen.Gen("bank-account", j, t); err != nil {
		log.Fatal(err)
	}
}

type Operation struct {
	Operation string `json:"operation"`
	Amount    any `json:"amount"`
}

func (o Operation) String() string {
	if a, ok := o.Amount.(float64); ok {
		return fmt.Sprintf("Operation{Name: %q, Amount: %d}", o.Operation, int(a))
	}
	return fmt.Sprintf("Operation{Name: %q}", o.Operation)
}

type testCase struct {
	Description string `json:"description"`
	Scenarios   []string `json:"scenarios"`
	Input       struct {
		Operations []Operation `json:"operations"`
	} `json:"input"`
	Expected    any `json:"expected"`
}

func (tc testCase) IsError() bool {
	return gen.IsError(tc.Expected)
}

func (tc testCase) ErrorMessage() string {
	return gen.ErrorMessage(tc.Expected)
}

var tmpl = `{{.Header}}

type Operation struct {
	Name   string
	Amount int64
}

type testCase struct {
	description string
	operations  []Operation
	expected    int64
	expectedErr string
}

var testCases = []testCase { {{range .J.bankAccount}}{{if eq .Scenarios nil}}
	{
		description: {{printf "%q" .Description}},
		operations:  []Operation { {{range .Input.Operations}}
			{{.}},{{end}}
		},{{if .IsError}}
		expectedErr: {{printf "%q" .ErrorMessage}},{{else}}
		expected:    {{.Expected}},{{end}}
	},{{end}}{{end}}
}`
