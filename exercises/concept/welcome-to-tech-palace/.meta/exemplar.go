package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder add a border to a welcome message.
func AddBorder(welcomeMessage string, int numStarsPerLine) string {
	stars := strings.Repeat("*", num)
	return stars + "\n" + text + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMessage string) string {
	msgWithoutStars := strings.ReplaceAll(message, "*", "")
	return strings.TrimSpace(msgWithoutStars)
}
