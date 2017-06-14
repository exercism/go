package main

import (
	"log"
	"strings"
	"text/template"

	"../../../gen"
)

func main() {
	t := template.New("").Funcs(template.FuncMap{
		"pokerHands": PokerHands,
	})
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("poker", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       []string
		Expected    []string
	}
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

type validTestCase struct {
	name  string
	hands []string
	best  []string
}

var validTestCases = []validTestCase {
{{range .J.Cases}}{
	name:  "{{.Description}}",
	hands: []string{
		{{range pokerHands .Input}} {{printf "%q" .}},
{{end}} },
	best:  {{pokerHands .Expected | printf "%#v"}},
},
{{end}}}
`
