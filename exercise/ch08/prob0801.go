package main

import "fmt"

// 引数に指定された整数または浮動小数点数の値を2倍にするジェネリックな関数を定義せよ。必要なジェネリックなインタフェースを定義せよ

type ValidTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Doubler[T ValidTypes](t T) T {
	return t * 2
}

func main() {
	fmt.Println(Doubler(10))
	fmt.Println(Doubler(11.2))
}
