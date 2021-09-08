package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	// Each elem in slice OR each field
	numOfValues := 0
	var getFieldFunc func(int) reflect.Value

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		getFieldFunc = val.Index
		numOfValues = val.Len()
	case reflect.Struct:
		getFieldFunc = val.Field
		numOfValues = val.NumField()
	case reflect.String:
		fn(val.String())
	}

	for i := 0; i < numOfValues; i++ {
		walk(getFieldFunc(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
