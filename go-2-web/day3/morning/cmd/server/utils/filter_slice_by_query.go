package utils

import (
	"fmt"
	"reflect"
	"slices"
)

func FilterSliceByQueryListValue[T any](field int, values []string, result []T) []T {
	return slices.DeleteFunc(result, func(e T) bool {
		value := reflect.ValueOf(e).Field(field - 1)
		var final string
		switch value.Type().Kind() {
		case reflect.Int:
			final = fmt.Sprintf("%d", value.Int())
		case reflect.Bool:
			final = fmt.Sprintf("%v", value.Bool())
		case reflect.Float32:
			final = fmt.Sprintf("%0.2f", value.Float())
		case reflect.String:
			final = value.String()
		}
		for _, i := range values {
			if i == final {
				return false
			}
		}
		return true
	})
}
