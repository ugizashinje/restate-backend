package transformer

import (
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/guregu/null.v4"
)

func Map(structure any) map[string]any {
	res := make(map[string]any)
	v := reflect.ValueOf(structure)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		// fieldName := v.Type().Field(i).Name
		// fieldValue := field.Interface()

		jsonTag := v.Type().Field(i).Tag.Get("json")
		tagSlice := strings.Split(jsonTag, ",")
		var key string
		if len(tagSlice) > 1 {
			key = strings.Trim(tagSlice[0], " \t")
		} else {
			key = strings.Trim(jsonTag, " \t")
		}

		switch field.Kind() {
		case reflect.Bool:
			res[key] = field.Interface()
		case reflect.String,
			reflect.Int, reflect.Int16, reflect.Int64, reflect.Int8,
			reflect.Uint, reflect.Uint16, reflect.Uint64, reflect.Uint8,
			reflect.Float32, reflect.Float64:
			if field.IsZero() {
				continue
			}
			res[key] = field.Interface()
		case reflect.Struct:
			name := v.Type().Field(i).Type.String()
			switch name {
			case "null.String":
				nullString := field.Interface().(null.String)
				if nullString.Valid {
					res[key] = nullString.String
				}
			case "null.Time":
				nullString := field.Interface().(null.Time)
				if nullString.Valid {
					res[key] = nullString.Time
				}
			}

		}

	}

	for k, v := range res {
		fmt.Printf("K: %s, V: %s \n", k, v)

	}
	return res
}
