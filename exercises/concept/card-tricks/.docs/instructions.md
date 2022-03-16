# Instructions

As a magician-to-be, Elyse needs to practice some basics. She has a stack of cards that she wants to manipulate.

To make things a bit easier she only uses the cards 1 to 10.

## 1. Create a slice with all the cards 

Return a slice with all cards from 1 to 10 

```go 
cards := AllItems()
fmt.Println(cards) 
// Output: [1 2 3 4 5 6 7 8 9 10]
``` 

## 2. Retrieve a card from a stack

Return the card at position `index` from the given stack.

```go
card := GetItem([]int{1, 2, 4, 1}, 2) // card == 4  
```

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to return `-1`:

```go
card := GetItem([]int{1, 2, 4, 1}, 10) // card == -1  
```

## 3. Exchange a card in the stack

Exchange the card at position `index` with the new card provided and return the adjusted stack.
Note that this will modify the input slice which is the expected behavior. 

```go
index := 2
newCard := 6
cards := SetItem([]int{1, 2, 4, 1}, index, newCard)  
fmt.Println(cards) 
// Output: [1 2 6 1]
```

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to append the new card to the end of the stack:

```go
index := -1
newCard := 6
cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
fmt.Println(cards) 
// Output: [1 2 4 1 6]
```

## 4. Remove a card from the stack

Remove the card at position `index` from the stack and return the stack. Note
that this may modify the input slice which is ok. 

```go 
cards := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
fmt.Println(cards) 
// Output: [3 2 4 8]
```

If the index is out of bounds (ie. if it is negative or after the end of the
stack), we want to leave the stack unchanged:

```go
cards := RemoveItem([]int{3, 2, 6, 4, 8}, 11)
fmt.Println(cards) 
// Output: [3 2 6 4 8]
```

## 5. PrependItems

Add values at the front of a stack. Note that `PrependItems` is a variadic
function, its type is: 

```go 
func PrependItems(slice []int, value ...int) []int
``` 

The function can be called with an arbitrary number of arguments for the `value` parameter:  

```go
slice := []int{3, 2, 6, 4, 8}
cards := PrependItems(slice, 5, 1) 
fmt.Println(cards)  
// Output: [5 1 3 2 6 4 8]
```

If no argument is given for the `value` parameter, then the result equals the
original slice. 

```go 
slice := []int{3, 2, 6, 4, 8}
cards := PrependItems(slice) 
fmt.Println(cards)  
// Output: [3 2 6 4 8]
``` 
 