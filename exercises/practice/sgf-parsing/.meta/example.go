package sgfparsing

import (
	"errors"
	"strings"
)

// Node represents an SGF node with properties and child nodes.
type Node struct {
	Properties map[string][]string
	Children   []*Node
}

func Parse(encoded string) (*Node, error) {
	s := strings.TrimSpace(encoded)
	if len(s) == 0 {
		return nil, errors.New("tree missing")
	}
	node, _, err := parseGameTree(s)
	return node, err
}

func parseGameTree(s string) (*Node, string, error) {
	if len(s) == 0 || s[0] != '(' {
		return nil, s, errors.New("tree missing")
	}
	s = s[1:] // consume '('

	if len(s) == 0 || s[0] != ';' {
		return nil, s, errors.New("tree with no nodes")
	}

	var nodes []*Node
	for len(s) > 0 && s[0] == ';' {
		s = s[1:] // consume ';'
		node, rest, err := parseNode(s)
		if err != nil {
			return nil, rest, err
		}
		nodes = append(nodes, node)
		s = rest
	}

	// subtrees become children of the last node in the sequence
	last := nodes[len(nodes)-1]
	for len(s) > 0 && s[0] == '(' {
		child, rest, err := parseGameTree(s)
		if err != nil {
			return nil, rest, err
		}
		last.Children = append(last.Children, child)
		s = rest
	}

	if len(s) == 0 || s[0] != ')' {
		return nil, s, errors.New("unexpected end of game tree")
	}
	s = s[1:] // consume ')'

	// chain sequence: nodes[i] → nodes[i+1]
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Children = []*Node{nodes[i+1]}
	}

	return nodes[0], s, nil
}

func parseNode(s string) (*Node, string, error) {
	node := &Node{
		Properties: make(map[string][]string),
		Children:   []*Node{},
	}

	for len(s) > 0 && isASCIILetter(s[0]) {
		// read property key
		i := 0
		for i < len(s) && isASCIILetter(s[i]) {
			if isLower(s[i]) {
				return nil, s, errors.New("property must be in uppercase")
			}
			i++
		}
		key := s[:i]
		s = s[i:]

		if len(s) == 0 || s[0] != '[' {
			return nil, s, errors.New("properties without delimiter")
		}

		var values []string
		for len(s) > 0 && s[0] == '[' {
			s = s[1:] // consume '['
			val, rest, err := parsePropertyValue(s)
			if err != nil {
				return nil, rest, err
			}
			values = append(values, val)
			s = rest
		}
		node.Properties[key] = values
	}

	return node, s, nil
}

func parsePropertyValue(s string) (string, string, error) {
	var buf strings.Builder
	i := 0
	for i < len(s) {
		c := s[i]
		if c == ']' {
			return buf.String(), s[i+1:], nil
		}
		if c == '\\' {
			i++
			if i >= len(s) {
				break
			}
			next := s[i]
			if next == '\n' {
				// escaped newline: soft line break, discard
				i++
				continue
			}
			if isNonNewlineWhitespace(next) {
				// escaped whitespace other than newline → space
				buf.WriteByte(' ')
				i++
				continue
			}
			// any other character → insert as-is
			buf.WriteByte(next)
			i++
			continue
		}
		if isNonNewlineWhitespace(c) {
			buf.WriteByte(' ')
			i++
			continue
		}
		buf.WriteByte(c)
		i++
	}
	return "", s[i:], errors.New("unexpected end of property value")
}

func isASCIILetter(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isNonNewlineWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\f' || c == '\v'
}
