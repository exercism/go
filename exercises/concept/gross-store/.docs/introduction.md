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

Here are some operations that you can use with map

```go
  // Add a value in a map, you can use the `=` operator:
  foo["bar"] = 42
  // Here we update the element of `bar`
  foo["bar"] = 73
  // To retrieve a value map, you can use
  baz := foo["bar"]
  // To delete an item from a map, you can use
  delete(foo, "bar")
```

If you're trying to retrieve a non existed key, will return the zero value of your value type, it can confuse you, especially if your element is the same as the `ElementType` default value (for example, 0 for an int), to check whether the key exists in your map, you can use

```go
  value, exists := foo["baz"]
  // If the key "baz" doesn't exists,
  // value: 0; exists: false
```
