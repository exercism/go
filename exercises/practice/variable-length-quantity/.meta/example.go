package variablelengthquantity

import (
	"errors"
	"slices"
)

const (
	// 8th bit is used to indicate if there is more to this number.
	moreBit byte = 1 << 7
	// The lower 7 bits contain the numeric value.
	byteMask byte = moreBit - 1
)

// EncodeVarint encodes intergers into bytes.
func EncodeVarint(input []uint32) []byte {
	var out []byte
	// Add one number at a time.
	for _, number := range input {
		var encoded []byte
		// Add one byte for each set of lower 7 bits. Enable the more bit.
		for i := number; i > 0; i >>= 7 {
			encoded = append(encoded, moreBit | (byte(i) & byteMask))
		}
		if number == 0 {
			encoded = append(encoded, 0)
		}
		// The last byte should not have the more bit set.
		encoded[0] &= byteMask
		slices.Reverse(encoded)
		out = append(out, encoded...)
	}
	return out
}

// DecodeVarint converts bytes to numbers.
func DecodeVarint(input []byte) ([]uint32, error) {
	var numbers []uint32
	// The last byte must not set the more bit.
	if input[len(input)-1] & moreBit != 0 {
		return nil, errors.New("invalid input")
	}
	var number uint32
	for _, val := range input {
		// Collect the bytes into one number.
		number = (number << 7) | uint32(byteMask & val)
		// When there is no more bytes, add the number to the return slice.
		if (moreBit & val) == 0 {
			numbers = append(numbers, number)
			number = 0
		}
	}
	return numbers, nil
}
