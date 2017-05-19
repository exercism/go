package acronym

import (
	"fmt"
	"testing"
)

const targetTestVersion = 3

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func ExampleAbbreviate() {
	var phrase string
	// basic
	phrase = "Portable Network Graphics"
	fmt.Println(Abbreviate(phrase))
	// lowercase words
	phrase = "Ruby on Rails"
	fmt.Println(Abbreviate(phrase))
	// camelcase
	phrase = "HyperText Markup Language"
	fmt.Println(Abbreviate(phrase))
	// punctuation
	phrase = "First In, First Out"
	fmt.Println(Abbreviate(phrase))
	// all caps words
	phrase = "PHP: Hypertext Preprocessor"
	fmt.Println(Abbreviate(phrase))
	// non-acronym all caps word
	phrase = "GNU Image Manipulation Program"
	fmt.Println(Abbreviate(phrase))
	// hyphenated
	phrase = "Complementary metal-oxide semiconductor"
	fmt.Println(Abbreviate(phrase))

	// Output:
	// PNG
	// ROR
	// HTML
	// FIFO
	// PHP
	// GIMP
	// CMOS
}
