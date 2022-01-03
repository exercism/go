package lasagna

// PreparationTim estimates the preparation time based on the number of layers and an average time per layer and returns it..
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}
	return len(layers) * avgPrepTime
}

// Quantities calculates and returns how many noodles and much sauce are needed for the given layers.
func Quantities(layers []string) (noodles int, sauce float64) {
	numLayers := len(layers)
	for i := 0; i < numLayers; i++ {
		if layers[i] == "noodles" {
			noodles += 50
		}

		if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// AddSecretIngredient adds the secret ingredient from the ingredient list that a friend provided to your ingredient list.
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe makes a new slice of float64s from an input slice scaled by a number of portions.
func ScaleRecipe(list []float64, portions int) []float64 {
	output := make([]float64, len(list))
	for i := 0; i < len(list); i++ {
		output[i] = list[i] * float64(portions) / 2
	}
	return output
}
