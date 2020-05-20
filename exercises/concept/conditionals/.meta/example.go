package _meta

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// Blackjack returns true if the player has a blackjack, false otherwise.
func Blackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// FirstTurn returns the "optimal" decision for the first turn, given the cards of the player and the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseCard(card1) + ParseCard(card2)
	// Single if statement
	if handScore == 22 {
		return "P"
	}
	dealerScore := ParseCard(dealerCard)
	// Nested if statement (no else)
	if Blackjack(card1, card2) {
		if dealerScore < 10 {
			return "W"
		}
		return "S"
	}
	// Switch statement
	switch {
	case (handScore >= 17) || (handScore >= 12 && dealerScore <= 6):
		return "S"
	case (handScore <= 11) || (handScore >= 12 && dealerScore > 6):
		return "H"
	default:
		return ""
	}
}
