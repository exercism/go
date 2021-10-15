package elon

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

// TODO: define the 'Drive()' method

// TODO: define the 'CanFinish(distance int) bool' method

// TODO: define the 'DisplayDistance() string' method

// TODO: define the 'DisplayBattery() string' method
