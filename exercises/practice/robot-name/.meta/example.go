// Package robotname implements a type Robot with a unique Name field.

package robotname

import (
	"crypto/rand"
	"errors"
	"math/big"
)

type Robot struct {
	name string
}

const (
	upperChar = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digit     = "0123456789"
)

// digitSpec specifies valid digit values for a multi-base string from most
// significant digit to least significant digit.
var digitSpec = []string{
	upperChar,
	upperChar,
	digit,
	digit,
	digit,
}

// digitIncrements stores the increment value of each digit position for later
// use during the encoding step.
var digitIncrements = make([]int, len(digitSpec))

// MaxNames is the maximum number of names that can be generated based on the
// defined digitSpec.
var MaxNames int

// Keep track of unused integer IDs. These are intended to be used to
// deterministically generate new robot names, and then be removed from this
// slice.
var availableIDs []int

// init initializes values for various global variables.
func init() {
	var m int

	m = 1
	numDigits := len(digitSpec)
	for i := numDigits - 1; i >= 0; i-- {
		digitIncrements[i] = m
		m *= len(digitSpec[i])
	}
	MaxNames = m

	availableIDs = make([]int, MaxNames)
	for i := 0; i < MaxNames; i++ {
		availableIDs[i] = i
	}
}

// randIntn wraps the crypto/rand integer generator to return a random
// int.
func randIntn(i int) int {
	bigi := big.NewInt(int64(i))
	rndm, _ := rand.Int(rand.Reader, bigi)
	return int(rndm.Int64())
}

// Given an integer "ID", encode it into a multi-base string according to the
// given list of strings, digitSpec, where each string represents valid symbols
// for its digit position in order of increasing value. The digit spec
// describes digits in this way in order of most significant to least
// significant, ie index 0 is the most significant digit, index 1 is the next
// most significant, etc
func intToMultiBaseString(id int, digitSpec []string) string {
	var pos, rmndr, minVal int
	var result string

	rmndr = id
	for i, digitChars := range digitSpec {
		minVal = digitIncrements[i]

		// The character for this digit position is determined by the number of
		// times that the minimum/increment value of that digit can evenly
		// divide into the remainder from the last digit's calculations (or the
		// original ID in the case of the most significant digit).
		pos = rmndr / minVal
		result += string(digitChars[pos])

		// Calculate the remainder to be used for the next digit.
		rmndr %= minVal
	}

	return result
}

// getUniqueID returns a random unused ID and returns an error if none are left.
func getUniqueID() (int, error) {
	numAvailable := len(availableIDs)
	if numAvailable == 0 {
		return 0, errors.New("all possible names have been taken")
	}

	i := randIntn(numAvailable)
	uniqueID := availableIDs[i]

	// Delete the ith element by replacing it with the last element in the
	// array and cutting off the original last element.
	availableIDs[i] = availableIDs[numAvailable-1]
	availableIDs = availableIDs[:numAvailable-1]

	return uniqueID, nil
}

// generateUniqueName returns an error if there are no valid unique names left, returns a
// valid unique name otherwise.
func generateUniqueName() (string, error) {
	id, err := getUniqueID()
	if err != nil {
		return "", err
	}

	return intToMultiBaseString(id, digitSpec), nil
}

// Name returns the name of the robot.
func (r *Robot) Name() (string, error) {
	var err error
	if r.name == "" {
		r.name, err = generateUniqueName()
		if err != nil {
			return "", err
		}
	}
	return r.name, nil
}

// Reset resets the robot's name to empty string.
func (r *Robot) Reset() {
	*r = Robot{}
}
