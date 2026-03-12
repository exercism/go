package bafflingbirthdays

import (
	rand "math/rand/v2"
	"time"
)

func SharedBirthday(dates []time.Time) bool {
	seen := make(map[int]bool)
	for _, date := range dates {
		mmdd := int(date.Month())*100 + date.Day()
		if seen[mmdd] {
			return true
		}
		seen[mmdd] = true
	}
	return false
}

func RandomBirthdates(size int) []time.Time {
	var birthdates []time.Time
	base := time.Date(1999, 1, 1, 0, 0, 0, 0, &time.Location{})
	for range size {
		date := base.Add(time.Hour * 24 * time.Duration(rand.IntN(365)))
		birthdates = append(birthdates, date)
	}
	return birthdates
}

func EstimatedProbability(size int) float64 {
	var shared float64
	rounds := 1024 * 8
	for range rounds {
		if SharedBirthday(RandomBirthdates(size)) {
			shared++
		}
	}
	return 100 * shared / float64(rounds)
}
