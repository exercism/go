package purchase

// NeedsLicence determines whether a licence is need to drive a type of vehicle. Only "car" and "truck" require a licence.
func NeedsLicence(kind string) bool {
	panic("NeedsLicence not implemented")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	panic("CalculateResellPrice not implemented")
}
