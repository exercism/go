package purchase

// NeedsLicence determines whether a licence is need to drive a type of vehicle. Only "car" and "truck" require a licence.
func NeedsLicence(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(first, second string) string {
	var choice string
	if first < second {
		choice = first
	} else {
		choice = second
	}
	return choice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var percentage float64
	if age < 3 {
		percentage = 0.8
	} else if age < 10 {
		percentage = 0.7
	} else {
		percentage = 0.5
	}
	return percentage * originalPrice
}
