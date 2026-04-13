package main

import (
	"../../../../gen"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"parse": &[]testCase{},
	}
	if err := gen.Gen("sgf-parsing", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Encoded string `json:"encoded"`
	} `json:"input"`
	Expected json.RawMessage `json:"expected"`
}

// expectedNode mirrors the tree structure in canonical-data.json expected values.
type expectedNode struct {
	Properties map[string][]string `json:"properties"`
	Children   []expectedNode      `json:"children"`
}

// ExpectedError returns the error string if this test case expects an error, otherwise "".
func (t testCase) ExpectedError() string {
       var e struct {
               Error string `json:"error"`
       }
       if err := json.Unmarshal(t.Expected, &e); err != nil {
               return ""
       }
       return e.Error
}

// ExpectedNode returns Go source code for the expected *Node value (or "nil" for error cases).
func (t testCase) ExpectedNode() string {
	if t.ExpectedError() != "" {
		return "nil"
	}
	var n expectedNode
	if err := json.Unmarshal(t.Expected, &n); err != nil {
		return "nil"
	}
	return renderNode(n)
}

func renderNode(n expectedNode) string {
	return fmt.Sprintf("&Node{\nProperties: %s,\nChildren: %s,\n}",
		renderProperties(n.Properties),
		renderChildren(n.Children),
	)
}

func renderProperties(props map[string][]string) string {
	if len(props) == 0 {
		return "map[string][]string{}"
	}
	keys := make([]string, 0, len(props))
	for k := range props {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		quotedVals := make([]string, len(props[k]))
		for i, v := range props[k] {
			quotedVals[i] = fmt.Sprintf("%q", v)
		}
		parts = append(parts, fmt.Sprintf("%q: {%s}", k, strings.Join(quotedVals, ", ")))
	}
	return fmt.Sprintf("map[string][]string{%s}", strings.Join(parts, ", "))
}

func renderChildren(children []expectedNode) string {
	if len(children) == 0 {
		return "[]*Node{}"
	}
	parts := make([]string, len(children))
	for i, c := range children {
		parts[i] = renderNode(c) + ","
	}
	return fmt.Sprintf("[]*Node{\n%s\n}", strings.Join(parts, "\n"))
}

var tmpl = `{{.Header}}

type testCase struct {
	description   string
	encoded       string
	expectedError string
	expected      *Node
}

var testCases = []testCase { {{range .J.parse}}
	{
		description:   {{printf "%q" .Description}},
		encoded:       {{printf "%q" .Input.Encoded}},
		expectedError: {{printf "%q" .ExpectedError}},
		expected:      {{.ExpectedNode}},
	},{{end}}
}
`
