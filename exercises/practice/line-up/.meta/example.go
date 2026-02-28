package lineup

import "fmt"

func Format(name string, number int) string {
	suffix := "th"
	if number % 100 / 10 != 1 {
		switch number % 10 {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		}
	}
	return fmt.Sprintf("%s, you are the %d%s customer we serve today. Thank you!", name, number, suffix)
}
