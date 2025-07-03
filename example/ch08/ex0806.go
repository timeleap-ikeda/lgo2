package main

import (
	"errors"
	"fmt"
)

type Integer interface { //liststart1
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
	// 全体が「型要素（type element）」
	// int, int8などのひとつひとつは「型ターム（type termp）」
} //listend1

func divAndRemainder[T Integer](num, denom T) (T, T, error) {  //liststart2
	if denom == 0 {
		return 0, 0, errors.New("0で割ることはできません")
	}
	return num / denom, num % denom, nil
}

func main() {
	var a uint = 18_446_744_073_709_551_615 // _は桁の区切り（「2.1.2 リテラル」参照）
	var b uint = 9_223_372_036_854_775_808
	fmt.Println(divAndRemainder(a, b))  // 1 9223372036854775807 <nil>

	var c uint = 18
	var d uint = 7
	fmt.Println(divAndRemainder(c, d))  // 2 4 <nil>

	var e int8 = 18
	var f int8 = 0
	fmt.Println(divAndRemainder(e, f))  // 0 0 0で割ることはできません
}  //listend2
