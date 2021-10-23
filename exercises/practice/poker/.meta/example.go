package poker

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

const (
	Jack  = 11
	Queen = 12
	King  = 13
	Ace   = 14
)

type kind int

const (
	highCard kind = iota
	onePair
	twoPair
	threeOfAKind
	straight
	flush
	fullHouse
	fourOfAKind
	straightFlush
)

type ordering int

const (
	lessThan    ordering = -1
	equalTo     ordering = 0
	greaterThan ordering = 1
)

type card struct {
	rank int
	suit rune
}

type rankCount struct {
	rank  int
	count int
}

type rankCountSlice []rankCount

func (s rankCountSlice) Len() int      { return len(s) }
func (s rankCountSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s rankCountSlice) Less(i, j int) bool {
	if s[i].count != s[j].count {
		return s[i].count < s[j].count
	}
	return s[i].rank < s[j].rank
}

// The strength of a hand in a way that's easily comparable.
// Each hand is considered to have a kind and a list of discriminators.
// The discriminators are ranks, hands of equal kind are first compared by the
// first discriminator of both hands and the one with the highest discriminator
// is better. If there's a tie the second discriminator is considered, etc.
type handValue struct {
	kind   kind
	discrs []int
}

// Returns lessThan, equalTo or greaterThan
func (hv handValue) Compare(other handValue) ordering {
	if hv.kind == other.kind {
		if len(hv.discrs) != len(other.discrs) {
			// Shouldn't happen, number of discriminators for a kind is fixed
			panic("Hands of same kind with different number of discriminators")
		}
		for i := 0; i < len(hv.discrs); i++ {
			if hv.discrs[i] != other.discrs[i] {
				if hv.discrs[i] < other.discrs[i] {
					return lessThan
				}
				return greaterThan
			}
		}
		return equalTo
	} else if hv.kind < other.kind {
		return lessThan
	}
	return greaterThan
}

func BestHand(hands []string) ([]string, error) {
	values := make([]handValue, len(hands))
	for i, hand := range hands {
		cards, err := parseHand(hand)
		if err != nil {
			return nil, err
		}
		counts := countRanks(cards)
		values[i] = evalHand(counts, cards)
	}

	idxes := selectBest(values)
	best := make([]string, len(idxes))
	for i, idx := range idxes {
		best[i] = hands[idx]
	}
	return best, nil
}

func selectBest(values []handValue) []int {
	var best []int
	for i, hv := range values {
		if len(best) == 0 {
			best = []int{i}
		} else {
			switch values[best[0]].Compare(hv) {
			case lessThan:
				best = []int{i}
			case equalTo:
				best = append(best, i)
			}
		}
	}
	return best
}

func evalHand(counts []rankCount, cards []card) handValue {
	isFlush := isFlush(cards)
	isStraight := isStraight(counts)
	var kind kind
	// This switch is why rank counts get ordered in a particular way.
	switch {
	case isStraight && isFlush:
		kind = straightFlush
	case counts[0].count == 4:
		kind = fourOfAKind
	case counts[0].count == 3 && counts[1].count == 2:
		kind = fullHouse
	case isFlush:
		kind = flush
	case isStraight:
		kind = straight
	case counts[0].count == 3:
		kind = threeOfAKind
	case counts[0].count == 2 && counts[1].count == 2:
		kind = twoPair
	case counts[0].count == 2:
		kind = onePair
	default:
		kind = highCard
	}
	discrs := make([]int, 0, len(cards))
	// There's an annoying edge case here. We require the number of
	// discriminators to be equal, however with a flush this can get
	// messed up as flushes aren't based on ranks but on suits. The
	// best way to get around that is to special case non-straight
	// flushes.
	switch {
	case kind == flush:
		for _, card := range cards {
			discrs = append(discrs, card.rank)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(discrs)))
	case (kind == straight || kind == straightFlush) &&
		counts[0].rank == Ace && counts[1].rank == 5:
		// For a straight with an Ace through 5,
		// adjust down the discrs value of the Ace to be 1.
		for i := 1; i < len(cards); i++ {
			discrs = append(discrs, counts[i].rank)
		}
		// Ace becomes a 1 in a low straight with Ace.
		discrs = append(discrs, 1)
	default:
		for _, rc := range counts {
			discrs = append(discrs, rc.rank)
		}
	}
	return handValue{kind, discrs}
}

func isFlush(cards []card) bool {
	first := cards[0].suit
	for i := 1; i < len(cards); i++ {
		if cards[i].suit != first {
			return false
		}
	}
	return true
}

func isStraight(counts []rankCount) bool {
	if len(counts) != 5 {
		return false
	}
	first := counts[0].rank
	if first == Ace {
		for i := 1; i < len(counts); i++ { //nolint // false positive
			if counts[i].rank != 6-i {
				goto normal
			}
			return true
		}
	}
normal:
	for i := 1; i < len(counts); i++ {
		if counts[i].rank != first-i {
			return false
		}
	}
	return true
}

// Count the number of cards with equal rank and return them sorted by count
// descending then rank descending.
func countRanks(cards []card) []rankCount {
	var counts []rankCount
loop:
	for _, card := range cards {
		for i, rc := range counts {
			if rc.rank == card.rank {
				counts[i] = rankCount{rc.rank, rc.count + 1}
				continue loop
			}
		}
		counts = append(counts, rankCount{card.rank, 1})
	}
	sort.Sort(sort.Reverse(rankCountSlice(counts)))
	return counts
}

func parseHand(s string) ([]card, error) {
	parts := strings.Split(strings.TrimSpace(s), " ")
	if len(parts) != 5 {
		return nil, fmt.Errorf("invalid number of hand parts: %d", len(parts))
	}
	cards := make([]card, 5)
	for i, part := range parts {
		card, err := parseCard(part)
		if err != nil {
			return nil, err
		}
		cards[i] = card
	}
	return cards, nil
}

func parseCard(s string) (card, error) {
	if len(s) < 1 {
		return card{}, fmt.Errorf("too short card string: %q", s)
	}
	// The first character should be a number between 1 and 9, so it's safe to use
	// byte matching.
	var rank int
	suitStart := 1
	if s[0] >= '2' && s[0] <= '9' {
		rank = int(s[0]) - '0'
	} else {
		switch s[0] {
		case '1':
			if len(s) < 2 || s[1] != '0' {
				return card{}, fmt.Errorf("invalid rank in card string: %q", s)
			}
			rank = 10
			suitStart = 2
		case 'J':
			rank = Jack
		case 'Q':
			rank = Queen
		case 'K':
			rank = King
		case 'A':
			rank = Ace
		default:
			return card{}, fmt.Errorf("invalid rank in card string: %q", s)
		}
	}
	suit, size := utf8.DecodeRuneInString(s[suitStart:])
	if suit == utf8.RuneError {
		return card{}, fmt.Errorf("Invalid UTF-8 character in card string: %q", s)
	} else if !(suit == '♡' || suit == '♢' || suit == '♤' || suit == '♧') {
		return card{}, fmt.Errorf("Invalid suit: %c", suit)
	}
	if suitStart+size != len(s) {
		return card{}, fmt.Errorf("Extraneous data after suit in card string: %q", s)
	}
	return card{rank, suit}, nil
}
