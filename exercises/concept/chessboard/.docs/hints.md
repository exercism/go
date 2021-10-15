# Hints

## General

- An [integer value][integers] can be defined as one or more consecutive digits.
- A [map value][maps] stores key-value data

## 1. Given a Chessboard and a Rank, count how many squares are occupied

- You can iterate a [map][maps]
- Check if the value is true. If it is increment. This is to count pieces.
- You have to [explicitly return an integer][return] from a function.

## 2. Given a Chessboard and a File, count how many squares are occupied

- You'll first need to check the file is within range.
- Loop over the chessboard.
- Add one if the square is occupied.

## 3. Count how many squares are present in the given chessboard

- There are many ways to solve this.
- This should return how many squares are configured in a chess-board.

## 4. Count how many squares are occupied in the given chessboard

- Get the CountInRank for all ranks in the chessboard.

[functions]: https://golang.org/ref/spec#Function_declarations
[return]: https://golang.org/ref/spec#Return_statements
[operators]: https://golang.org/ref/spec#Operators
[integers]: https://golang.org/ref/spec#Integer_literals
[calls]: https://golang.org/ref/spec#Calls
