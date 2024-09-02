package o

import "reflect"

func IsNull(a any) bool {
	if a == nil {
		return true
	}
	v := reflect.ValueOf(a)
	switch v.Kind() {
	case reflect.Pointer, reflect.Slice, reflect.Map, reflect.Chan, reflect.UnsafePointer, reflect.Func, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

func NotNull(a any) bool {
	return !IsNull(a)
}
