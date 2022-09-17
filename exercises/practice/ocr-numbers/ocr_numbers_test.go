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

func TestRecognize(t *testing.T) {
	for _, test := range tests {
	    t.Run(test.description, func(t *testing.T) {
	        res := Recognize(test.in)
	        assert.DeepEqual(t, res, test.out)
        }
	}
}
