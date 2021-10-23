package accumulate

// Accumulate transform given strings by using given function
func Accumulate(s []string, f func(st string) string) (result []string) {
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}
