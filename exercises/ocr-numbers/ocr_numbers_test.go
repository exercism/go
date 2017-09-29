// Go requirements:
//
// Define a function recognizeDigit as README Step 1 except make it recognize
// all ten digits 0 to 9.  Pick what you like for parameters and return values
// but make it useful as a subroutine for README step 2.
//
// For README Step 2 define,
//
//    func Recognize(string) []string
//
// and implement it using recognizeDigit.
//
// Input strings tested here have a \n at the beginning of each line and
// no trailing \n on the last line. (This makes for readable raw string
// literals.)
//
// For bonus points, gracefully handle misformatted data.  What should you
// do with a partial cell?  Discard it?  Pad with spaces?  Report it with a
// "?" character?  What should you do if the first character is not \n?

package ocr

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in  string
	out []string
}{
	{`
 _ 
| |
|_|
   `, []string{"0"}},
	{`
   
  |
  |
   `, []string{"1"}},
	{`
 _ 
 _|
|_ 
   `, []string{"2"}},
	{`
 _ 
 _|
 _|
   `, []string{"3"}},
	{`
   
|_|
  |
   `, []string{"4"}},
	{`
 _ 
|_ 
 _|
   `, []string{"5"}},
	{`
 _ 
|_ 
|_|
   `, []string{"6"}},
	{`
 _ 
  |
  |
   `, []string{"7"}},
	{`
 _ 
|_|
|_|
   `, []string{"8"}},
	{`
 _ 
|_|
 _|
   `, []string{"9"}},
	{`
    _ 
  || |
  ||_|
      `, []string{"10"}},
	{`
   
| |
| |
   `, []string{"?"}},
	{`
       _     _        _  _ 
  |  || |  || |  |  || || |
  |  ||_|  ||_|  |  ||_||_|
                           `, []string{"110101100"}},
	{`
       _     _           _ 
  |  || |  || |     || || |
  |  | _|  ||_|  |  ||_||_|
                           `, []string{"11?10?1?0"}},
	{`
    _  _     _  _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|| |
  ||_  _|  | _||_|  ||_| _||_|
                              `, []string{"1234567890"}},
	{`
    _  _ 
  | _| _|
  ||_  _|
         
    _  _ 
|_||_ |_ 
  | _||_|
         
 _  _  _ 
  ||_||_|
  ||_| _|
         `, []string{"123", "456", "789"}},
}

var _ = recognizeDigit // step 1.

func TestRecognize(t *testing.T) {
	for _, test := range tests {
		if res := Recognize(test.in); !reflect.DeepEqual(res, test.out) {
			t.Fatalf("Recognize(`%s`) = %q, want %q.", test.in, res, test.out)
		}
	}
}
