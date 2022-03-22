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

Depending on your two cards and the card of the dealer, there is a strategy for the first turn of the game, in which you have the following options:

- Stand (S)
- Hit (H)
- Split (P)
- Automatically win (W)

Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

- If you have a pair of aces you must always split them.
- If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
- If your cards sum up to a value within the range [17, 20] you should always stand.
- If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
- If your cards sum up to 11 or lower you should always hit.

## 1. Calculate the value of any given card.

Implement a function to calculate the numerical value of a card:

```go
value := ParseCard("ace")
fmt.Println(value)
// Output: 11
```

## 2. Implement the decision logic for the first turn.

Write a function that implements the decision logic as described above:

```go
func FirstTurn(card1, card2, dealerCard string) string
```

Here are some examples for the expected outcomes:

```go
FirstTurn("ace", "ace", "jack") == "P"
FirstTurn("ace", "king", "ace") == "S"
FirstTurn("five", "queen", "ace") == "H"
```
