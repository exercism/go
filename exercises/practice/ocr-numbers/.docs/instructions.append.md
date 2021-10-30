# Instructions append

## Implementation Notes

Define a function recognizeDigit as README Step 1 except make it recognize
all ten digits 0 to 9.  Pick what you like for parameters and return values
but make it useful as a subroutine for README step 2.

For README Step 2 define,

   func Recognize(string) []string

and implement it using recognizeDigit.

Input strings tested here have a \n at the beginning of each line and
no trailing \n on the last line. (This makes for readable raw string
literals.)

For bonus points, gracefully handle misformatted data.  What should you
do with a partial cell?  Discard it?  Pad with spaces?  Report it with a
"?" character?  What should you do if the first character is not \n?
