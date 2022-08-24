package main

import (
	"log"
	"strings"
	"text/template"

	"../../../../gen"
)

func main() {
	t := template.New("").Funcs(template.FuncMap{
		"pokerHands": PokerHands,
	})
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]interface{}{
		"bestHands": &[]testCase{},
	}
	if err := gen.Gen("poker", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Hands []string `json:"hands"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

var suitCharToSymbol = map[byte]string{'C': "♧", 'D': "♢", 'H': "♡", 'S': "♤"}

// PokerHands converts hands using C/D/H/S as the suit
// to one which uses the symbols ♧ / ♢ / ♡ / ♤ instead.
// For example, "2H 3C 4D 5S 9S" is converted to "2♡ 3♧ 4♢ 5♤ 9♤".
func PokerHands(inputHands []string) []string {
	hands := make([]string, len(inputHands))
	for i, inputHand := range inputHands {
		cards := strings.Fields(inputHand)
		handStr := ""
		for j, c := range cards {
			rankStr := c[0 : len(c)-1]
			handStr += rankStr + suitCharToSymbol[c[len(c)-1]]
			if j != len(cards)-1 {
				handStr += " "
			}
		}
		hands[i] = handStr
	}
	return hands
}

// template applied to above data structure generates the Go test cases
var tmpl = `package poker

{{.Header}}

type testCase struct {
	description string
	hands    []string
	expected []string
}

var validTestCases = []testCase {
{{range .J.bestHands}}{
	description:  	{{printf "%q"  .Description }},
	hands: 			{{pokerHands .Input.Hands | printf "%#v"}},
	expected:  		{{pokerHands .Expected | printf "%#v"}},
},
{{end}}}
`
