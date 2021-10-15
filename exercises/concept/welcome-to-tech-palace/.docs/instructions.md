# Instructions

There is an appliance store called "Tech Palace" nearby.
The owner of the store recently installed a big display to use for marketing messages and to show a special greeting when customers scan their loyalty cards at the entrance.
The display consists of lots of small LED lights and can show multiple lines of text.

The store owner needs your help with the code that is used to generate the text for the new display.

## 1. Create the welcome message

For most customers who scan their loyalty cards, the store owner wants to see `Welcome to the Tech Palace, ` followed by the name of the customer in capital letters on the display.

Implement the function `WelcomeMessage` that accepts the name of the customer as a `string` argument and returns the desired message as a `string`.

```go
WelcomeMessage("Judy")
// Output: Welcome to the Tech Palace, JUDY
```

## 2. Add a fancy border

For loyal customers that buy a lot at the store, the owner wants the welcome display to be more fancy by adding a line of stars before and after the welcome message.
They are not sure yet how many stars should be in the lines so they want that to be configurable.

Write a function `AddBorder` that accepts a welcome message (a `string`) and the number of stars per line (type `int`) as arguments.
It should return a `string` that consists of 3 lines, a line with the desired number of stars, then the welcome message as it was passed in, then another line of stars.

```go
AddBorder("Welcome!", 10)
// Output:
// **********
// Welcome!
// **********
```

## 3. Clean up old marketing messages

Before installing this new display, the store had a similar display that could only show non-configurable, static lines.
The owner would like to reuse some of the old marketing messages on the new display.
However, the data already includes a star border and some unfortunate whitespaces.
Your task is to clean up the messages so they can be re-used.

Implement a function `CleanUpMessage` that accepts the old marketing message as a string.
The function should first remove all stars from the text and afterwards remove the leading and trailing whitespaces from the remaining text.
The function should then return the cleaned up message.

```go
message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

CleanUpMessage(message)
// Output: BUY NOW, SAVE 10%
```
