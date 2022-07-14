# Instructions

In this exercise you'll be writing code to handle multilayer cache with Go channels.
Cache layer defines as a struct blow:

```go
type Cache struct {
	ResponseTime time.Duration
	Data         map[string]string
}
```

As you see, the `ResponseTime` is the time it takes to get(or set) the value from(to) the cache, and the `Data` is the data stored in the cache.

You have two tasks for reading and writing data form(to) cache layers.

## 1. Get value from fastest layer

We have multiple layers of cache, and we want to get the value from the fastest layer that has the key we want. the layers of cache given to you as `cacheLayers` parameter.
You have to get a value from the layer have the value for this key as fast as possible and return it with `nil` error or return `ErrKeyNotExist` error if the key does not exist in any layer.


```go
GetValueFromFastestLayer(cacheLayers []Cache, key string) (string, error)
```

## 2. Set value to all layers
The key and value are given to you as `key` and `value` parameters.
You have to set the value to all `cacheLayers` as fast as possible, you have maximum `max(cacheLayers.ResponseTime)` to set the value to all layers.

```go
SetValueToAllLayers(cacheLayers []Cache, key string, value string) error
```
