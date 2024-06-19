package transformer

import (
	"fmt"
	"reflect"
	"strings"
	"time"

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

func Patch(src any, patch any) error {
	patchMap := make(map[string]any)
	v := reflect.ValueOf(patch)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		// fieldName := v.Type().Field(i).Name
		// fieldValue := field.Interface()

		key := v.Type().Field(i).Name

		switch field.Kind() {
		case reflect.Bool:
			patchMap[key] = field.Interface()
		case reflect.String,
			reflect.Int, reflect.Int16, reflect.Int64, reflect.Int8,
			reflect.Uint, reflect.Uint16, reflect.Uint64, reflect.Uint8,
			reflect.Float32, reflect.Float64:
			if field.IsZero() {
				continue
			}
			patchMap[key] = field.String()
		case reflect.Struct:
			name := v.Type().Field(i).Type.String()
			switch name {
			case "null.String":
				nullString := field.Interface().(null.String)
				if nullString.Valid {
					patchMap[key] = nullString.String
				}
			case "null.Int":
				nullInt := field.Interface().(null.Int)
				if nullInt.Valid {
					patchMap[key] = nullInt.Int64
				}
			case "null.Bool":
				nullInt := field.Interface().(null.Bool)
				if nullInt.Valid {
					patchMap[key] = nullInt.Bool
				}
			case "null.Time":
				nullString := field.Interface().(null.Time)
				if nullString.Valid {
					patchMap[key] = nullString.Time
				}
			}

		}

	}

	for k, v := range patchMap {
		fmt.Printf("K: %s, V: %s \n", k, v)
	}

	v = reflect.ValueOf(src)
	v = v.Elem()
	for k, e := range patchMap {
		fieldByName := reflect.ValueOf(src).Elem().FieldByName(k)
		if !fieldByName.IsValid() {
			continue
		}

		srcValue := v.FieldByName(k)
		switch {
		case fieldByName.CanInt():
			if i, ok := e.(int64); ok {
				srcValue.SetInt(i)
			}
		case fieldByName.Kind() == reflect.Bool:
			if b, ok := e.(bool); ok {
				srcValue.SetBool(b)
			}
		case fieldByName.Kind() == reflect.String:
			if str, ok := e.(string); ok {
				srcValue.SetString(str)
			}
		case fieldByName.Kind() == reflect.Struct:
			switch fieldByName.Type().String() {
			case "null.String":
				if str, ok := e.(string); ok {
					srcValue.Set(reflect.ValueOf(null.StringFrom(str)))
				}
			case "null.Int":
				if i, ok := e.(int64); ok {
					srcValue.Set(reflect.ValueOf(null.IntFrom(i)))
				}
			case "null.Bool":
				if b, ok := e.(bool); ok {
					srcValue.Set(reflect.ValueOf(null.BoolFrom(b)))
				}
			case "null.Time":
				if t, ok := e.(time.Time); ok {
					srcValue.Set(reflect.ValueOf(null.TimeFrom(t)))
				}

			}
		}

	}

	return nil
}
