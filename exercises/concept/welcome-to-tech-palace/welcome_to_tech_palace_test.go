package techpalace

import "testing"

func TestWelcomeMessage(t *testing.T) {
	tests := []struct {
		customer string
		expected string
	}{
		{
			customer: "Judy",
			expected: "Welcome to the Tech Palace, JUDY"
		},
		{
			customer: "Lars",
			expected: "Welcome to the Tech Palace, LARS"
		},
		{
			customer: "Peter-James",
			expected: "Welcome to the Tech Palace, PETER-JAMES"
		},
		{
			customer: "MJ",
			expected: "Welcome to the Tech Palace, MJ"
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
		welcomeMessage string
		numStarsPerLine int
		expected string
	}{
		{
			welcomeMessage: "Welcome!",
			numStarsPerLine: 10,
			expected: "**********\nWelcome!\n**********"
		},
		{
			welcomeMessage: "Hi",
			numStarsPerLine: 2,
			expected: "**\nHi\n**"
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
		oldMessage string
		expected string
	}{
		{
			oldMessage: "**************************\n*    BUY NOW, SAVE 10%   *\n**************************",
			expected: "BUY NOW, SAVE 10%"
		},
		{
			oldMessage: "**********\n*DISCOUNT*\n**********",
			expected: "DISCOUNT"
		},
		{
			oldMessage: "*****\n SALE\n*****",
			expected: "SALE"
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
