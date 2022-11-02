package speed
// TODO: define the 'Car' type struct
type Car struct{
    battery int
    batteryDrain int
    speed int
    distance int
}
// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	var NewTrack Car
    NewTrack.speed=speed
    NewTrack.batteryDrain=batteryDrain
    NewTrack.battery = 100
    NewTrack.distance = 0
    return NewTrack
}
// TODO: define the 'Track' type struct
type Track struct{
    distance int
}
// NewTrack creates a new track
func NewTrack(distance int) Track {
	var track Track
    track.distance=distance
    return track
}
// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery >= car.batteryDrain{
       car.distance+=car.speed
 	   car.battery -=  car.batteryDrain  
    }
	return car
}
// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	for car.battery >= car.batteryDrain{
       car.distance+=car.speed
 	   car.battery -=  car.batteryDrain  
    }
    if car.distance >= track.distance{
        return true
    } 
    return false
}
