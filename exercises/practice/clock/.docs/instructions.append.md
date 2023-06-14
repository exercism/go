# Instructions append

The  clock package offers a straightforward implementation of a 24-hour clock in Go. It includes a Clock struct that represents the current time, along with various functions for manipulating and displaying the time. 
 
## Implementation Notes 
To ensure that two clocks representing the same time are equal to each other, the values of your Clock type must work with the  ==  operator. If your New  function returns a pointer instead of a value, your clocks will likely not work with  == . 
While it's not mandatory to use the time. Time type in the standard library as a basis for your Clock type, it could be useful to examine how constructors like Date and Now return values instead of pointers. Additionally, note that most time.Time methods use value receivers rather than pointer receivers. 
 
## Functions 
### New(hour, minute int) Clock 
The  New  function creates a new Clock instance with the specified hour and minute values. The hour value must be between 0 and 23, and the minute value must be between 0 and 59. 
 
### (c Clock) String() string 
The  String function returns a string representation of the current time in the format "HH:MM". For example, "02:45" represents 2:45 AM, and "18:30" represents 6:30 PM. 
 
### (c Clock) Add(minutes int) Clock 
The  Add function adds the specified number of minutes to the current time and returns a new  Clock instance with the updated time. The minutes value can be positive or negative. 
 
### (c Clock) Subtract(minutes int) Clock 
The  Subtract function subtracts the specified number of minutes from the current time and returns a new  Clock instance with the updated time. The minutes value can be positive or negative.

For some useful guidelines on when to use a value receiver or a pointer
receiver see [Go Wiki: Receiver Type](https://github.com/golang/go/wiki/CodeReviewComments#receiver-type)
