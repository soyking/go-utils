package gu

import "reflect"

// 检查一个类型是否为空
// 支持 指针 字符串 slice 结构体
func CheckEmpty(v reflect.Value) bool {
	// 指针获取其指向内容
	for v.Kind() == reflect.Ptr {
		// 指针为 nil
		if v.IsNil() {
			return true
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.String:
		if v.String() == "" {
			return true
		}
	case reflect.Slice:
		if v.Len() == 0 {
			return true
		}
	case reflect.Struct:
		n := v.NumField()
		// 结构体遍历所有 field 检查
		for i := 0; i < n; i++ {
			// 递归检查
			if CheckEmpty(v.Field(i)) == true {
				return true
			}
		}
	default:
		return false
	}

	return false
}
