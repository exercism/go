In Go, uninitialized variables and their elements are given default values.

These default values are called the zero values for their respective types:

| Type      | Zero Value |
| --------- | ---------- |
| bool      | false      |
| numeric   | 0          |
| string    | ""         |
| pointer   | nil        |
| func      | nil        |
| interface | nil        |
| slice     | nil        |
| channel   | nil        |
| map       | nil        |

The identifier `nil`, meaning zero, is the zero value for the more complex types in Go.
