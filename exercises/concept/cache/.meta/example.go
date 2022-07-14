package cache

func GetValueFromFastestLayer(cacheLayers []Cache, key string) (string, error) {
	ch := make(chan string, len(cacheLayers))
	for _, cache := range cacheLayers {
		go func(cache Cache) {
			value, _ := cache.Get(key)
			ch <- value
		}(cache)
	}

	for range cacheLayers {
		value := <-ch
		if value != "" {
			return value, nil
		}
	}

	return "", ErrKeyNotExist
}

func SetValueToAllLayers(cacheLayers []Cache, key string, value string) error {
	ch := make(chan error, len(cacheLayers))
	for _, cache := range cacheLayers {
		go func(cache Cache) {
			err := cache.Set(key, value)
			ch <- err
		}(cache)
	}

	for range cacheLayers {
		<-ch
	}

	return nil
}
