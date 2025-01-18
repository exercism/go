# Hints

## General

- A `struct` is a sequence of named elements, called fields, each one having a name and a type.

## 1. Create a remote controlled car

- Instantiate a new [struct][struct] and fill the fields except `distance`.
- Create a function that returns the type of the newly created struct and update the fields accordingly

## 2. Create a race track

- Define a new [struct][struct] with 1 field
- Create a function that returns the type of the new created struct and updates the fields accordingly

## 3. Drive the car

- Add a new field to keep track of the distance driven
- You need to check if there is enough battery to drive the car

## 4. Check if a remote controlled car can finish a race

- Assume the car is just starting the race
- You need to calculate the maximum distance a car can drive with the current level of battery
- The number of times a car can be driven can be calculated by `battery / batteryDrain`.
- The maximum distance the car can cover is the product of the car's speed and the number of times it can be driven.
- Knowing the maximum distance the car can drive, compare it with the distance of the race track

[struct]: https://tour.golang.org/moretypes/2
