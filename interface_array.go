package gu

import (
	"errors"
	"reflect"
)

// 任意数组不能作为 []interface{}
// 但是可以作为 interface{}
func ToArrayInterface(src interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(src)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("not slice")
	}

	var dst []interface{}
	for i := 0; i < v.Len(); i++ {
		dst = append(dst, v.Index(i).Interface())
	}

	return dst, nil
}
