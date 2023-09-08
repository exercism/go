package strain

func Keep[T any](list []T, fn func(T) bool) []T {
	result := []T{}

	for _, entry := range list {
		if fn(entry) {
			result = append(result, entry)
		}
	}

	return result
}

func Discard[T any](list []T, fn func(T) bool) []T {
	result := []T{}

	for _, entry := range list {
		if !fn(entry) {
			result = append(result, entry)
		}
	}

	return result
}
