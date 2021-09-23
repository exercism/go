package techpalace

import "testing"

func TestWelcomeMessage(t *testing.T) {
	tests := []struct {
		name     string
		customer string
		expected string
	}{
		{
			name:     "Welcome message for customer with first letter capitalized",
			customer: "Judy",
			expected: "Welcome to the Tech Palace, JUDY",
		},
		{
			name:     "Welcome message for customer with only lowercase letters",
			customer: "lars",
			expected: "Welcome to the Tech Palace, LARS",
		},
		{
			name:     "Welcome message for customer with dash in name",
			customer: "Peter-James",
			expected: "Welcome to the Tech Palace, PETER-JAMES",
		},
		{
			name:     "Welcome message for customer with only uppercase letters",
			customer: "MJ",
			expected: "Welcome to the Tech Palace, MJ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WelcomeMessage(tt.customer); got != tt.expected {
				t.Errorf("WelcomeMessage(\"%s\") = \"%s\", expected \"%s\"", tt.customer, got, tt.expected)
			}
		})
	}
}

func TestAddBorder(t *testing.T) {
	tests := []struct {
		name            string
		welcomeMessage  string
		numStarsPerLine int
		expected        string
	}{
		{
			name:            "Add border with 10 stars per line",
			welcomeMessage:  "Welcome!",
			numStarsPerLine: 10,
			expected:        "**********\nWelcome!\n**********",
		},
		{
			name:            "Add border with 2 stars per line",
			welcomeMessage:  "Hi",
			numStarsPerLine: 2,
			expected:        "**\nHi\n**",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBorder(tt.welcomeMessage, tt.numStarsPerLine); got != tt.expected {
				t.Errorf("AddBorder(\"%s\", %d) = \"%s\", expected \"%s\"", tt.welcomeMessage, tt.numStarsPerLine, got, tt.expected)
			}
		})
	}
}

func TestCleanupMessage(t *testing.T) {
	tests := []struct {
		name       string
		oldMessage string
		expected   string
	}{
		{
			name:       "Cleanup message with many stars and leading and trailing whitespace",
			oldMessage: "**************************\n*    BUY NOW, SAVE 10%   *\n**************************",
			expected:   "BUY NOW, SAVE 10%",
		},
		{
			name:       "Cleanup message without leading or trailing whitespace",
			oldMessage: "**********\n*DISCOUNT*\n**********",
			expected:   "DISCOUNT",
		},
		{
			name:       "Cleanup message without leading whitespace",
			oldMessage: "*****\n SALE\n*****",
			expected:   "SALE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanupMessage(tt.oldMessage); got != tt.expected {
				t.Errorf("CleanupMessage(\"%s\") = \"%s\", expected \"%s\"", tt.oldMessage, got, tt.expected)
			}
		})
	}
}
