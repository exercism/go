package flatten

import "reflect"

// Flatten recursively flattens slices and ignores nil values
func Flatten(nested interface{}) []interface{} {
	flattened := []interface{}{}
	nestedSlice, ok := nested.([]interface{})
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
