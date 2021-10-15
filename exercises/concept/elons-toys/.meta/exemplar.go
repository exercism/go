package elon

import "fmt"

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
	}
}

// Drive drives the car one time. If there is not enough battry to drive on more time,
// the car will not move but use the leftover battery.
func (car *Car) Drive() {
	if car.battery < car.batteryDrain {
		return
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain
}

// CanFinish checks if a car is able to travel a certain distance
func (car *Car) CanFinish(distance int) bool {
	maxDistance := car.battery / car.batteryDrain * car.speed
	return distance <= maxDistance
}

// DisplayDistance displays the distance the car is driven.
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery displays the battery level.
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}
