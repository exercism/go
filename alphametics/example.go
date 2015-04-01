package alphametics

import "sync"

const TestVersion = 1

// Set this to true to debug the example. It enables several useful logging tools.
const debugging = false

type accumulator struct {
	carry    int
	usedNums uint16       // bit n denotes if digit n is taken
	solution map[rune]int // partially filled solution map
}

type ipc struct {
	resultChan chan map[rune]int
	abortChan  chan struct{}
	waitGroup  *sync.WaitGroup
}

func Solve(a, b, c string) map[rune]int {
	aRunes, bRunes, cRunes := prepareRunes(a, b, c)

	// Setup the machinery for solving in parallel.
	resultChan := make(chan map[rune]int)
	abortChan := make(chan struct{}) // Closed when goroutines can abort

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		// This goroutine ensures the main thread will return nil if no answer can
		// be found instead of blocking forever.
		waitGroup.Wait()
		resultChan <- nil
	}()
	go solveStep(aRunes, bRunes, cRunes, accumulator{}, &ipc{resultChan, abortChan, &waitGroup})

	result := <-resultChan
	close(abortChan)
	return result
}

func solveStep(aRunes, bRunes, cRunes []rune, acc accumulator, ipc *ipc) {
	defer ipc.waitGroup.Done()
	if abortSignalled(ipc) {
		return
	}
	aRune, bRune, cRune := aRunes[0], bRunes[0], cRunes[0]
	c, cFound := acc.solution[cRune]
	for _, a := range availableDigits(aRune, acc, 0, 0) {
		for _, b := range availableDigits(bRune, acc, aRune, a) {
			if cFound {
				cCalc := (a + b + acc.carry) % 10
				if cCalc != c {
					continue
				}
			} else {
				c = (a + b + acc.carry) % 10
				if (aRune == cRune && aRune != 0 && c != a) ||
					(bRune == cRune && bRune != 0 && c != b) {
					continue
				}
				if isUsed(c, acc.usedNums) {
					continue
				}
			}
			newAcc := acc
			newAcc.carry = (a + b + acc.carry) / 10
			newUsed := acc.usedNums
			// Needed because availableDigits returns []int{0} if the rune is 0.
			if aRune != 0 {
				newUsed = setUsed(a, newUsed)
			}
			if bRune != 0 {
				newUsed = setUsed(b, newUsed)
			}
			newAcc.usedNums = setUsed(c, newUsed)
			newAcc.solution = copyMap(acc.solution)
			newAcc.solution[aRune] = a
			newAcc.solution[bRune] = b
			newAcc.solution[cRune] = c
			if len(aRunes) == 1 {
				// This was the final iteration, so we've found a result

				// Last check, the left-most (in our case last) character of each part
				// of the equation cannot be 0.
				if resultIsValid([][]rune{aRunes, bRunes, cRunes}, newAcc.solution) {
					ipc.resultChan <- newAcc.solution
				}
				return
			} else {
				ipc.waitGroup.Add(1)
				go solveStep(aRunes[1:], bRunes[1:], cRunes[1:], newAcc, ipc)
			}
		}
	}
}

func resultIsValid(runesSlice [][]rune, solution map[rune]int) bool {
outer:
	for _, runes := range runesSlice {
		for i := len(runes) - 1; i >= 0; i-- {
			if runes[i] != 0 {
				if solution[runes[i]] == 0 {
					return false
				} else {
					continue outer
				}
			}
		}
	}
	return true
}

// Gets the available digits to iterate over.
//
// NOTE: Returns []int{0} if 'char' is 0. This is convenient but don't forget to
// check if char == 0 before writing that 0 to the accumulator.
//
// If char == otherChar and otherChar isn't 0 (so a valid char), the otherDigit will be used.
// This is used to shortcut a situation where you have two equal letters.
func availableDigits(char rune, acc accumulator, otherChar rune, otherDigit int) []int {
	if char == 0 {
		return []int{0}
	}
	if otherChar != 0 && char == otherChar {
		return []int{otherDigit}
	}
	num, found := acc.solution[char]
	if found {
		return []int{num}
	}
	digits := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		if !isUsed(i, acc.usedNums) {
			if !(otherChar != 0 && otherChar != char && otherDigit == i) {
				digits = append(digits, i)
			}
		}
	}
	return digits
}

func isUsed(num int, usedNums uint16) bool {
	return usedNums&(1<<uint16(num)) != 0
}

func setUsed(num int, usedNums uint16) uint16 {
	return usedNums | 1<<uint16(num)
}

func copyMap(m map[rune]int) map[rune]int {
	result := make(map[rune]int, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

func abortSignalled(ipc *ipc) bool {
	select {
	case <-ipc.abortChan:
		return true
	default:
		return false
	}
}

// Turn the strings into zero rune padded rune slices in reverse order.
//
// This preparation simplifies the main logic.
func prepareRunes(a, b, c string) ([]rune, []rune, []rune) {
	// To simplify the logic ensure a is never longer than b.
	var aRunes, bRunes []rune
	if len(a) <= len(b) {
		aRunes = []rune(a)
		bRunes = []rune(b)
	} else {
		aRunes = []rune(b)
		bRunes = []rune(a)
	}
	cRunes := []rune(c)

	aPadded := make([]rune, len(cRunes))
	copy(aPadded[len(aPadded)-len(aRunes):], aRunes)
	bPadded := make([]rune, len(cRunes))
	copy(bPadded[len(bPadded)-len(bRunes):], bRunes)

	reverseRunes(aPadded)
	reverseRunes(bPadded)
	reverseRunes(cRunes)
	return aPadded, bPadded, cRunes
}

func reverseRunes(runes []rune) {
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
}
