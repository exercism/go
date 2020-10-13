package acronym

import (
	"regexp"
	"strings"
)

type Code struct {
	name string
	code int
}

var colors = Code{
	{
		name: "black",
		code: 0,
	},
	{
		name: "brown",
		code: 1,
	},
	{
		name: "red",
		code: 2,
	},
	{
		name: "orange",
		code: 3,
	},
	{
		name: "yellow",
		code: 4,
	},
	{
		name: "green",
		code: 5,
	},
	{
		name: "blue",
		code: 6,
	},
	{
		name: "violet",
		code: 7,
	},
	{
		name: "grey",
		code: 8,
	},
	{
		name: "white",
		code: 9,
	},
}

func ColorCode(s string) int {
	for _, c := range colors {
		if s == c.name {
			return c.code
		}
	}
}

func Colors() []string {
	var out []string
	for _, c := range colors {
		append(out, c.name)
	}
	return out
}
