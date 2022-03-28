package blackjack

import "testing"

func TestParseCard(t *testing.T) {
	tests := []struct {
		name string
		card string
		want int
	}{
		{
			name: "parse ace",
			card: "ace",
			want: 11,
		},
		{
			name: "parse two",
			card: "two",
			want: 2,
		},
		{
			name: "parse three",
			card: "three",
			want: 3,
		},
		{
			name: "parse four",
			card: "four",
			want: 4,
		},
		{
			name: "parse five",
			card: "five",
			want: 5,
		},
		{
			name: "parse six",
			card: "six",
			want: 6,
		},
		{
			name: "parse seven",
			card: "seven",
			want: 7,
		},
		{
			name: "parse eight",
			card: "eight",
			want: 8,
		},
		{
			name: "parse nine",
			card: "nine",
			want: 9,
		},
		{
			name: "parse ten",
			card: "ten",
			want: 10,
		},
		{
			name: "parse jack",
			card: "jack",
			want: 10,
		},
		{
			name: "parse queen",
			card: "queen",
			want: 10,
		},
		{
			name: "parse king",
			card: "king",
			want: 10,
		},
		{
			name: "parse other",
			card: "joker",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCard(tt.card); got != tt.want {
				t.Errorf("ParseCard(%s) = %d, want %d", tt.card, got, tt.want)
			}
		})
	}
}

func TestFirstTurn(t *testing.T) {
	type hand struct {
		card1, card2 string
	}
	tests := []struct {
		name   string
		hand   hand
		dealer string
		want   string
	}{
		{
			name:   "pair of aces",
			hand:   hand{card1: "ace", card2: "ace"},
			dealer: "ace",
			want:   "P",
		},
		{
			name:   "pair of jacks",
			hand:   hand{card1: "jack", card2: "jack"},
			dealer: "ace",
			want:   "S",
		},
		{
			name:   "pair of kings",
			hand:   hand{card1: "king", card2: "king"},
			dealer: "ace",
			want:   "S",
		},
		{
			name:   "pair of twos",
			hand:   hand{card1: "two", card2: "two"},
			dealer: "ace",
			want:   "H",
		},
		{
			name:   "pair of fives",
			hand:   hand{card1: "five", card2: "five"},
			dealer: "ace",
			want:   "H",
		},
		{
			name:   "blackjack with ace for dealer",
			hand:   hand{card1: "ace", card2: "jack"},
			dealer: "ace",
			want:   "S",
		},
		{
			name:   "blackjack with queen for dealer",
			hand:   hand{card1: "king", card2: "ace"},
			dealer: "queen",
			want:   "S",
		},
		{
			name:   "blackjack with five for dealer",
			hand:   hand{card1: "ace", card2: "ten"},
			dealer: "five",
			want:   "W",
		},
		{
			name:   "blackjack with nine for dealer",
			hand:   hand{card1: "ace", card2: "king"},
			dealer: "nine",
			want:   "W",
		},
		{
			name:   "score of 20",
			hand:   hand{card1: "ten", card2: "king"},
			dealer: "ace",
			want:   "S",
		},
		{
			name:   "score of 19",
			hand:   hand{card1: "ten", card2: "nine"},
			dealer: "ace",
			want:   "S",
		},
		{
			name:   "score of 18",
			hand:   hand{card1: "ten", card2: "eight"},
			dealer: "ace",
			want:   "S",
		},
		{
			name:   "score of 17",
			hand:   hand{card1: "seven", card2: "king"},
			dealer: "ace",
			want:   "S",
		},
		{
			name:   "score of 16 with six for dealer",
			hand:   hand{card1: "ten", card2: "six"},
			dealer: "six",
			want:   "S",
		},
		{
			name:   "score of 16 with seven for dealer",
			hand:   hand{card1: "ten", card2: "six"},
			dealer: "seven",
			want:   "H",
		},
		{
			name:   "score of 16 with ace for dealer",
			hand:   hand{card1: "ten", card2: "six"},
			dealer: "ace",
			want:   "H",
		},
		{
			name:   "score of 15 with six for dealer",
			hand:   hand{card1: "ten", card2: "five"},
			dealer: "six",
			want:   "S",
		},
		{
			name:   "score of 15 with seven for dealer",
			hand:   hand{card1: "ten", card2: "five"},
			dealer: "seven",
			want:   "H",
		},
		{
			name:   "score of 15 with king for dealer",
			hand:   hand{card1: "ten", card2: "five"},
			dealer: "king",
			want:   "H",
		},
		{
			name:   "score of 14 with six for dealer",
			hand:   hand{card1: "ten", card2: "four"},
			dealer: "six",
			want:   "S",
		},
		{
			name:   "score of 14 with seven for dealer",
			hand:   hand{card1: "ten", card2: "four"},
			dealer: "seven",
			want:   "H",
		},
		{
			name:   "score of 14 with queen for dealer",
			hand:   hand{card1: "ten", card2: "four"},
			dealer: "queen",
			want:   "H",
		},
		{
			name:   "score of 13 with six for dealer",
			hand:   hand{card1: "ten", card2: "three"},
			dealer: "six",
			want:   "S",
		},
		{
			name:   "score of 13 with seven for dealer",
			hand:   hand{card1: "ten", card2: "three"},
			dealer: "seven",
			want:   "H",
		},
		{
			name:   "score of 13 with queen for dealer",
			hand:   hand{card1: "ten", card2: "three"},
			dealer: "queen",
			want:   "H",
		},
		{
			name:   "score of 12 with six for dealer",
			hand:   hand{card1: "ten", card2: "two"},
			dealer: "six",
			want:   "S",
		},
		{
			name:   "score of 12 with seven for dealer",
			hand:   hand{card1: "ten", card2: "two"},
			dealer: "seven",
			want:   "H",
		},
		{
			name:   "score of 12 with queen for dealer",
			hand:   hand{card1: "ten", card2: "two"},
			dealer: "queen",
			want:   "H",
		},
		{
			name:   "score of 11 with queen for dealer",
			hand:   hand{card1: "nine", card2: "two"},
			dealer: "queen",
			want:   "H",
		},
		{
			name:   "score of 10 with two for dealer",
			hand:   hand{card1: "eight", card2: "two"},
			dealer: "two",
			want:   "H",
		},
		{
			name:   "score of 5 with queen for dealer",
			hand:   hand{card1: "three", card2: "two"},
			dealer: "queen",
			want:   "H",
		},
		{
			name:   "score of 4 with five for dealer",
			hand:   hand{card1: "two", card2: "two"},
			dealer: "five",
			want:   "H",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstTurn(tt.hand.card1, tt.hand.card2, tt.dealer); got != tt.want {
				t.Errorf("FirstTurn(%s, %s, %s) = %s, want %s", tt.hand.card1, tt.hand.card2, tt.dealer, got, tt.want)
			}
		})
	}
}
