package util

import (
	"errors"
	"reflect"
)

func InSlice(val interface{}, arr interface{}) bool {
	slice, err := Sliceconv(arr)
	if err != nil {
		return false
	}

	for _, item := range slice {
		if item == val {
			return true
		}
	}

	return false
}

func Sliceconv(slice interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("warning: sliceconv: param \"slice\" should be on slice value")
	}

	l := v.Len()
	r := make([]interface{}, l)
	for i := 0; i < l; i++ {
		r[i] = v.Index(i).Interface()
	}
	return r, nil
}
