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
		"buildTree": &[]testCase{},
	}
	if err := gen.Gen("tree-building", j, t); err != nil {
		log.Fatal(err)
	}
}

type Node struct {
	ID       int    `json:"recordId"`
	Children []Node `json:"children"`
}

func (n *Node) String() string {
	var out strings.Builder
	if len(n.Children) == 0 {
		return fmt.Sprintf("&Node{ID: %d}", n.ID)
	}
	out.WriteString(fmt.Sprintf("&Node{\nID: %d,\n", n.ID))
	out.WriteString("Children: []*Node{\n")
	for _, c := range n.Children {
		out.WriteString(c.String())
		out.WriteString(",\n")
	}
	out.WriteString("},\n")
	out.WriteString("}")
	return out.String()
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Records []struct {
			RecordID int `json:"recordId"`
			ParentID int `json:"parentId"`
		} `json:"records"`
	} `json:"input"`
	Expected struct {
		Error string `json:"error"`
		Node  *Node  `json:"node"`
	} `json:"expected"`
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	input       []Record
	wantErr     bool
	expected    *Node
}

var testCases = []testCase { {{range .J.buildTree}}
	{
		description: {{printf "%q" .Description}},
		input:       []Record{ {{range .Input.Records}}
			Record{ID: {{.RecordID}}, Parent: {{.ParentID}}},{{end}}
		},{{if ne .Expected.Error ""}}
		wantErr:     true,{{else if eq .Expected.Node nil}}
		expected:    nil,{{else}}
		expected:    {{.Expected.Node}},{{end}}
	},{{end}}
}`
