# Introduction

In go, `map` is a built-in data type that maps keys to values. In other programming language, you might be familiar with the concept of `map` as a dictionary, hash table, key/value store or an associative array.

Syntactically, `map` looks like this:

```go
map[KeyType]ElementType
```

It is also important to know that each key is unique, meaning that assigning the same key twice will overwrite the value of the corresponding key.

To create a map, you can do:

```go
  // With map literal
  foo := map[string]int{}
```

or

```go
  // or with make function
  foo := make(map[string]int)
```

Here are some operations that you can do with a map

```go
  // Add a value in a map with the `=` operator:
  foo["bar"] = 42
  // Here we update the element of `bar`
  foo["bar"] = 73
  // To retrieve a map value, you can use
  baz := foo["bar"]
  // To delete an item from a map, you can use
  delete(foo, "bar")
```

If you try to retrieve the value for a key which does not exist in the map, it will return the zero value of the value type.
This can confuse you, especially if the default value of your `ElementType` (for example, 0 for an int), is a valid value.
To check whether a key exists in your map, you can use

```go
  value, exists := foo["baz"]
  // If the key "baz" does not exist,
  // value: 0; exists: false
```
