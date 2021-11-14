package purchase

import (
	"testing"
)

func TestNeedsLicence(t *testing.T) {
	tests := []struct {
		name     string
		kind     string
		expected bool
	}{
		{
			name:     "need a licence for a car",
			kind:     "car",
			expected: true,
		},
		{
			name:     "need a licence for a truck",
			kind:     "truck",
			expected: true,
		},
		{
			name:     "does not need a licence for a bike",
			kind:     "bike",
			expected: false,
		},
		{
			name:     "does not need a licence for a stroller",
			kind:     "stroller",
			expected: false,
		},
		{
			name:     "does not need a licence for a e-scooter",
			kind:     "e-scooter",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := NeedsLicence(test.kind)
			if actual != test.expected {
				t.Errorf(
					"NeedsLicence(\"%s\") = %t, want %t",
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
		expected string
		other    string
	}{
		{
			name:     "chooses Bugatti over Ford",
			expected: "Bugatti Veyron",
			other:    "Ford Pinto",
		}, {
			name:     "chooses Chery over Kia",
			expected: "Chery EQ",
			other:    "Kia Niro Elektro",
		}, {
			name:     "chooses Ford Focus over Ford Pinto",
			expected: "Ford Focus",
			other:    "Ford Pinto",
		}, {
			name:     "chooses 2018 over 2020",
			expected: "2018 Bergamont City",
			other:    "2020 Gazelle Medeo",
		}, {
			name:     "chooses Bugatti over ford",
			expected: "Bugatti Veyron",
			other:    "ford",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expectedResponse := test.expected + " is clearly the better choice"
			actual := ChooseVehicle(test.expected, test.other)
			if actual != expectedResponse {
				t.Errorf(
					"ChooseVehicle(\"%s\", \"%s\") = \"%s\", want \"%s\"",
					test.expected,
					test.other,
					actual,
					expectedResponse)
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
			name:          "price is reduced to 70%% for age 10",
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
	}

	for _, test := range tests {
		actual := CalculateResellPrice(test.originalPrice, test.age)
		if actual != test.expected {
			t.Errorf(
				"CalculateResellPrice(%v, %v) = %v, want %v",
				test.originalPrice,
				test.age,
				actual,
				test.expected)
		}
	}
}
