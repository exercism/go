package structs

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int
	battery  int
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}


// Track implements a race track.
type Track struct{
	distance int
}


// NewTrack created a new track
func NewTrack(distance int) Track{
	panic("Please implement the NewTrack function") 
}

// Drive drives the car one time. If there is not enough battry to drive on more time,
// the car will not move but use the leftover battery.
func Drive(car Car) Car {
	panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	panic("Please implement the CanFinish function")
}
