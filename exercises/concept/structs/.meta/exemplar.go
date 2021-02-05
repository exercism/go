package _meta

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}


// NewTrack creates a new track with given distance.
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battry to drive on more time,
// the car will not move but use the leftover battery.
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	maxDistance := car.battery / car.batteryDrain * car.distance
	return track.distance <= maxDistance
}
