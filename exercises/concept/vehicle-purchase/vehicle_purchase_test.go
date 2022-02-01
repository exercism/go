package purchase

import (
	"math"
	"testing"
)

const floatEqualityThreshold = 1e-5

func TestNeedsLicense(t *testing.T) {
	tests := []struct {
		name     string
		kind     string
		expected bool
	}{
		{
			name:     "need a license for a car",
			kind:     "car",
			expected: true,
		},
		{
			name:     "need a license for a truck",
			kind:     "truck",
			expected: true,
		},
		{
			name:     "does not need a license for a bike",
			kind:     "bike",
			expected: false,
		},
		{
			name:     "does not need a license for a stroller",
			kind:     "stroller",
			expected: false,
		},
		{
			name:     "does not need a license for a e-scooter",
			kind:     "e-scooter",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := NeedsLicense(test.kind)
			if actual != test.expected {
				t.Errorf(
					"NeedsLicense(%q) = %t, want %t",
					test.kind,
					actual,
					test.expected)
			}
		})
	}
}

func TestChooseVehicle(t *testing.T) {
	tests := []struct {
		name     string
		choice1  string
		choice2  string
		expected string
	}{
		{
			name:     "chooses Bugatti over Ford",
			choice1:  "Bugatti Veyron",
			choice2:  "Ford Pinto",
			expected: "Bugatti Veyron is clearly the better choice.",
		}, {
			name:     "chooses Chery over Kia",
			choice1:  "Chery EQ",
			choice2:  "Kia Niro Elektro",
			expected: "Chery EQ is clearly the better choice.",
		}, {
			name:     "chooses Ford Focus over Ford Pinto",
			choice1:  "Ford Focus",
			choice2:  "Ford Pinto",
			expected: "Ford Focus is clearly the better choice.",
		}, {
			name:     "chooses 2018 over 2020",
			choice1:  "2018 Bergamont City",
			choice2:  "2020 Gazelle Medeo",
			expected: "2018 Bergamont City is clearly the better choice.",
		}, {
			name:     "chooses Bugatti over ford",
			choice1:  "Bugatti Veyron",
			choice2:  "ford",
			expected: "Bugatti Veyron is clearly the better choice.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ChooseVehicle(test.choice1, test.choice2)
			if actual != test.expected {
				t.Errorf(
					"ChooseVehicle(%q, %q) = %q, want %q",
					test.choice1,
					test.choice2,
					actual,
					test.expected)
			}

			// The result should be independent of the order in which the choices are given.
			actual = ChooseVehicle(test.choice2, test.choice1)
			if actual != test.expected {
				t.Errorf(
					"ChooseVehicle(%q, %q) = %q, want %q",
					test.choice2,
					test.choice1,
					actual,
					test.expected)
			}
		})
	}
}

func TestCalculateResellPrice(t *testing.T) {
	tests := []struct {
		name          string
		originalPrice float64
		age           float64
		expected      float64
	}{
		{
			name:          "price is reduced to 80%% for age below 3",
			originalPrice: 40000,
			age:           2,
			expected:      32000,
		},
		{
			name:          "price is reduced to 80%% for age below 3",
			originalPrice: 40000,
			age:           2.5,
			expected:      32000,
		},
		{
			name:          "price is reduced to 70%% for age 7",
			originalPrice: 40000,
			age:           7,
			expected:      28000,
		},
		{
			name:          "price is reduced to 50%% for age 10",
			originalPrice: 25000,
			age:           10,
			expected:      12500,
		},
		{
			name:          "price is reduced to 50%% for age 11",
			originalPrice: 50000,
			age:           11,
			expected:      25000,
		},
		{
			name:          "float price is reduced to 70%% for age 8,",
			originalPrice: 39000.000001,
			age:           8,
			expected:      27300.0000007,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := CalculateResellPrice(test.originalPrice, test.age)
			if !floatingPointEquals(actual, test.expected) {
				t.Errorf(
					"CalculateResellPrice(%v, %v) = %v, want %v",
					test.originalPrice,
					test.age,
					actual,
					test.expected)
			}
		})
	}
}

func floatingPointEquals(got, want float64) bool {
	absoluteDifferenceBelowThreshold := math.Abs(got-want) <= floatEqualityThreshold
	relativeDifferenceBelowThreshold := math.Abs(got-want)/(math.Abs(got)+math.Abs(want)) <= floatEqualityThreshold
	return absoluteDifferenceBelowThreshold || relativeDifferenceBelowThreshold
}
