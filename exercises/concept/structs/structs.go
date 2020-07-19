package structs

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	panic("Please implement the NewCar function")
}

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

// TODO: add race track and instantiator function

// Drive drives the car one time. If there is not enough battry to drive on more time,
// the car will not move but use the leftover battery.
func Drive(car Car) Car {
	panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
// TODO: uncomment to implement CanFinish
// func CanFinish(car Car, track Track) bool {
// 	panic("Please implement the CanFinish function")
// }
