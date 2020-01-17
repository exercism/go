package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden map[string][]string

func (g *Garden) Plants(child string) ([]string, bool) {
	p, ok := (*g)[child]
	return p, ok
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	rows := strings.Split(diagram, "\n")
	if len(rows) != 3 || rows[0] != "" {
		return nil, errors.New("diagram must have two rows")
	}
	if len(rows[1]) != len(rows[2]) {
		return nil, errors.New("diagram rows must be same length")
	}
	if len(rows[1]) != 2*len(children) {
		return nil, errors.New("each diagram row must have two cups per child")
	}
	g := Garden{}
	alpha := append([]string{}, children...)
	sort.Strings(alpha)
	for _, n := range alpha {
		g[n] = make([]string, 0, 4)
	}
	if len(g) != len(alpha) {
		return nil, errors.New("no two children can have the same name")
	}
	for _, row := range rows[1:] {
		for nx, n := range alpha {
			for cx := range rows[1:] { // a little hack
				var p string
				switch row[2*nx+cx] {
				case 'G':
					p = "grass"
				case 'C':
					p = "clover"
				case 'R':
					p = "radishes"
				case 'V':
					p = "violets"
				default:
					return nil, errors.New("plant codes must be one of G, C, R, or V")
				}
				g[n] = append(g[n], p)
			}
		}
	}
	return &g, nil
}
