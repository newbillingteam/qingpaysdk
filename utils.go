package qingpay

import (
	"fmt"
	"net/url"
	"reflect"
)

func ToUrlValues(a map[string]interface{}) url.Values {
	if a == nil {
		return nil
	}
	urlEncoder := url.Values{}
	for key, value := range a {
		v := fmt.Sprintf("%v", value)
		urlEncoder.Add(key, v)
	}
	return urlEncoder
}

// func structToMap(t interface{}) map[string]interface{} {
// 	defer func() {
// 		if r := recover(); r != nil {

// 		}
// 	}()

// 	val := reflect.ValueOf(t)

// 	if val.Kind() != reflect.Struct {
// 		return nil
// 	}

// 	out := map[string]interface{}{}

// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)

// 		var value interface{}

// 		switch field.Kind() {
// 		case reflect.Struct:
// 			value = structToMap(field.Interface())
// 		case reflect.Ptr:
// 			indirectType := field.Elem()

// 			if indirectType.Kind() == reflect.Struct {
// 				value = structToMap(indirectType.Interface())
// 			} else {
// 				value = indirectType.Interface()
// 			}
// 		default:
// 			value = field.Interface()
// 		}

// 		out[val.Type().Field(i).Name] = value
// 	}

// 	return out
// }

func structToMap(t interface{}) map[string]interface{} {
	defer func() {
		if r := recover(); r != nil {
			panic("Panicked in structToMap. This should never happen.")
		}
	}()

	val := reflect.ValueOf(t)

	if val.Kind() != reflect.Struct {
		return nil
	}

	out := map[string]interface{}{}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		var value interface{}

		switch field.Kind() {
		case reflect.Struct:
			value = structToMap(field.Interface())
		case reflect.Ptr:
			indirectType := field.Elem()

			if indirectType.Kind() == reflect.Struct {
				value = structToMap(indirectType.Interface())
			} else if indirectType.Kind() != reflect.Invalid {
				value = indirectType.Interface()
			}
		default:
			value = field.Interface()
		}

		out[val.Type().Field(i).Name] = value
	}

	return out
}
