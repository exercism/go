package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

// Number takes in a potential phone number string and returns the number
// without any formatting if it's valid.
//
// * If the phone number is less than 10 digits assume that it is bad number
// * If the phone number is 10 digits assume that it is good
// * If the phone number is 11 digits and the first number is 1, trim the 1 and use the first 10 digits
// * If the phone number is 11 digits and the first number is not 1, then it is a bad number
// * If the phone number is more than 11 digits assume that it is a bad number
func Number(s string) (string, error) {
	//remove any characters that aren't digits
	clean := strings.Map(
		func(r rune) rune {
			if unicode.IsDigit(r) {
				return r
			}
			return -1
		}, s)

	numDigits := len(clean)
	switch {
	// bad number if less than 10 digits
	case numDigits < 10:
		return "", fmt.Errorf("Phone number %q contains too few digits", s)
	// good number if exactly 10 digits
	case numDigits == 10:
		return clean, nil
	// good number if 11 digits and 1st digit is a 1.  Return last 10 digits.
	case numDigits == 11 && clean[0] == '1':
		return clean[1:], nil
	}
	// all other numbers are a bad
	return "", fmt.Errorf("Phone number %q contains too many digits", s)
}

// AreaCode takes in a phone number string and returns the area code (first three digits)
func AreaCode(s string) (string, error) {
	num, err := Number(s)
	if err != nil {
		return "", err
	}
	return num[:3], nil
}

// Format takes in a phone number string and returns it as a well formatted
// human readable string
func Format(s string) (string, error) {
	f, err := Number(s)
	if err != nil {
		return "", err
	}
	// should follow the sample format of: (123) 456-7890
	return fmt.Sprintf("(%s) %s-%s", f[:3], f[3:6], f[6:]), nil
}
