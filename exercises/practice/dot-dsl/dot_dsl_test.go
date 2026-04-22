package dotdsl

import (
	"fmt"
	"maps"
	"strings"
	"testing"
)

func propsEqual(a, b any) bool {
	return fmt.Sprintf("%T %v", a, a) == fmt.Sprintf("%T %v", b, b)
}

func propsMapEqual(a, b Properties) bool {
	return maps.EqualFunc(a, b, propsEqual)
}

func TestParse(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if got, err := Parse(tc.input); err != nil {
				if tc.expectedErr == "" {
					t.Fatalf("dsl = %q\nParse(dsl) returned unexpected error %v", tc.input, err)
				}
			} else if tc.expectedErr != "" {
				t.Fatalf("dsl = %q\nParse(dsl) did not return an error; expected error %v", tc.input, tc.expectedErr)
			} else if got == nil {
				t.Fatalf("dsl = %q\nParse(dsl) = nil, expected a Graph", tc.input)
			} else {
				var diffs []string
				if !maps.EqualFunc(got.nodes, tc.nodes, propsMapEqual) {
					diffs = append(diffs, fmt.Sprintf("got nodes: %#v\nwant nodes: %#v", got.nodes, tc.nodes))
				}
				if !maps.EqualFunc(got.edges, tc.edges, propsMapEqual) {
					diffs = append(diffs, fmt.Sprintf("got edges: %#v\nwant edges: %#v", got.edges, tc.edges))
				}
				if !maps.EqualFunc(got.attrs, tc.attrs, propsEqual) {
					diffs = append(diffs, fmt.Sprintf("got attrs: %#v\nwant attrs: %#v", got.attrs, tc.attrs))
				}
				if len(diffs) != 0 {
					t.Fatalf("dsl = %q\nParse(dsl) gave incorrect graph.\n%s", tc.input, strings.Join(diffs, "\n"))
				}
			}
		})
	}
}

/*
func BenchmarkValid(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			Valid(tc.input)
		}
	}
}
*/
