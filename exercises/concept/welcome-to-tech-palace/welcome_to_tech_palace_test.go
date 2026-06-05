package techpalace

import "testing"

func TestWelcomeMessage(t *testing.T) {
	tests := []struct {
		name     string
		customer string
		want     string
	}{
		{
			name:     "Welcome message for customer with first letter capitalized",
			customer: "Judy",
			want:     "Welcome to the Tech Palace, JUDY",
		},
		{
			name:     "Welcome message for customer with only lowercase letters",
			customer: "lars",
			want:     "Welcome to the Tech Palace, LARS",
		},
		{
			name:     "Welcome message for customer with dash in name",
			customer: "Peter-James",
			want:     "Welcome to the Tech Palace, PETER-JAMES",
		},
		{
			name:     "Welcome message for customer with only uppercase letters",
			customer: "MJ",
			want:     "Welcome to the Tech Palace, MJ",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := WelcomeMessage(tc.customer); got != tc.want {
				t.Errorf("WelcomeMessage(\"%s\") = \"%s\", want \"%s\"", tc.customer, got, tc.want)
			}
		})
	}
}

func TestAddBorder(t *testing.T) {
	tests := []struct {
		name            string
		welcomeMessage  string
		numStarsPerLine int
		want            string
	}{
		{
			name:            "Add border with 10 stars per line",
			welcomeMessage:  "Welcome!",
			numStarsPerLine: 10,
			want:            "**********\nWelcome!\n**********",
		},
		{
			name:            "Add border with 2 stars per line",
			welcomeMessage:  "Hi",
			numStarsPerLine: 2,
			want:            "**\nHi\n**",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := AddBorder(tc.welcomeMessage, tc.numStarsPerLine); got != tc.want {
				t.Errorf("AddBorder(\"%s\", %d) = \"%s\", want \"%s\"", tc.welcomeMessage, tc.numStarsPerLine, got, tc.want)
			}
		})
	}
}

func TestCleanupMessage(t *testing.T) {
	tests := []struct {
		name       string
		oldMessage string
		want       string
	}{
		{
			name:       "Cleanup message with many stars and leading and trailing whitespace",
			oldMessage: "**************************\n*    BUY NOW, SAVE 10%   *\n**************************",
			want:       "BUY NOW, SAVE 10%",
		},
		{
			name:       "Cleanup message without leading or trailing whitespace",
			oldMessage: "**********\n*DISCOUNT*\n**********",
			want:       "DISCOUNT",
		},
		{
			name:       "Cleanup message with leading whitespace",
			oldMessage: "*****\n SALE\n*****",
			want:       "SALE",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := CleanupMessage(tc.oldMessage); got != tc.want {
				t.Errorf("CleanupMessage(\"%s\") = \"%s\", want \"%s\"", tc.oldMessage, got, tc.want)
			}
		})
	}
}
