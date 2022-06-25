# Instructions

Elaine is working on a new children's game that features animals and magic wands.
It is time to code functions for rolling a die, generating random wand energy and shuffling a slice.

## 1. Seed the random number generator.

Implement a `SeedWithTime` function that seeds the `math.rand` package with the current computer time.

## 2. Roll a die.

Implement a `RollADie` function.

This will be the traditional twenty-sided die with numbers 1 to 20.

```go
d := RollADie() // d will be assigned a random int, 1 <= d <= 20
```

## 3. Generate wand energy.

Implement a `GenerateWandEnergy` function.
The wand energy should be a random floating point number between 0.0 and 12.0.

```go
f := GenerateWandEnergy()  // f will be assigned a random float64, 0.0 <= f < 12.0
```

## 4. Shuffle a slice.

The game features eight different animals:

- ant
- beaver
- cat
- dog
- elephant
- fox
- giraffe
- hedgehog

Write a function `ShuffleAnimals` that returns a slice with the eight animals in random order.
