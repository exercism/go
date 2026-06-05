package lasagnamaster

import (
	"math"
	"reflect"
	"testing"
)

type preparationTimeTests struct {
	name           string
	layers         []string
	time, expected int
}

func TestPreparationTime(t *testing.T) {
	tests := []preparationTimeTests{
		{
			name: "Preparation time for many layers with custom average time",
			layers: []string{
				"sauce",
				"noodles",
				"béchamel",
				"meat",
				"mozzarella",
				"noodles",
				"ricotta",
				"eggplant",
				"béchamel",
				"noodles",
				"sauce",
				"mozzarella",
			},
			time:     1,
			expected: 12,
		},
		{
			name: "Preparation time for few layers",
			layers: []string{
				"sauce",
				"noodles",
			},
			time:     3,
			expected: 6,
		},
		{
			name: "Preparation time for default case",
			layers: []string{
				"sauce",
				"noodles",
			},
			time:     0,
			expected: 4,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := PreparationTime(tc.layers, tc.time); got != tc.expected {
				t.Errorf("PreparationTime(%v, %d) = %d; want %d", tc.layers, tc.time, got, tc.expected)
			}
		})
	}
}

type quantitiesTest struct {
	name       string
	layers     []string
	expNoodles int
	expSauce   float64
}

func TestQuantities(t *testing.T) {
	tests := []quantitiesTest{
		{
			name:       "few layers",
			layers:     []string{"noodles", "sauce", "noodles"},
			expNoodles: 100,
			expSauce:   0.2,
		},
		{
			name: "many layers",
			layers: []string{
				"sauce",
				"noodles",
				"béchamel",
				"meat",
				"mozzarella",
				"noodles",
				"ricotta",
				"eggplant",
				"béchamel",
				"noodles",
				"sauce",
				"mozzarella",
			},
			expNoodles: 150,
			expSauce:   0.4,
		},
		{
			name: "no noodles",
			layers: []string{
				"sauce",
				"meat",
				"mozzarella",
				"sauce",
				"mozzarella",
			},
			expNoodles: 0,
			expSauce:   0.4,
		},
		{
			name: "no sauce",
			layers: []string{
				"noodles",
				"meat",
				"mozzarella",
				"noodles",
				"mozzarella",
			},
			expNoodles: 100,
			expSauce:   0.0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotNoodles, gotSauce := Quantities(tc.layers)
			if gotNoodles != tc.expNoodles {
				t.Errorf("quantities(%v) = %d noodles; want %d", tc.layers, gotNoodles, tc.expNoodles)
			}
			if gotSauce != tc.expSauce {
				t.Errorf("quantities(%v) = %f sauce; want %f", tc.layers, gotSauce, tc.expSauce)
			}
		})
	}
}

type secretTest struct {
	name        string
	friendsList []string
	myList      []string
	expected    []string
}

func TestAddSecretIngredient(t *testing.T) {
	tests := []secretTest{
		{
			name:        "Adds secret ingredient",
			friendsList: []string{"sauce", "noodles", "béchamel", "marjoram"},
			myList:      []string{"sauce", "noodles", "meat", "tomatoes", "?"},
			expected:    []string{"sauce", "noodles", "meat", "tomatoes", "marjoram"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			friendsList := make([]string, len(tc.friendsList))
			copy(friendsList, tc.friendsList)
			myList := make([]string, len(tc.myList))
			copy(myList, tc.myList)
			AddSecretIngredient(tc.friendsList, tc.myList)
			if !reflect.DeepEqual(tc.myList, tc.expected) {
				t.Errorf("addSecretIngredient(%v, %v) = %v want %v", tc.friendsList, myList, tc.myList, tc.expected)
			}
			if !reflect.DeepEqual(friendsList, tc.friendsList) {
				t.Errorf("addSecretIngredient permuted friendsList (was %v, now %v), should not alter inputs", tc.friendsList, friendsList)
			}
		})
	}
}

type scaleRecipeTest struct {
	name     string
	input    []float64
	portions int
	expected []float64
}

func TestScaleRecipe(t *testing.T) {
	tests := []scaleRecipeTest{
		{
			name:     "scales up correctly",
			input:    []float64{0.5, 250, 150, 3, 0.5},
			portions: 6,
			expected: []float64{1.5, 750, 450, 9, 1.5},
		},
		{
			name:     "scales up correctly (2)",
			input:    []float64{0.6, 300, 1, 0.5, 50, 0.1, 100},
			portions: 3,
			expected: []float64{0.9, 450, 1.5, 0.75, 75, 0.15, 150},
		},
		{
			name:     "scales down correctly",
			input:    []float64{0.5, 250, 150, 3, 0.5},
			portions: 1,
			expected: []float64{0.25, 125, 75, 1.5, 0.25},
		},
		{
			name:     "empty recipe",
			input:    []float64{},
			portions: 100,
			expected: []float64{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inputList := make([]float64, len(tc.input))
			copy(inputList, tc.input)
			got := ScaleRecipe(inputList, tc.portions)
			if len(got) != len(tc.expected) {
				t.Errorf("ScaleRecipe(%v, %d) produced slice of length %d, expected %d", inputList, tc.portions, len(got), len(tc.expected))
			}
			for i := range tc.expected {
				if math.Abs(got[i]-tc.expected[i]) > 0.000001 {
					t.Errorf("Got %f Expected %f for index %d", got[i], tc.expected[i], i)
				}
			}
			if !reflect.DeepEqual(inputList, tc.input) {
				t.Errorf("ScaleRecipe permuted list (was %v, now %v), should not alter inputs", tc.input, inputList)
			}
		})
	}
}
