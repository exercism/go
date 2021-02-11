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

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	// Nested if statement (no else)
	if isBlackjack {
		if dealerScore < 10 {
			return "W"
		}
		return "S"
	}
	return "P"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore int, dealerScore int) string {
	// Switch statement
	switch {
	case (handScore >= 17) || (handScore >= 12 && dealerScore < 7):
		return "S"
	case (handScore <= 11) || (handScore >= 12 && dealerScore >= 7):
		return "H"
	default:
		return ""
	}
}

// FirstTurn implements the logic for the first turn with the help of the other implemented functions.
// This function is already implemented and does not need to be edited.
func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	if 20 < handScore {
		return LargeHand(IsBlackjack(card1, card2), dealerScore)
	}
	return SmallHand(handScore, dealerScore)
}
