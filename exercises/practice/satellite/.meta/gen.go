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
		"treeFromTraversals": &[]testCase{},
	}
	if err := gen.Gen("satellite", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Preorder []string `json:"preorder"`
		Inorder  []string `json:"inorder"`
	} `json:"input"`
	Expected map[string]any `json:"expected"`
}

func NodeString(node map[string]any) string {
	if len(node) == 0 {
		return "nil"
	}
	l, _ := node["l"].(map[string]any)
	r, _ := node["r"].(map[string]any)
	return fmt.Sprintf(
		"&Node{Value: %q, Left: %s, Right: %s}",
		node["v"],
		NodeString(l),
		NodeString(r),
	)
}

func (tc testCase) Tree() string {
	return NodeString(tc.Expected)
}

var tmpl = `{{.Header}}

var testCases = []struct {
	description string
	preorder    []string
	inorder     []string
	expected    *Node
	err         string
}{ {{range .J.treeFromTraversals}}
	{
		description: {{printf "%q" .Description}},
		preorder:    {{printf "%#v" .Input.Preorder}},
		inorder:     {{printf "%#v" .Input.Inorder}}, {{if not .Expected.error}}
		expected:    {{.Tree}},{{end}} {{if .Expected.error}}
		err:         {{printf "%q" .Expected.error}},{{end}}
	},{{end}}
}`
