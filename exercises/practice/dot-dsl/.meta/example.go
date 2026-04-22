package dotdsl

import (
	"fmt"
	"strconv"
	"strings"
)

// Properties holds the properties of a node or edge.
// The values can be int, bool or string.
type Properties map[string]any

// Graph stores the parts of a dot graph.
// All entities are stored as a Properties map (`nil` Properties when none set)
// attrs is the Properties for the entire Graph, vs a specific node or edge.
type Graph struct {
	nodes map[string]Properties
	edges map[string]Properties
	attrs Properties
}

// extractPropSection splits a line like "a -- b [prop=val]" into "a -- b" and "prop=val".
func extractPropSection(line string) (string, string, error) {
	// Parse the property, if any.
	if !strings.ContainsAny(line, "[]") {
		return line, "", nil
	}
	if strings.Count(line, "[") != 1 || strings.Count(line, "]") != 1 {
		return "", "", fmt.Errorf("incorrect number of []")
	}
	if !strings.HasSuffix(line, "]") {
		return "", "", fmt.Errorf("incorrectly placed ']'")
	}
	line, props, ok := strings.Cut(strings.TrimSuffix(line, "]"), "[")
	if !ok || len(props) < 3 || !strings.Contains(props[1:len(props)-1], "=") {
		return "", "", fmt.Errorf("badly formatted props")
	}
	return line, props, nil
}

// parseProp turns a prop like `foo=2` into `foo` and `int(5)`.
// It handles bool, int, quoted strings and (as a fallback) unquoted strings.
func parseProp(rawProp string) (string, any, error) {
	propName, rawPropVal, hasProp := strings.Cut(rawProp, "=")
	if !hasProp || len(propName) == 0 || len(rawPropVal) == 0 {
		return "", nil, fmt.Errorf("badly formatted prop")
	}

	var propVal any
	if rawPropVal == "true" {
		propVal = true
	} else if rawPropVal == "false" {
		propVal = false
	} else if v, err := strconv.Atoi(rawPropVal); err == nil {
		propVal = v
	} else if strings.HasPrefix(rawPropVal, `"`) && strings.HasSuffix(rawPropVal, `"`) {
		propVal = strings.Trim(rawPropVal, `"`)
	} else {
		propVal = rawPropVal
	}
	return propName, propVal, nil
}

// extractNodes parses a line like "a -- b -- c" to return []string{"a", "b", "c"}.
func extractNodes(line string) ([]string, error) {
	nodes := strings.Split(line, "--")
	if len(nodes) == 0 || len(nodes) == 1 && nodes[0] == "" {
		return nil, nil
	}
	for i, node := range nodes {
		node = strings.TrimSpace(node)
		nodes[i] = node
		if len(node) == 0 {
			return nil, fmt.Errorf("badly formatted node, line %q", line)
		}
		for _, r := range node {
			if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '1')) {
				return nil, fmt.Errorf("node name should be alnum, got %s", node)
			}
		}
	}
	return nodes, nil
}

// nodesToEdges takes a slice of nodes eg []string{"a", "b", "c"} and returns edges,
// like []string{"{a b", "{b c}"}
func nodesToEdges(nodes []string) []string {
	if len(nodes) < 2 {
		return nil
	}
	edges := make([]string, len(nodes)-1)
	for i := range len(nodes) - 1 {
		a, b := nodes[i], nodes[i+1]
		if a > b {
			a, b = b, a
		}
		edges[i] = fmt.Sprintf("{%s %s}", a, b)
	}
	return edges
}

// addKeys updates a map, setting missing entries to nil entries.
func addKeys(keys []string, myMap map[string]Properties) {
	for _, key := range keys {
		if _, ok := myMap[key]; !ok {
			myMap[key] = nil
		}
	}
}

// setProp updates the Properties in a map for a set of keys.
func setProp(keys []string, myMap map[string]Properties, prop string, value any) {
	for _, key := range keys {
		if myMap[key] == nil {
			myMap[key] = Properties{}
		}
		myMap[key][prop] = value
	}
}

// parseLine updates the Graph with a single line.
func (g *Graph) parseLine(line string) error {
	// Check the line format.
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
		return nil
	}
	if !strings.HasSuffix(line, ";") {
		return fmt.Errorf("line does not end in a semicolon")
	}
	line = strings.TrimSuffix(line, ";")

	// Parse the property, if any.
	var propName string
	var propVal any
	var hasProp bool

	if lineSection, propSection, err := extractPropSection(line); err != nil {
		return err
	} else if propSection != "" {
		hasProp = true
		line = lineSection
		if propName, propVal, err = parseProp(propSection); err != nil {
			return err
		}
	}

	nodes, err := extractNodes(line)
	if err != nil {
		return err
	}
	edges := nodesToEdges(nodes)
	addKeys(nodes, g.nodes)
	addKeys(edges, g.edges)

	if hasProp {
		switch len(nodes) {
		case 0:
			g.attrs[propName] = propVal
		case 1:
			setProp(nodes, g.nodes, propName, propVal)
		default:
			setProp(edges, g.edges, propName, propVal)
		}
	}
	return nil
}

// Parse creates a Graph from a text blob.
func Parse(data string) (*Graph, error) {
	g := &Graph{make(map[string]Properties), make(map[string]Properties), Properties{}}
	lines := strings.Split(data, "\n")
	// Check the graph start/end.
	if len(lines) < 2 || lines[0] != "graph {" || lines[len(lines)-1] != "}" {
		return nil, fmt.Errorf("not a graph")
	}
	for i, line := range lines[1 : len(lines)-1] {
		if err := g.parseLine(line); err != nil {
			return nil, fmt.Errorf("line %d: %q -- parse error %w", i, line, err)
		}
	}
	return g, nil
}
