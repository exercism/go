package elon

import "fmt"

// Drive drives the car one time. If there is not enough battry to drive on more time,
// the car will not move but use the leftover battery.
func (car *Car) Drive() {
	if car.battery < car.batteryDrain {
		return
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain
}

// DisplayDistance displays the distance the car is driven.
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery displays the battery level.
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish checks if a car is able to finish a track of a certain distance.
func (car *Car) CanFinish(trackDistance int) bool {
	maxDistance := car.battery / car.batteryDrain * car.speed
	return trackDistance <= maxDistance
}
