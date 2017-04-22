package romannumerals

import (
	"bytes"
	"fmt"
)

const testVersion = 3

type arabicToRoman struct {
	arabic int
	roman  string
}

func ToRomanNumeral(input int) (string, error) {
	buffer := bytes.NewBufferString("")

	if input <= 0 || input >= 4000 {
		return "", fmt.Errorf("the number %d is undefined in the roman numeral system", input)
	}

	mappings := []arabicToRoman{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, m := range mappings {
		for input >= m.arabic {
			buffer.WriteString(m.roman)
			input -= m.arabic
		}
	}

	return buffer.String(), nil
}
