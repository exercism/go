# Empty  `interface{}`

The empty interface can hold any type of data.
It is a special case of using [interfaces](TODO:_link_to_interface) where no method is specified and therefore any type can satisfy the empty interface.

To a programmer being used to a language like python (dynamically typed) or javascript (weakly typed) this sounds perfect.
But I'd recommend not using the empty interface at all, until one has figured out how to use Go's type system without it.
After that there are only a few very rare cases where it's usage actually makes sense.

## Problematic Usage

Often the first thing newcomers to Go try to use `interface{}` for is generic functions.

```go
func double(v interface{}) interface{} {
	switch val := v.(type) {
	case int:
		return val * 2
	case int64:
		return val * 2
	case float64:
		return val * 2
	case float32:
		return val * 2
	}

	log.Printf("type %T not supported by double function", v)
	return math.NaN()
}
```

This function has the following problems:
- Every type that can be multiplied by 2 needs to be handled separately. `v` cannot be multiplied by 2 directly as it
could be any type including a `string`. There are a lot more numerical types in Go than just the ones implemented above.
- Reflection is needed to check the type of the variable. In this case this is done with the `type switch`. In other cases
using the `reflect` package might be necessary. Reflection is expensive.
- Calling this function needs type casting afterwards `double(5).(int)` as the returned empty interface is of unknown type.
Not casting the type afterwards will litter the entire code base with variables of unknown type.
A variable of type `interface{}` cannot be used, unless casted to its underlying type first.
- Using `interface{}` disables type checks at compile time and introduces failures at runtime.
When using the function above with a non-supported type it will return the special "not a number" (`float64`) value.
Not checking for that after every call to the function is a big risk. Especially if the variable passed to the function
is already of type `interface{}` and with that its type unknown to the compiler and to the developer.

## Useful Usage

The empty `interface{}` can for example be useful if it does not matter what type is provided and any type should be supported.

One could imagine a customized `logData` function taking a message and additional data of any type:

```go
func logData(msg string, data ...interface{}) {
	log.Println(msg)
	fmt.Println("\tRelated data:")
	for _, val := range data {
		fmt.Printf("\t\t%+v\n", val)
	}
}
```

Another very useful example is for wrapping some data of any type into a json response without having to create another struct:

```go
func handleRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	t := time.Now()
	response, err := getResponse(r.Body)
	if err != nil {
		writeErr(w, err)
		return
	}

	res, err := json.Marshal(map[string]interface{}{
		"took": time.Since(t),
		"data": response,
	})
	if err != nil {
		writeErr(w, err)
		return
	}

	if _, err := w.Write(res); err != nil {
		w.WriteHeader(500)
		log.Println(err)
	}
}

func writeErr(w http.ResponseWriter, respErr error) {
	w.WriteHeader(500)

	res, err := json.Marshal(map[string]interface{}{
		"error": respErr,
	})
	if err != nil {
		log.Println(err)
	}

	if _, err := w.Write(res); err != nil {
		log.Println(err)
	}
}
```
