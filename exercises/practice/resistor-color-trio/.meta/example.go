package resistorcolortrio

import (
	"fmt"
	"strconv"
)

var colorMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func Exp(base, exp int) int {
	if exp == 0 {
		return 1
	}

	res := 1
	for i := 0; i < exp; i++ {
		res *= base
	}

	return res
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	value := 10*colorMap[colors[0]] + colorMap[colors[1]]
	value *= Exp(10, colorMap[colors[2]])
	label := strconv.Itoa(value)

	var unit string
	switch {
	case len(label) < 4:
		unit = "ohms"
	case len(label) < 7:
		label = strconv.Itoa(value / 1000)
		unit = "kiloohms"
	case len(label) <= 8:
		label = strconv.Itoa(value / 1000_000)
		unit = "megaohms"
	case len(label) >= 9:
		label = strconv.Itoa(value / 1000_000_000)
		unit = "gigaohms"
	}

	if value < 1000 {
		label = strconv.Itoa(value)
	}

	return fmt.Sprintf("%s %s", label, unit)
}
