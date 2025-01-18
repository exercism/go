package lasagna

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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreparationTime(tt.layers, tt.time); got != tt.expected {
				t.Errorf("PreparationTime(%v, %d) = %d; want %d", tt.layers, tt.time, got, tt.expected)
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
				"mozzarella"},
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
				"mozzarella"},
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
				"mozzarella"},
			expNoodles: 100,
			expSauce:   0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNoodles, gotSauce := Quantities(tt.layers)
			if gotNoodles != tt.expNoodles {
				t.Errorf("quantities(%v) = %d noodles; want %d", tt.layers, gotNoodles, tt.expNoodles)
			}
			if gotSauce != tt.expSauce {
				t.Errorf("quantities(%v) = %f sauce; want %f", tt.layers, gotSauce, tt.expSauce)
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			friendsList := make([]string, len(tt.friendsList))
			copy(friendsList, tt.friendsList)
			myList := make([]string, len(tt.myList))
			copy(myList, tt.myList)
			AddSecretIngredient(tt.friendsList, tt.myList)
			if !reflect.DeepEqual(tt.myList, tt.expected) {
				t.Errorf("addSecretIngredient(%v, %v) = %v want %v", tt.friendsList, myList, tt.myList, tt.expected)
			}
			if !reflect.DeepEqual(friendsList, tt.friendsList) {
				t.Errorf("addSecretIngredient permuted friendsList (was %v, now %v), should not alter inputs", tt.friendsList, friendsList)
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputList := make([]float64, len(tt.input))
			copy(inputList, tt.input)
			got := ScaleRecipe(inputList, tt.portions)
			if len(got) != len(tt.expected) {
				t.Errorf("ScaleRecipe(%v, %d) produced slice of length %d, expected %d", inputList, tt.portions, len(got), len(tt.expected))
			}
			for i := range tt.expected {
				if math.Abs(got[i]-tt.expected[i]) > 0.000001 {
					t.Errorf("Got %f Expected %f for index %d", got[i], tt.expected[i], i)
				}
			}
			if !reflect.DeepEqual(inputList, tt.input) {
				t.Errorf("ScaleRecipe permuted list (was %v, now %v), should not alter inputs", tt.input, inputList)
			}
		})
	}
}
