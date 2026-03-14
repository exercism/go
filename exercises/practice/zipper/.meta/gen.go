package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"strings"
	"text/template"
)

// Node is used to build a tree.
type Node struct {
	value int
	right *Node
	left  *Node
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("&Node{value: %d, left: %s, right: %s}", n.value, n.left, n.right)
}

func newNode(n any) *Node {
	if n, ok := n.(map[string]any); !ok {
		return nil
	} else {
		if v, ok := n["value"].(float64); !ok {
			panic("Cannot parse data")
		} else {
			return &Node{
				value: int(v),
				right: newNode(n["right"]),
				left:  newNode(n["left"]),
			}
		}
	}
}

// Operation is a Stringifier that holds an operation.

type Operation map[string]any

func (op Operation) String() string {
	funcCall, ok := op["operation"].(string)
	if !ok {
		panic("Operation lacks a string name")
	}
	parts := []string{fmt.Sprintf("funcCall: %q", funcCall)}

	if item, ok := op["item"]; ok {
		switch v := item.(type) {
		case float64:
			parts = append(parts, fmt.Sprintf("intArg: %d", int(v)))
		case map[string]any:
			parts = append(parts, fmt.Sprintf("nodeArg: %s", newNode(v)))
		case nil:
			parts = append(parts, "nodeArg: nil")
		default:
			panic("Unknown item type")
		}
	}
	return fmt.Sprintf("operation{%s}", strings.Join(parts, ", "))
}

// The testCase for expectedValue tests

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		InitialTree map[string]any `json:"initialTree"`
		Operations  []Operation    `json:"operations"`
	} `json:"input"`
	Expected struct {
		Type  string `json:"type"`
		Value any    `json:"value"`
	} `json:"expected"`
}

func (tc testCase) InitialTree() *Node {
	return newNode(tc.Input.InitialTree)
}

func (tc testCase) ExpectedTree() *Node {
	return newNode(tc.Expected.Value)
}

func (tc testCase) ExpectedIntVal() int {
	if tc.Expected.Type != "int" {
		panic("Not expecting an int")
	}
	v, ok := tc.Expected.Value.(float64)
	if !ok {
		panic("Expected is not a number")
	}
	return int(v)
}

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"expectedValue":            &[]testCase{},
		"sameResultFromOperations": &[]any{}, // Not implemented
	}
	if err := gen.Gen("zipper", j, t); err != nil {
		log.Fatal(err)
	}
}

var tmpl = `{{.Header}}

type operation struct {
	funcCall string
	intArg   int
	nodeArg  *Node
}

type want struct {
	wantType string
	intVal   int
	nodeVal  *Node
}

type testCase struct {
	description   string
	input         *Node
	operations    []operation
	want          want
}

var testCases = []testCase { {{range .J.expectedValue}}
	{
		description:   {{printf "%q" .Description}},
		input:         {{.InitialTree}},
		operations:    []operation{ {{range .Input.Operations}}
			{{.}},{{end}}
		},
		want:          want{
			wantType:  {{printf "%q" .Expected.Type}},{{if eq .Expected.Type "int"}}
			intVal:    {{.ExpectedIntVal}},{{end}}{{if eq .Expected.Type "tree"}}
			nodeVal:   {{.ExpectedTree}},{{end}}
		},
	},{{end}}
}`
