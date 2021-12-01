# Instructions

In this exercise you'll be writing code to analyze the production in a car factory.

## 1. Calculate the number of working cars produced per hour

The cars are produced on an assembly line. 
The assembly line has a certain speed, that can be changed. 
The faster the assembly line speed is, the more cars are produced. 
However, changing the speed of the assembly line also changes the number of cars that are produced successfully, that is cars without any errors in their production.

Implement a function that takes in the number of cars produced per hour and the success rate and calculates the number of successful cars made per hour. The success rate is given as a percentage, from `0` to `100`:

```go
rate := CalculateWorkingCarsPerHour(1547, 90)
// Output: 1392.3
```

**Note** the return value should be a `float64`.

## 2. Calculate the number of working cars produced per minute

Implement a function that takes in the number of cars produced per hour and the success rate and calculates how many cars are successfully produced each minute:

```go
rate := CalculateWorkingCarsPerMinute(1105, 90)
// Output: 16
```

**Note** the return value should be an `int`.

## 3. Calculate the cost of production 

Each car normally costs $10,000 to produce, regardless of whether it is successful or not.
But with a bit of planning, the 3 cars can be produced together for $29,000.
10 cars can be produced together for $95,000.

For example, 37 cars can be produced in the following groups:
37 = 3 x groups of 10 + 2 groups of 3 + 1 

So the cost for 37 cars is:
3*95,000+2*29,000+10,000=353,000

Implement the function `CalculateCost` that calculates the cost of producing a number of cars, regardless of whether they are successful:

```go
cost := CalculateCost(37)
// Output: 353,000

cost = CalculateCost(21)
// Output: 200,000
```

**Note** the return value should be an `uint`.
