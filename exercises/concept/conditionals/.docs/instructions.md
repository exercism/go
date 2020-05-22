# Instructions

In this exercise we will simulate the first turn of a [Blackjack](https://en.wikipedia.org/wiki/Blackjack) game.

You will receive two cards and will be able to see the face up card of the dealer. All cards are represented using a string such as "ace", "king", "three", "two", etc. The values of each card are:

| card  | value | card  | value |
| :---: | :---: | :---: | :---: |
|  ace  |  11   | eight |   8   |
|  two  |   2   | nine  |   9   |
| three |   3   |  ten  |  10   |
| four  |   4   | jack  |  10   |
| five  |   5   | queen |  10   |
|  six  |   6   | king  |  10   |
| seven |   7   | other |   0   |

**Note**: Commonly, aces can take the value of 1 or 11 but for simplicity we will assume that they can only take the value of 11.

Depending on your two cards and the card of the dealer, there is an optimal strategy for the first turn of the game, in which you have the following options:

    - Stand (S)
    - Hit (H)
    - Split (P)
    - Automatically win (W)

Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

- If you have a Blackjack (two cards that sum up to a value of 21) and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
- If you have a pair of aces you must always split them.
- If your cards sum up to 17 or higher you should always stand.
- If your cards sum up to 11 or lower you should always hit.
- If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.

You have three tasks:

## 1. Calculate the score of any given card

Implement a function to calculate the numerical value of a card given its name.

```go
value := ParseCard("ace")
fmt.Println(value)
// Output: 11
```

## 2. Determine if two cards make up a Blackjack.

Implement a function that returns true if two cards form a Blackjack, false otherwise.

```go
isBlackjack := Blackjack("queen", "ace")
fmt.Println(isBlackjack)
// Output: true
```

## 3. Determine the best gameplay choice for your first turn following Alex's strategy.

Implement a function that returns the string representation of a decision given your cards, the dealer's card and Alex's strategy.

```go
playerCard1 := "ace"
playerCard2 := "ace"
dealerCard := "two"
choice := FirstTurn(playerCard1, playerCard2, dealerCard)
fmt.Println(choice)
// Output: "P"
```
