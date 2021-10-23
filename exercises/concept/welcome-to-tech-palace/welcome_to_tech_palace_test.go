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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WelcomeMessage(tt.customer); got != tt.want {
				t.Errorf("WelcomeMessage(\"%s\") = \"%s\", want \"%s\"", tt.customer, got, tt.want)
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBorder(tt.welcomeMessage, tt.numStarsPerLine); got != tt.want {
				t.Errorf("AddBorder(\"%s\", %d) = \"%s\", want \"%s\"", tt.welcomeMessage, tt.numStarsPerLine, got, tt.want)
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
			name:       "Cleanup message without leading whitespace",
			oldMessage: "*****\n SALE\n*****",
			want:       "SALE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanupMessage(tt.oldMessage); got != tt.want {
				t.Errorf("CleanupMessage(\"%s\") = \"%s\", want \"%s\"", tt.oldMessage, got, tt.want)
			}
		})
	}
}
