# Instructions

In this exercise you are going to write some code to help you prepare to buy a vehicle.

You have three tasks, one to determine if you need a licence, one to help you choose between two vehicles and one to estimate the acceptable price for a used vehicle.

## 1. Determine if you will need a driver's licence

Some vehicle kinds require a driver's licence to operate them.
Assume only the kinds `'car'` and `'truck'` require a licence, everything else can be operated without a licence.

Implement the `NeedsLicence(kind)` function that takes the kind of vehicle and returns a boolean indicating whether you need a licence for that kind of vehicle.

```go
needLicence := NeedsLicence('car')
// Output: true

needLicence = NeedsLicence('bike')
// Output: false
```

## 2. Choose between two potential vehicles to buy

You evaluate your options of available vehicles.
You manage to narrow it down to two options but you need help making the final decision.
For that implement the function `ChooseVehicle(option1, option2)` that takes two vehicles as arguments and returns a decision that includes the option that comes first in dictionary order.

```go
vehicle := ChooseVehicle('Wuling Hongguang', 'Toyota Corolla')
// Output: "Toyota Corolla is clearly the better choice."

ChooseVehicle('Volkswagen Beetle', 'Volkswagen Golf')
// Output: "Volkswagen Beetle is clearly the better choice."
```

## 3. Calculate an estimation for the price of a used vehicle

Now that you made your decision you want to make sure you get a fair price at the dealership.
Since you are interested in buying a used vehicle, the price depends on how old the vehicle is.
For a rough estimate, assume if the vehicle is less than 3 years old, it costs 80% of the original price it had when it was brand new.
If it is more than 10 years old, it costs 50%.
If the vehicle is at least 3 years old but not older than 10 years, it costs 70% of the original price.

Implement the `CalculateResellPrice(originalPrice, age)` function that applies this logic using `if`, `else if` and `else` (there are other ways if you want to practice).
It takes the original price and the age of the vehicle as arguments and returns the estimated price in the dealership.

```go
CalculateResellPrice(1000, 1)
// Output: 800

CalculateResellPrice(1000, 5)
// Output: 700

CalculateResellPrice(1000, 15)
// Output: 500
```

Note the value returned is a `float64`.