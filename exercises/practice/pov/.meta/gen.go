package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"fromPov": &[]povTestCase{},
		"pathTo":  &[]pathTestCase{},
	}
	if err := gen.Gen("pov", j, t); err != nil {
		log.Fatal(err)
	}
}

type Tree struct {
	Label    string `json:"label"`
	Children []Tree `json:"children"`
}

func (n *Tree) String() string {
	if n == nil {
		return "nil"
	}
	args := []string{fmt.Sprintf("%q", n.Label)}
	for _, child := range n.Children {
		args = append(args, child.String())
	}
	return fmt.Sprintf("New(%s)", strings.Join(args, ", "))
}

type povTestCase struct {
	Description string `json:"description"`
	Input       struct {
		From string `json:"from"`
		Tree *Tree  `json:"tree"`
	} `json:"input"`
	Expected *Tree `json:"expected"`
}

type pathTestCase struct {
	Description string `json:"description"`
	Input       struct {
		From string `json:"from"`
		To   string `json:"to"`
		Tree *Tree  `json:"tree"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

var tmpl = `{{.Header}}

type povTestCase struct {
	description string
	from        string
	tree        *Tree
	expected    *Tree
}

var povTestCases = []povTestCase { {{range .J.fromPov}}
	{
		description: {{printf "%q" .Description}},
		from:        {{printf "%q" .Input.From}},
		tree:        {{.Input.Tree}},
		expected:    {{.Expected}},
	},{{end}}
}

type pathTestCase struct {
	description string
	from        string
	to          string
	tree        *Tree
	expected    []string
}

var pathTestCases = []pathTestCase { {{range .J.pathTo}}
	{
		description: {{printf "%q" .Description}},
		from:        {{printf "%q" .Input.From}},
		to:          {{printf "%q" .Input.To}},
		tree:        {{.Input.Tree}},
		expected:    {{printf "%#v" .Expected}},
	},{{end}}
}`
