package ocr

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	description string
	in          string
	out         []string
}{
	{
		description: "single digit - 0",
		in: `
 _ 
| |
|_|
   `,
		out: []string{"0"}},
	{
		description: "single digit - 1",
		in: `
   
  |
  |
   `,
		out: []string{"1"}},
	{
		description: "single digit - 2",
		in: `
 _ 
 _|
|_ 
   `,
		out: []string{"2"}},
	{
		description: "single digit - 3",
		in: `
 _ 
 _|
 _|
   `,
		out: []string{"3"}},
	{
		description: "single digit - 4",
		in: `
   
|_|
  |
   `,
		out: []string{"4"}},
	{
		description: "single digit - 5",
		in: `
 _ 
|_ 
 _|
   `,
		out: []string{"5"}},
	{
		description: "single digit - 6",
		in: `
 _ 
|_ 
|_|
   `,
		out: []string{"6"}},
	{
		description: "single digit - 7",
		in: `
 _ 
  |
  |
   `,
		out: []string{"7"}},
	{
		description: "single digit - 8",
		in: `
 _ 
|_|
|_|
   `,
		out: []string{"8"}},
	{
		description: "single digit - 9",
		in: `
 _ 
|_|
 _|
   `,
		out: []string{"9"}},
	{
		description: "multiple digits - 10",
		in: `
    _ 
  || |
  ||_|
      `,
		out: []string{"10"}},
	{
		description: "multiple digits - 11",
		in: `
   
| |
| |
   `,
		out: []string{"?"}},
	{
		description: "multiple digits - 110101100",
		in: `
       _     _        _  _ 
  |  || |  || |  |  || || |
  |  ||_|  ||_|  |  ||_||_|
                           `,
		out: []string{"110101100"}},
	{
		description: "multiple digits - 11?10?1?0",
		in: `
       _     _           _ 
  |  || |  || |     || || |
  |  | _|  ||_|  |  ||_||_|
                           `,
		out: []string{"11?10?1?0"}},
	{
		in: `
    _  _     _  _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|| |
  ||_  _|  | _||_|  ||_| _||_|
                              `,
		out: []string{"1234567890"}},
	{
		description: "multiple numbers with multiple digits - 123 456 789",
		in: `
    _  _ 
  | _| _|
  ||_  _|
         
    _  _ 
|_||_ |_ 
  | _||_|
         
 _  _  _ 
  ||_||_|
  ||_| _|
         `,
		out: []string{"123", "456", "789"}},
}

var _ = recognizeDigit // step 1.

func TestRecognize(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if got := Recognize(tc.in); !reflect.DeepEqual(got, tc.out) {
				t.Fatalf("Recognize(%q) = %q, want: %q", tc.in, got, tc.out)
			}
		})
	}
}
