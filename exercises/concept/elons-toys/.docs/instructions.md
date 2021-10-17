# Instructions

Note: This exercise is a continuation of the `need-for-speed` exercise.

In this exercise you'll be organizing races between various types of remote controlled cars. Each car has its own speed and battery drain characteristics.

Cars start with full (100%) batteries. Each time you drive the car using the remote control, it covers the car's speed in meters and decreases the remaining battery percentage by its battery drain.

If a car's battery is below its battery drain percentage, you can't drive the car anymore.

The remote controlled car has a fancy LED display that shows two bits of information:

- The total distance it has driven, displayed as: `"Driven <METERS> meters"`.
- The remaining battery charge, displayed as: `"Battery at <PERCENTAGE>%"`.

Each race track has its own distance. Cars are tested by checking if they can finish the track without running out of battery.

## 1. Drive the car

Implement the `Drive` method on the `Car` that updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage:

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)
car.Drive()
// car is now Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
```

Note: If a car's battery is below its battery drain percentage, you can't drive the car anymore.

## 2. Display the distance driven

Implement a `DisplayDistance` method on `Car` to return the distance as displayed on the LED display as a `string`:

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

car.DisplayDistance()
// Output: "Driven 0 meters"
```

## 3. Display the battery percentage

Implement the `DisplayBattery` method on `Car` to return the distance as displayed on the LED display as a `string`:

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

car.DisplayBattery()
// Output: "Battery at 100%"
```

## 4. Check if a remote control car can finish a race

To finish a race, a car has to be able to drive the race's distance. This means not draining its battery before having crossed the finish line. Implement the `CanFinish` method that takes a `trackDistance int` as its parameter and returns `true` if the car can finish the race; otherwise, return `false`:

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

trackDistance := 100

car.CanFinish(trackDistance)
// Output: true
```
