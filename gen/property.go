package gen

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const tagName = "property"

func classifyByProperty(group interface{}) error {
	v := reflect.ValueOf(group)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("Can only classify pointers, not %s (%v)", v.Kind(), v)
	}

	e := v.Elem()
	t := e.Type()

	var raws []json.RawMessage
	properties := make(map[string]reflect.Value)

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get(tagName)
		field := e.Field(i)
		switch {
		case tag == "RAW":
			raws = field.Interface().([]json.RawMessage)
		case tag != "":
			properties[tag] = field
		default:
			switch field.Kind() {
			case reflect.Slice:
				elem := field.Type().Elem()
				if elem.Kind() == reflect.Struct {
					for j := 0; j < field.Len(); j++ {
						err := classifyByProperty(field.Index(j).Addr().Interface())
						if err != nil {
							return fmt.Errorf("Couldn't classify %s element %d: %s", t.Field(i).Name, j, err)
						}
					}
				}
			case reflect.Struct:
				err := classifyByProperty(field.Addr().Interface())
				if err != nil {
					return fmt.Errorf("Couldn't classify %s: %s", t.Field(i).Name, err)
				}
			}
		}
	}

	for _, raw := range raws {
		var prop struct {
			Property string
		}
		if err := json.Unmarshal(raw, &prop); err != nil {
			return err
		}
		caseSlice, ok := properties[prop.Property]
		if !ok {
			return fmt.Errorf("Found property %s but no element tagged with that property", prop.Property)
		}
		oneCase := reflect.New(caseSlice.Type().Elem()).Interface()
		if err := json.Unmarshal(raw, &oneCase); err != nil {
			return err
		}
		caseSlice.Set(reflect.Append(caseSlice, reflect.ValueOf(oneCase).Elem()))
	}

	return nil
}
