package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"regexp"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"react": &[]testCase{},
	}
	if err := gen.Gen("react", j, t); err != nil {
		log.Fatal(err)
	}
}

type Cell struct {
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	InitialValue    int      `json:"initial_value"`
	ComputeFunction string   `json:"compute_function"`
	Inputs          []string `json:"inputs"`
}

func (c Cell) String() string {
	if c.Type == "input" {
		return fmt.Sprintf("CreateCell{cType: %q, name: %q, value: %d}", c.Type, c.Name, c.InitialValue)
	} else if c.Type == "compute" && (len(c.Inputs) == 1 || len(c.Inputs) == 2) {
		fnArgs := "x"
		fnName := "fn1"
		if len(c.Inputs) == 2 {
			fnArgs = "x, y"
			fnName = "fn2"
		}
		fn := c.ComputeFunction
		fn = strings.ReplaceAll(fn, "inputs[0]", "x")
		fn = strings.ReplaceAll(fn, "inputs[1]", "y")
		if m := regexp.MustCompile(`if (.*) then (.*) else (.*)`).FindStringSubmatch(fn); m != nil {
			fn = fmt.Sprintf("if %s { return %s }; return %s", m[1], m[2], m[3])
		} else {
			fn = fmt.Sprintf("return %s", fn)
		}
		fn = fmt.Sprintf("func(%s int) int {%s}", fnArgs, fn)
		return fmt.Sprintf(
			"CreateCell{cType: %q, name: %q, inputs: %#v, %s: %s, fnString: %q}",
			fmt.Sprintf("%s%d", c.Type, len(c.Inputs)),
			c.Name,
			c.Inputs,
			fnName,
			fn,
			fn,
		)
	}
	panic(fmt.Sprintf("Invalid cell %#v", c))
}

type Operation struct {
	Type              string         `json:"type"`
	Cell              string         `json:"cell"`
	Name              string         `json:"name"`
	Value             int            `json:"value"`
	ExpectCallbacks   map[string]int `json:"expect_callbacks"`
	ExpectNoCallbacks []string       `json:"expect_callbacks_not_to_be_called"`
}

func (o Operation) String() string {
	switch o.Type {
	case "set_value":
		parts := []string{
			fmt.Sprintf("oType: %q", "SetValue"),
			fmt.Sprintf("cell: %q", o.Cell),
			fmt.Sprintf("value: %d", o.Value),
		}
		if len(o.ExpectCallbacks) != 0 {
			parts = append(parts, fmt.Sprintf("wantCallbacks: %#v", o.ExpectCallbacks))
		}
		if len(o.ExpectNoCallbacks) != 0 {
			parts = append(parts, fmt.Sprintf("wantNoCallbacks: %#v", o.ExpectNoCallbacks))
		}
		return fmt.Sprintf("Operation{%s}", strings.Join(parts, ", "))
	case "add_callback":
		return fmt.Sprintf("Operation{oType: %q, cell: %q, name: %q}", "AddCallback", o.Cell, o.Name)
	case "remove_callback":
		return fmt.Sprintf("Operation{oType: %q, cell: %q, name: %q}", "Cancel", o.Cell, o.Name)
	case "expect_cell_value":
		return fmt.Sprintf("Operation{oType: %q, cell: %q, value: %d}", "Value", o.Cell, o.Value)
	default:
		panic("Unknown operation type, " + o.Type)
	}
}

type testCase struct {
	Description string   `json:"description"`
	Comments    []string `json:"comments"`
	Input       struct {
		Cells []Cell `json:"cells"`
		Operations []Operation `json:"operations"`
	} `json:"input"`
}

var tmpl = `{{.Header}}

type CreateCell struct {
	cType    string
	name     string
	value    int
	inputs   []string
	fn1      func (int) int
	fn2      func (int, int) int
	fnString string
}

type Operation struct {
	oType           string
	cell            string
	name            string
	value           int
	wantCallbacks   map[string]int
	wantNoCallbacks []string
}

type testCase struct {
	description   string
	cells         []CreateCell
	operations    []Operation
}

var testCases = []testCase { {{range .J.react}}
	{
		description:   {{printf "%q" .Description}},
		cells: []CreateCell{ {{range .Input.Cells}}
			{{.}},{{end}}
		},
		operations:    []Operation{ {{range .Input.Operations}}
			{{.}},{{end}}
		},
	},{{end}}
}`
