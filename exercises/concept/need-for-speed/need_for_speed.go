package speed

// TODO: define the 'Car' type struct

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	panic("Please implement the NewCar function")
}

// TODO: define the 'Track' type struct

// NewTrack created a new track
func NewTrack(distance int) Track {
	panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	panic("Please implement the CanFinish function")
}
