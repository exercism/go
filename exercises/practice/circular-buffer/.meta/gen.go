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
		"run": &[]testCase{},
	}
	if err := gen.Gen("circular-buffer", j, t); err != nil {
		log.Fatal(err)
	}
}

type Op struct {
	Operation     string `json:"operation"`
	Item          int    `json:"item"`
	ShouldSucceed bool   `json:"should_succeed"`
	Expected      int    `json:"expected"`
}

func (o Op) String() string {
	switch o.Operation {
	case "read":
		if o.ShouldSucceed {
			return fmt.Sprintf("Operation{name: %q, value: %q, wantErr: %t}", "ReadByte", byte(o.Expected+'0'), !o.ShouldSucceed)
		}
		return fmt.Sprintf("Operation{name: %q, wantErr: %t}", "ReadByte", !o.ShouldSucceed)
	case "write":
		return fmt.Sprintf("Operation{name: %q, value: %q, wantErr: %t}", "WriteByte", byte(o.Item+'0'), !o.ShouldSucceed)
	case "overwrite":
		return fmt.Sprintf("Operation{name: %q, value: %q}", "Overwrite", byte(o.Item+'0'))
	case "clear":
		return fmt.Sprintf("Operation{name: %q}", "Reset")
	}
	panic("Unknown op")
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Capacity   int  `json:"capacity"`
		Operations []Op `json:"operations"`
	} `json:"input"`
}

var tmpl = `{{.Header}}

type Operation struct {
	name    string
	value   byte
	wantErr bool
}

type testCase struct {
	description string
	capacity    int
	operations  []Operation
}

var testCases = []testCase { {{range .J.run}}
	{
		description: {{printf "%q" .Description}},
		capacity:    {{.Input.Capacity}},
		operations:  []Operation{ {{range .Input.Operations}}
			{{.}},{{end}}
		},
	},{{end}}
}`
