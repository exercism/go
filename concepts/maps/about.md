In go, `map` is a built-in data type that represent hash table. In other programming language, you might know map as `dict` or associative array or a key/value store. If you're not familiar with such concept, map you can think `map` like a dictionary, which every word is the `key` and the definition is the `element` of the map.

> Before we begin, I'd like to point you to [go spec for map][gospec-map] and [go blog for map][goblog-map] to dig further into `map`.

Syntactically, `map` looks like this:

```go
map[KeyType]ElementType
```

_`KeyType` must be any [comparable type][gospec-comparable], while `ElementType` can be any valid type in go, which means you can store anything from primitive variable to a slice._

It's important to remember that `map` in go is **unordered**, if you try to loop trough a `map` and print the element, you might surprize yourself seeing that your elements printed in random order (give it a try if you like).

It is also important to know that each key is unique, meaning that assigning the same key twice will overwrite the value of the corresponding key.

`map` is reference type, which means if you pass it around, go won't copy the whole map. Instead what go will do is go copy the pointer of the map, this makes passing map to a function or variable cheap. The value of an uninitialized map is `nil`.

You can define map as follows (we also called this a **nil map**);

```go
  var foo map[string]int
```

To initialize a map, you can do:

```go
  // With map literal
  foo := map[string]int{}
```

or

```go
  // or with make function
  foo := make(map[string]int)
```

> A nil map is different from initialized map, writing to an nil map will cause a runtime error
>
> ```txt
> panic: assignment to entry in nil map
> ```
>
> Therefore it's important to initialize a map before using it

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

As we've seen before, detecting whether a map `map` is initialized or not is easy, we can simply compare them to `nil` but what about other maps? Well, since `maps` isn't a comparable, we can't use `==` operator, but `reflect` package has something to rescue us, it's called `DeepEqual`, we can use it like

```go
import "reflect"

// ....

equal := reflect.DeepEqual(map[string]int{}, map[string]int{})
fmt.Println(equal)
// Output: true
```

_But wait, if map isn't a comparable why are we able to compare them with `nil`? Well, the spec has made an exception for this, see the [comparable spec][gospec-comparable]_

The last one, if you're trying to write to a `map` from multiple go routine, that will trigger the race detector, [see this link][godoc-race-detector] and [here][goblog-race-detector]. Alternatively, you can use `sync.Map` or `atomic` or `mutex` to work around this issue.

[godoc-race-detector]: https://golang.org/doc/articles/race_detector.html
[goblog-race-detector]: https://blog.golang.org/race-detector
[goblog-map]: https://blog.golang.org/maps
[gospec-map]: https://golang.org/ref/spec#Map_types
[gospec-comparable]: https://golang.org/ref/spec#Comparison_operators
