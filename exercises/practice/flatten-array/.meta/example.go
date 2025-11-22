package flatten

import "reflect"

// Flatten recursively flattens slices and ignores nil values
func Flatten(nested any) []any {
	flattened := []any{}
	nestedSlice, ok := nested.([]any)
	if !ok {
		return flattened
	}
	for _, e := range nestedSlice {
		t := reflect.TypeOf(e)
		if t == nil {
			continue
		}
		switch kind := t.Kind(); kind {
		case reflect.Slice:
			flattened = append(flattened, Flatten(e)...)
		default:
			flattened = append(flattened, e)
		}
	}
	return flattened
}
