package zebra

import "fmt"

/* The Zebra Puzzle facts given:

1. There are five houses.
2. The Englishman lives in the red house.
3. The Spaniard owns the dog.
4. Coffee is drunk in the green house.
5. The Ukrainian drinks tea.
6. The green house is immediately to the right of the ivory house.
7. The Old Gold smoker owns snails.
8. Kools are smoked in the yellow house.
9. Milk is drunk in the middle house.
10. The Norwegian lives in the first house.
11. The man who smokes Chesterfields lives in the house next to the man with the fox.
12. Kools are smoked in the house next to the house where the horse is kept.
13. The Lucky Strike smoker drinks orange juice.
14. The Japanese smokes Parliaments.
15. The Norwegian lives next to the blue house.

*/

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

// SolvePuzzle returns a solution answering the two questions for the zebra puzzle,
// which are "Who drinks water?", and "Who owns the zebra?"
func SolvePuzzle() (solution Solution) {

	// 1. There are five houses.
	const (
		firstHouse = iota // name/identity of each house
		secondHouse
		middleHouse
		fourthHouse
		fifthHouse
		numHouses
	)
	houses := []int{firstHouse, secondHouse, middleHouse, fourthHouse, fifthHouse}

	// Generate the permutations of all the house identities.
	houseIdentityPermutations := permutations(houses, numHouses)

	// Note: Below, the pattern of var naming of "iXyz" is notation for "potential house (i)dentity of Xyz"
	// So for example, when comparing pEnglishman == pRed, the test is determining whether the permutation
	// of placements for both the Englishman and the red house agree to meet the given fact #2.

	for _, houseColors := range houseIdentityPermutations {

		iRed, iGreen, iIvory, iYellow, iBlue := assign(houseColors)

		// 6. The green house is immediately to the right of the ivory house.
		if !justRightOf(iGreen, iIvory) {
			continue
		}

		for _, residents := range houseIdentityPermutations {

			iEnglishman, iSpaniard, iUkrainian, iNorwegian, iJapanese := assign(residents)

			// 2. The Englishman lives in the red house.
			// 10. The Norwegian lives in the first house.
			// 15. The Norwegian lives next to the blue house.
			if iEnglishman != iRed || iNorwegian != firstHouse || !nextTo(iNorwegian, iBlue) {
				continue
			}

			for _, beverages := range houseIdentityPermutations {

				iCoffee, iTea, iMilk, iOrangeJuice, iWater := assign(beverages)

				// 4. Coffee is drunk in the green house.
				// 5. The Ukrainian drinks tea.
				// 9. Milk is drunk in the middle house.
				if iCoffee != iGreen || iUkrainian != iTea || iMilk != middleHouse {
					continue
				}

				for _, smokeBrands := range houseIdentityPermutations {

					iOldGold, iKools, iChesterfields, iLuckyStrike, iParliaments := assign(smokeBrands)

					// 8. Kools are smoked in the yellow house.
					// 13. The Lucky Strike smoker drinks orange juice.
					// 14. The Japanese smokes Parliaments.
					if iKools != iYellow || iLuckyStrike != iOrangeJuice || iJapanese != iParliaments {
						continue
					}

					for _, pets := range houseIdentityPermutations {

						iDog, iSnails, iFox, iHorse, iZebra := assign(pets)

						// 3. The Spaniard owns the dog.
						// 7. The Old Gold smoker owns snails.
						// 11. The man who smokes Chesterfields lives in the house next to the man with the fox.
						// 12. Kools are smoked in the house next to the house where the horse is kept.
						if iSpaniard != iDog || iOldGold != iSnails ||
							!nextTo(iChesterfields, iFox) || !nextTo(iKools, iHorse) {
							continue
						}

						// At this point all criteria are met, so we arrived at the solution,
						// and can fill in an array h of house facts (actually we only need the residents).

						var h [numHouses]struct {
							resident   string
							color      string
							pet        string
							beverage   string
							smokeBrand string
						}

						h[iEnglishman].resident = "Englishman"
						h[iSpaniard].resident = "Spaniard"
						h[iUkrainian].resident = "Ukrainian"
						h[iJapanese].resident = "Japanese"
						h[iNorwegian].resident = "Norwegian"

						solution = Solution{
							DrinksWater: h[iWater].resident,
							OwnsZebra:   h[iZebra].resident}

						if !showHouseFacts {
							return
						}
						h[iRed].color = "red"
						h[iGreen].color = "green"
						h[iIvory].color = "ivory"
						h[iYellow].color = "yellow"
						h[iBlue].color = "blue"
						h[iDog].pet = "dog"
						h[iSnails].pet = "snails"
						h[iFox].pet = "fox"
						h[iHorse].pet = "horse"
						h[iZebra].pet = "zebra"
						h[iCoffee].beverage = "coffee"
						h[iTea].beverage = "tea"
						h[iMilk].beverage = "milk"
						h[iOrangeJuice].beverage = "orange juice"
						h[iWater].beverage = "water"
						h[iOldGold].smokeBrand = "OldGold"
						h[iKools].smokeBrand = "Kools"
						h[iChesterfields].smokeBrand = "Chesterfields"
						h[iLuckyStrike].smokeBrand = "LuckyStrike"
						h[iParliaments].smokeBrand = "Parliaments"
						var houseNames = [5]string{"first", "second", "middle", "fourth", "fifth"}

						for p := firstHouse; p <= fifthHouse; p++ {
							var separator string
							if h[p].pet[len(h[p].pet)-1] != 's' {
								separator = "a "
							}
							fmt.Printf("The %-10s lives in the %-6s house which is %-7s owns %-8s drinks %-13s and smokes %-13s\n",
								h[p].resident,
								houseNames[p],
								h[p].color+",",
								separator+h[p].pet+",",
								h[p].beverage+",",
								h[p].smokeBrand+".")
						}

						return
					}
				}
			}
		}
	}
	return
}

const showHouseFacts = false // when true, print all house facts for fun.

// assign is notational helper function which returns the successive members
// of a permutation slice as five distinct values.
func assign(p []int) (a, b, c, d, e int) {
	return p[0], p[1], p[2], p[3], p[4]
}

// justRightOf returns true if positionally x is just to the right of y.
func justRightOf(x, y int) bool {
	return (x - y) == 1
}

// nextTo returns true if positionally, x is next to y, differing only by 1.
func nextTo(x, y int) bool {
	return (x-y) == 1 || (y-x) == 1
}

// permutations returns a slice containing the r length permutations of the elements in iterable.
// The implementation is modeled after the Python intertools.permutations().
func permutations(iterable []int, r int) (perms [][]int) {
	pool := iterable
	n := len(pool)

	if r > n {
		return
	}

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	result := make([]int, r)
	for i, el := range indices[:r] {
		result[i] = pool[el]
	}

	p := make([]int, len(result))
	copy(p, result)
	perms = append(perms, p)

	for n > 0 {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				index := indices[i]
				for j := i; j < n-1; j++ {
					indices[j] = indices[j+1]
				}
				indices[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]

				for k := i; k < r; k++ {
					result[k] = pool[indices[k]]
				}

				p = make([]int, len(result))
				copy(p, result)
				perms = append(perms, p)

				break
			}
		}

		if i < 0 {
			return
		}

	}
	return
}
