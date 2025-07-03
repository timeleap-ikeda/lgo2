package main

import(
	"fmt"
	"reflect"
)

func main() {
	type Foo struct { //liststart1
		A int    `myTag:"value"`
		B string `myTag:"value2"`
	}

	var f Foo
	ft := reflect.TypeOf(f)
	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i) // 各フィールドを表現する構造体
		fmt.Println(i, curField.Name, curField.Type.Name(), curField.Tag.Get("myTag"))
	} //listend1
}
