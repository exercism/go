# Introduction

Go does not have a concept of empty, null, or undefined for variable values. Variables declared without an explicit initial value default to the zero value for their respective type.

The zero value for primitive types such as booleans, numeric types, and strings are `false`, `0`, and `""`, respectively.

The identifier `nil`, meaning zero, is the zero value for more complex types such as pointers, functions, interfaces, slices, channels, and maps.

The following table details the zero value for Go's types.

| Type      | Zero Value |
| --------- | ---------- |
| boolean   | `false`    |
| numeric   | `0`        |
| string    | `""`       |
| pointer   | `nil`      |
| function  | `nil`      |
| interface | `nil`      |
| slice     | `nil`      |
| channel   | `nil`      |
| map       | `nil`      |

You may have noticed struct types are absent from the above table. That is because the zero value for a struct type depends on its fields. Structs are set to their zero value when all of its fields are set to their respective zero value.
