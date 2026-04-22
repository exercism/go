package main

import (
	"../../../../gen"
	"fmt"
	"iter"
	"log"
	"maps"
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
		"parse": &[]testCase{},
	}
	if err := gen.Gen("dot-dsl", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Text []string `json:"text"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (tc testCase) JoinedText() string {
	return strings.Join(tc.Input.Text, "\n")
}

func (tc testCase) IsError() bool {
	return gen.IsError(tc.Expected)
}

func (tc testCase) ErrorMessage() string {
	return gen.ErrorMessage(tc.Expected)
}

func sortedKeys(keys iter.Seq[string]) []string {
	var out []string
	for key := range keys {
		out = append(out, key)
	}
	slices.Sort(out)
	return out
}

func parseProperties(rawProperties any) map[string]string {
	properties, ok := rawProperties.(map[string]any)
	if !ok {
		panic(fmt.Sprintf("Failed to parse properties %v", rawProperties))
	}
	if len(properties) == 0 {
		return nil
	}
	out := map[string]string{}
	for key, value := range properties {
		switch value.(type) {
		case bool:
			out[key] = fmt.Sprintf("%t", value)
		case float64:
			out[key] = fmt.Sprintf("%d", int(value.(float64)))
		case string:
			out[key] = fmt.Sprintf("%q", value)
		default:
			panic(fmt.Sprintf("Unknown property type, %s %v", key, value))
		}
	}
	return out
}

func (tc testCase) ExpectedMap(key string) map[string]map[string]string {
	data, ok := tc.Expected.(map[string]any)
	if !ok {
		panic("Failed to parse Expected")
	}
	data, ok = data[key].(map[string]any)
	if !ok {
		panic("Failed to parse Expected")
	}
	if len(data) == 0 {
		return nil
	}
	out := map[string]map[string]string{}
	for node, rawProperties := range data {
		out[node] = parseProperties(rawProperties)
	}
	return out
}

func (tc testCase) FormatNodesEdges(key string) string {
	data := tc.ExpectedMap(key)
	if data == nil {
		return "nil"
	}
	var parts []string
	for _, name := range sortedKeys(maps.Keys(data)) {
		properties := data[name]
		if len(properties) == 0 {
			parts = append(parts, fmt.Sprintf("%q: nil", name))
		} else {
			var props []string
			for _, prop := range sortedKeys(maps.Keys(properties)) {
				props = append(props, fmt.Sprintf("%q: %s", prop, properties[prop]))
			}
			parts = append(parts, fmt.Sprintf("%q: {%s}", name, strings.Join(props, ", ")))
		}
	}
	return fmt.Sprintf("map[string]Properties{%s}", strings.Join(parts, ", "))
}

func (tc testCase) Nodes() string {
	return tc.FormatNodesEdges("nodes")
}

func (tc testCase) Edges() string {
	return tc.FormatNodesEdges("edges")
}

func (tc testCase) Attrs() string {
	data, ok := tc.Expected.(map[string]any)
	if !ok {
		panic("Failed to parse Expected")
	}
	data, ok = data["attrs"].(map[string]any)
	if !ok {
		panic("Failed to parse Expected")
	}
	properties := parseProperties(data)
	if properties == nil {
		return "nil"
	}
	var props []string
	for _, prop := range sortedKeys(maps.Keys(properties)) {
		props = append(props, fmt.Sprintf("%q: %s", prop, properties[prop]))
	}
	return fmt.Sprintf("Properties{%s}", strings.Join(props, ", "))
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	input       string
	nodes       map[string]Properties
	edges       map[string]Properties
	attrs       Properties
	expectedErr string
}

var testCases = []testCase { {{range .J.parse}}
	{
		description: {{printf "%q" .Description}},
		input:       {{printf "%q" .JoinedText}},{{if .IsError}}
		expectedErr: {{printf "%q" .ErrorMessage}},{{else}}
		nodes:       {{.Nodes}},
		edges:       {{.Edges}},
		attrs:       {{.Attrs}},{{end}}
	},{{end}}
}`
