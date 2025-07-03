package main

import (
	"fmt"
	"reflect"
)

func hasNoValue(i any) bool {  //liststart1
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	case reflect.Pointer, reflect.Slice, reflect.Map, reflect.Func,
		 reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}  //listend1

func main() {
	var a any
	fmt.Println(a == nil, hasNoValue(a)) // true true

	var b *int
	fmt.Println(b == nil, hasNoValue(b)) // true true

	var c any = b
	fmt.Println(c == nil, hasNoValue(c)) // false true

	var d int
	fmt.Println(hasNoValue(d)) // false

	var e any = d
	fmt.Println(e == nil, hasNoValue(e)) // false false
}
