# Instructions

In this exercise you are going to write some more code related to preparing and cooking your brilliant lasagna from your favorite cookbook.

You have four tasks.
The first one is related to the cooking itself, the other three are about the perfect preparation.

## 1. Estimate the preparation time

For the next lasagna that you will prepare, you want to make sure you have enough time reserved so you can enjoy the cooking.
You already planned which layers your lasagna will have.
Now you want to estimate how long the preparation will take based on that.

Implement a function `PreparationTime` that accepts an array of layers as a `[]string` and the average preparation time per layer in minutes as an `int`.
The function should return the estimate for the total preparation time based on the number of layers as an `int`.
Go has no default values for functions.
If the average preperation time is passed as `0` (the default initial value for an `int`), then the default value of `2` should be used.

```go
layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
PreparationTime(layers, 3)
// => 18
PreparationTime(layers, 0)
// => 12
```

## 2. Compute the amounts of noodles and sauce needed

Besides reserving the time, you also want to make sure you have enough sauce and noodles to cook the lasagna of your dreams.
For each noodle layer in your lasagna, you will need 50 grams of noodles.
For each sauce layer in your lasagna, you will need 0.2 liters of sauce.

Define the function `Quantities` that takes an array of layers as parameter as a `[]string`.
The function will then determine the quantity of noodles and sauce needed to make your meal.
The result should be returned as two values of `noodles` as an `int` and `sauce` as a `float64`.

```go
Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"})
// => 100, 0.4
```

## 3. Add the secret ingredient

A while ago you visited a friend and ate lasagna there.
It was amazing and had something special to it.
The friend sent you the list of ingredients and told you the last item on the list is the "secret ingredient" that made the meal so special.
Now you want to add that secret ingredient to your recipe as well.

Write a function `AddSecretIngredient` that accepts two arrays of ingredients of type `[]string` as parameters.
The first parameter is the list your friend sent you, the second is the ingredient list for your own recipe.
The function should generate a new slice and add the last item from your friends list to the end of your list.
Neither argument should not be modified.

```go
friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
myList := []string{"noodles", "meat", "sauce", "mozzarella"}

AddSecretIngredient(friendsList, myList)
// => []string{"noodles", "meat", "sauce", "mozzarella", "kampot pepper"}
```

## 4. Scale the recipe

The amounts listed in your cookbook only yield enough lasagna for two portions.
Since you want to cook for more people next time, you want to calculate the amounts for different numbers of portions.

Implement a function `ScaleRecipe` that takes two parameters.

- A slice of `float64` amounts needed for 2 portions.
- The number of portions you want to cook.

The function should return a slice of `float64` of the amounts needed for the desired number of portions.
You want to keep the original recipe though.
This means the `amounts` argument should not be modified in this function.

```go
quantities := []float64{ 1.2, 3.6, 10.5 }
scaledQuantities := ScaleRecipe(quantities, 4)
// => []float64{ 2.4, 7.2, 21 }
```
