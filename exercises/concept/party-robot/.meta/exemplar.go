package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and stands out his age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbour string, direction string, distance float64) string {
	info := Welcome(name) + "\n"
	info += fmt.Sprintf("You have been assigned to table %X. ", table)
	info += fmt.Sprintf("Your table is %s, exactly %.1f meters from here.\n", direction, distance)
	info += fmt.Sprintf("You will be sitting next to %s.", neighbour)
	return info
}
