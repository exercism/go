// Package affinecipher contains tools that implement affine-cipher.
package affinecipher

import (
	"errors"
	"strings"
)

type operation int

const (
	encode operation = iota + 1
	decode
)

const _totalAlphabets = 26 // total number of letters in alphabet.

// Encode encodes the provided message with the provided keys, using affine-cipher.
func Encode(text string, a, b int) (string, error) {
	if gcd(a, _totalAlphabets) != 1 {
		return "", errors.New("affinecipher.Encode: a and b must be co-prime")
	}
	return cipher(encode, text, a, b)
}

// Decode decodes the provided encoded message with the provided keys, using affine-cipher.
func Decode(text string, a, b int) (string, error) {
	if gcd(a, _totalAlphabets) != 1 {
		return "", errors.New("affinecipher.Decode: a and b must be co-prime")
	}
	mmi := multiInv(a, _totalAlphabets)
	return cipher(decode, text, mmi, b)
}

// cipher functions takes a operation, the text and as well as key values, and performs the
// operation on the text using the provided keys.
func cipher(op operation, text string, a, b int) (string, error) {
	text = strings.ToLower(text)
	var output strings.Builder
	accum := 0
	opFunc := operationFunc(op)
	for _, char := range []byte(text) {
		if !(char >= 'a' && char <= 'z') && // if not a letter and
			!(char >= '1' && char <= '9') { // if not a number
			continue
		}

		// if its encoding then write a space every 5 letters.
		if op == encode && accum != 0 && (accum%5) == 0 {
			output.WriteByte(' ')
		}
		accum++

		if char >= '1' && char <= '9' { // write numbers as it is.
			output.WriteByte(char)
			continue
		}
		output.WriteByte(opFunc(char, a, b))
	}
	return output.String(), nil
}

// gcd function finds the greatest common divisor for the numbers passed into it.
func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b > 0 {
		a, b = b, (a % b)
	}
	return a
}

// multiInv finds the modular multiple inverse of the passed num and mod.
func multiInv(num, mod int) int {
	t1, t2 := 0, 1
	a, b := num, mod

	if a < b {
		a, b = b, a
	}

	for b > 0 {
		t := t1 - (t2 * (a / b))
		a, b = b, (a % b)
		t1, t2 = t2, t
	}

	return t1
}

// operationFunc function returns the appropriate function for the passed operation.
func operationFunc(op operation) func(byte, int, int) byte {
	if op == encode {
		return encryptionFunc
	}
	return decryptionFunc
}

// encryptionFunc function is for encrypting text.
func encryptionFunc(letter byte, num1, num2 int) byte {
	temp := ((num1 * int(letter-'a')) + num2) % _totalAlphabets
	return byte(temp) + 'a'
}

// decryptionFunc function is for decrypting text.
func decryptionFunc(letter byte, mmi, num2 int) byte {
	temp := (mmi * (int(letter-'a') - num2)) % _totalAlphabets
	if temp < 0 {
		temp += _totalAlphabets
	}
	return byte(temp) + 'a'
}
