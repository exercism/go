package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	var choice string
	if option1 < option2 {
		choice = option1
	} else {
		choice = option2
	}
	return choice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var percentage float64
	// Disables ifElseChain rule of gocritic, which suggests convert the following code to a switch,
	// however the point of this exercise is to teach if-else-if-else, not switch.
	//nolint:gocritic
	if age < 3 {
		percentage = 0.8
	} else if age < 10 {
		percentage = 0.7
	} else {
		percentage = 0.5
	}
	return percentage * originalPrice
}
