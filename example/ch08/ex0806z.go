package main

import (
	_ "errors"
	"fmt"
)

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
	// 全体が「型要素（type element）」
	// int, int8などのひとつひとつは「型ターム（type termp）」
}

func main() {
	var zzz Integer = 123
	fmt.Println(zzz)
}
