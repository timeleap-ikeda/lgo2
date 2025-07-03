package main

import (
	"errors"
	"fmt"
)

type Integer interface { //liststart0
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
} //listend0
// 全体が「型要素（type element）」
// int, int8などのひとつひとつは「型ターム（type termp）」

func divAndRemainder[T Integer](num, denom T) (T, T, error) {  //liststart2
	if denom == 0 {
		return 0, 0, errors.New("0で割ることはできません")
	}
	return num / denom, num % denom, nil
}

func main() {
	type MyInt int // intはIntegerに含まれてはいるが... //liststart1
	var myA MyInt = 10
	var myB MyInt = 20
	fmt.Println(divAndRemainder(myA, myB)) //コンパイル時のエラー   //listend1
}
