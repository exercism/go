package resistorcolortolerance

import (
	"fmt"
	"math"
	"slices"
)

var Colors = []string{
	"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white",
}

var Tolerance = map[string]string{
	"grey":   "±0.05%",
	"violet": "±0.1%",
	"blue":   "±0.25%",
	"green":  "±0.5%",
	"brown":  "±1%",
	"red":    "±2%",
	"gold":   "±5%",
	"silver": "±10%",
}

var Units = []string{
	"ohms", "kiloohms", "megaohms", "gigaohms",
}

func Label(colors []string) string {
	if len(colors) == 1 {
		return "0 ohms"
	}

	var ohms float64
	for i := 0; i < len(colors)-2; i++ {
		ohms *= 10
		ohms += float64(slices.Index(Colors, colors[i]))
	}
	ohms *= math.Pow(10.0, float64(slices.Index(Colors, colors[len(colors)-2])))

	var shifts int
	for ohms >= 1000 {
		shifts++
		ohms /= 1000
	}
	return fmt.Sprintf("%g %s %s", ohms, Units[shifts], Tolerance[colors[len(colors)-1]])
}
