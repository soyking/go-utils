package gu

import (
	"fmt"
	"reflect"
	"strings"
)

// 把一个结构体内容转成一个字符串，可以作为存储的 key
// 支持字符串，整数，布尔值，字符串数组
// 传入指针, 默认 _ 连接
func StructToKey(r interface{}, connector ...string) string {
	v := reflect.ValueOf(r).Elem()
	s := []string{}
	n := v.NumField()
	for i := 0; i < n; i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			s = append(s, f.String())
		case reflect.Slice:
			for j := 0; j < f.Len(); j++ {
				if f.Index(j).Kind() == reflect.String {
					s = append(s, f.Index(j).String())
				}
			}
		case reflect.Int:
			s = append(s, fmt.Sprintf("%d", f.Int()))
		case reflect.Bool:
			s = append(s, fmt.Sprintf("%t", f.Bool()))
		}
	}

	c := "_"
	if len(connector) > 0 {
		c = connector[0]
	}
	return strings.Join(s, c)
}
